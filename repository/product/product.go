package product

import (
	"acme/model"
)

type ProductRepository interface {
	GetProducts() ([]model.Product, error)
	GetProduct(id int) (model.Product, error)
	AddProduct(user model.Product) (id int, err error)
	UpdateProduct(id int, user *model.Product) (model.Product, error)
	DeleteProduct(id int) error
	Close()
}
