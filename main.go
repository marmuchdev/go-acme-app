package main

import (
	"acme/api"
	"acme/config"
	"acme/db/postgres"
	"acme/repository/product"
	"acme/repository/user"
	"acme/service"
	"fmt"
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		// Continue with the next handler
		next.ServeHTTP(writer, request)
	})
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World!")
}

func main() {

	// Load configuration
	//config := config.Postgres
	//config := config.InMemory
	//default to .env file
	config := config.LoadDatabaseConfig()
	//for inmemory
	//config := config.LoadDatabaseConfig(".env.inmemory")

	var userRepo user.UserRepository
	var productRepo product.ProductRepository

	switch config.Type {
	case "postgres":
		connectionString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", config.User, config.DBName, config.Password, config.Host, config.SSLMode)

		db, err := postgres.PostgresConnection(connectionString)
		if err != nil {
			panic(err)
		}

		//this is where we set up all our repositories using the postgres db
		userRepo = user.NewPostgresUserRepository(db.DB)
		productRepo = product.NewPostgresProductRepository(db.DB)

	case "inmemory":
		//for in-memory, we don't need db connection details as the repository itself does this
		userRepo = user.NewInMemoryUserRepository()

	default:
		fmt.Errorf("unsupported database type: %s", config.Type)
	}

	// Initialize services for users
	userService := service.NewUserService(userRepo)
	userAPI := api.NewUserAPI(userService)

	// Initialize services for products
	productService := service.NewProductService(productRepo)
	productAPI := api.NewProductAPI(productService)

	// Initialize router
	router := http.NewServeMux()
	//Users handlers
	router.HandleFunc("GET /", rootHandler)
	router.HandleFunc("GET /api/users", userAPI.GetUsers)
	router.HandleFunc("POST /api/users", userAPI.CreateUser)
	router.HandleFunc("GET /api/users/{id}", userAPI.GetSingleUser)
	router.HandleFunc("DELETE /api/users/{id}", userAPI.DeleteSingleUser)
	router.HandleFunc("PUT /api/users/{id}", userAPI.UpdateSingleUser)

	//Products handlers
	router.HandleFunc("GET /api/products", productAPI.GetProducts)
	router.HandleFunc("POST /api/products", productAPI.CreateProduct)
	router.HandleFunc("GET /api/products/{id}", productAPI.GetSingleProduct)
	router.HandleFunc("DELETE /api/products/{id}", productAPI.DeleteSingleProduct)
	router.HandleFunc("PUT /api/products/{id}", productAPI.UpdateSingleProduct)

	// Starting the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", CorsMiddleware(router))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
