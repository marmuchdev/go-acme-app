package api

import (
	"acme/model"
	"acme/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type ProductAPI struct {
	productService *service.ProductService
}

func NewProductAPI(productService *service.ProductService) *ProductAPI {
	return &ProductAPI{
		productService: productService,
	}
}

func (api *ProductAPI) UpdateSingleProduct(writer http.ResponseWriter, request *http.Request) {

	id, err := api.parseId(request.PathValue("id"))

	if err != nil {
		http.Error(writer, "Bad Request ID", http.StatusBadRequest)
		return
	}

	product, err := decodeProduct(request.Body)

	if err != nil {
		http.Error(writer, "Bad Request Body", http.StatusBadRequest)
		return
	}

	updated, err := api.productService.UpdateProduct(id, product)

	if err != nil {
		http.Error(writer, "User not found to update", http.StatusNotFound)
		return
	}

	json.NewEncoder(writer).Encode(updated)

}

func (api *ProductAPI) DeleteSingleProduct(writer http.ResponseWriter, request *http.Request) {

	id, err := api.parseId(request.PathValue("id"))

	if err != nil {
		http.Error(writer, "Bad Request ID", http.StatusBadRequest)
		return
	}

	err = api.productService.DeleteProduct(id)

	if err != nil {
		http.Error(writer, "Could not delete user", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)

}

func (api *ProductAPI) GetSingleProduct(writer http.ResponseWriter, request *http.Request) {

	id, err := api.parseId(request.PathValue("id"))

	if err != nil {
		http.Error(writer, "Bad Request ID", http.StatusBadRequest)
		return
	}

	product, err := api.productService.GetProduct(id)

	if err != nil {
		http.Error(writer, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(writer).Encode(product)

}

func (api *ProductAPI) CreateProduct(writer http.ResponseWriter, request *http.Request) {

	product, err := decodeProduct(request.Body)

	if err != nil {
		http.Error(writer, "Bad Request Body", http.StatusBadRequest)
		return
	}

	id, err := api.productService.CreateProduct(product)

	if err != nil {
		http.Error(writer, "Product not created", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, "Product created successfully: %d", id)

}

func (api *ProductAPI) GetProducts(writer http.ResponseWriter, request *http.Request) {

	products, err := api.productService.GetProducts()

	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(products)

}

func (api *ProductAPI) parseId(idStr string) (id int, err error) {

	id, err = strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error parsing ID:", err)
		return 0, err
	}

	return id, nil

}

func decodeProduct(body io.ReadCloser) (product model.Product, err error) {

	err = json.NewDecoder(body).Decode(&product)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		return model.Product{}, err
	}

	return product, nil
}
