package product

import (
	"acme/model"
	"errors"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type PostgresProductRepository struct {
	DB *sqlx.DB
}

func NewPostgresProductRepository(db *sqlx.DB) *PostgresProductRepository {
	return &PostgresProductRepository{DB: db}
}

func (repo *PostgresProductRepository) GetProducts() ([]model.Product, error) {

	products := []model.Product{}

	err := sqlx.Select(repo.DB, &products, "SELECT * FROM products")
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return []model.Product{}, errors.New("Database could not be queried")
	}

	return products, nil
}

// GetUser retrieves a user by ID from the database.
func (repo *PostgresProductRepository) GetProduct(id int) (model.Product, error) {
	product := []model.Product{}
	err := sqlx.Select(repo.DB, &product, "SELECT * FROM products WHERE id = ($1)", strconv.Itoa(id))
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return model.Product{}, errors.New("Database could not be queried")
	}

	return product[0], nil
}

func (repo *PostgresProductRepository) AddProduct(product model.Product) (id int, err error) {

	err = repo.DB.QueryRow("INSERT INTO products (name, price, stock_count) VALUES ($1,$2,$3) RETURNING id", product.Name, product.Price, product.Stock_count).Scan(&id)
	if err != nil {
		fmt.Println("Error inserting product into the database:", err)
		return 0, errors.New("Could not insert product")
	}

	return id, nil
}
func (repo *PostgresProductRepository) UpdateProduct(id int, product *model.Product) (model.Product, error) {
	query := "UPDATE products SET name = ($1) WHERE id = ($2) RETURNING id, name"
	rows, err := repo.DB.Queryx(query, product.Name, id)

	if err != nil {
		fmt.Println("Error querying the database:", err)
		return model.Product{}, errors.New("Database could not be queried")
	}
	defer rows.Close()

	var updatedProduct []model.Product

	for rows.Next() {
		var u model.Product
		if err := rows.StructScan(&u); err != nil {
			return model.Product{}, err
		}
		updatedProduct = append(updatedProduct, u)
	}
	return updatedProduct[0], nil
}

func (repo *PostgresProductRepository) DeleteProduct(id int) error {

	product := []model.Product{}
	err := sqlx.Select(repo.DB, &product, "DELETE FROM products WHERE id = ($1)", strconv.Itoa(id))
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return errors.New("Database could not be queried")
	}
	return nil
}
func (repo *PostgresProductRepository) Close() {

}
