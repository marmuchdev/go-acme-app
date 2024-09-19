package service

import (
	"acme/model"
	"acme/repository/product"
	"errors"
	"fmt"
)

type ProductService struct {
	repository product.ProductRepository
}

// NewUserService creates a new instance of UserService.
func NewProductService(repo product.ProductRepository) *ProductService {
	return &ProductService{
		repository: repo,
	}
}

func (s *ProductService) GetProducts() ([]model.Product, error) {

	products, err := s.repository.GetProducts()

	if err != nil {
		fmt.Println("Error getting products from DB:", err)
		return nil, errors.New("There was an error getting the products from the database.")
	}

	return products, nil

}

func (s *ProductService) DeleteProduct(id int) error {
	err := s.repository.DeleteProduct(id)

	if err != nil {
		fmt.Println("Error deleting product from DB:", err)
		return errors.New("Could not delete product")
	}

	return nil
}

func (s *ProductService) GetProduct(id int) (product model.Product, err error) {
	product, err = s.repository.GetProduct(id)

	if err != nil {
		fmt.Println("Error getting user from DB:", err)
		return model.Product{}, errors.New("Could not find user")
	}

	return product, nil

}

func (s *ProductService) UpdateProduct(id int, product model.Product) (UpdateProduct model.Product, err error) {
	updated, err := s.repository.UpdateProduct(id, &product)

	if err != nil {
		fmt.Println("Error updating product in DB:", err)
		return model.Product{}, errors.New("Could not update product")
	}

	return updated, nil

}

func (s *ProductService) CreateProduct(product model.Product) (id int, err error) {
	id, err = s.repository.AddProduct(product)

	if err != nil {
		fmt.Println("Error creating product in DB:", err)
		return 0, errors.New("Could not create product")
	}

	return id, nil
}
