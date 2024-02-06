package service

import (
	"github.com/ormesino/e-commerce/internal/database"
	"github.com/ormesino/e-commerce/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{
		ProductDB: productDB,
	}
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetProductByCategoryID(id string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductByCategoryID(id)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) CreateProduct(name, description, imageURL, categoryID string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, imageURL, categoryID, price)
	_, err := ps.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
