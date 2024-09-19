package product

import (
	"acme/model"
	"errors"
	"slices"
)

type InMemoryProductRepository struct{}

var count int = 3
var products []model.Product

func NewInMemoryProductRepository() *InMemoryProductRepository {
	InitDB() // Initialize the in-memory database with sample data
	return &InMemoryProductRepository{}
}

func InitDB() {
	products = []model.Product{
		{ID: 1, Name: "User 1", Price: 9.99, Stock_count: 99},
		{ID: 1, Name: "User 1", Price: 9.99, Stock_count: 99},
		{ID: 1, Name: "User 1", Price: 9.99, Stock_count: 99},
	}
}
func (repo *InMemoryProductRepository) GetProducts() ([]model.Product, error) {
	return products, nil
}

func (repo *InMemoryProductRepository) AddProduct(product model.Product) (id int, err error) {
	count++
	product.ID = count

	products = append(products, product)

	return count, nil
}

func (repo *InMemoryProductRepository) GetProduct(id int) (model.Product, error) {
	var product model.Product

	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}

	return product, errors.New("product id not found.")

}

func (repo *InMemoryProductRepository) DeleteProduct(id int) error {

	for index, product := range products {
		if product.ID == id {
			products = slices.Delete(products, index, index+1)
			return nil
		}
	}

	return errors.New("Product id not found to delete.")

}

/*
WITHOUT using pointers
*/

func (repo *InMemoryProductRepository) UpdateProduct(id int, updatedProduct *model.Product) (model.Product, error) {

	for index, product := range products {
		if product.ID == id {
			products[index].Name = updatedProduct.Name
			return product, nil
		}
	}

	return model.Product{}, errors.New("User id not found to update.")

}

/*
	USING Pointers

func UpdateUser(id int, updatedUser model.User) *model.User {

	for index, user := range users {
		if user.ID == id {
			user := &users[index]
			user.Name = updatedUser.Name
			return user
		}
	}

	return &User{}

}
*/
func (repo *InMemoryProductRepository) Close() {

}
