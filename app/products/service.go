package products

import (
	"errors"
)

type Service interface {
	GetProducts() ([]Product, error)
	InsertProduct(input InputProduct) (Product, error)
	GetProductById(input GetProductDetailInput) (Product, error)
	UpdateProduct(input InputProduct, IDProduct GetProductDetailInput) (Product, error)
	DeleteProduct(input GetProductDetailInput) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProducts() ([]Product, error) {

	product, err := s.repository.FindAll()

	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) InsertProduct(input InputProduct) (Product, error) {
	product := Product{
		Name:    input.Name,
		Product: input.Product,
		Price:   input.Price,
	}

	newProduct, err := s.repository.Insert(product)

	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) GetProductById(input GetProductDetailInput) (Product, error) {
	product, err := s.repository.FindById(input.IDProduct)
	if err != nil {
		return product, err
	}

	if product.IDProduct == 0 {
		return product, errors.New("no product found on that id")
	}

	return product, nil
}

func (s *service) UpdateProduct(input InputProduct, inputIDProduct GetProductDetailInput) (Product, error) {
	product, err := s.repository.FindById(inputIDProduct.IDProduct)

	if err != nil {
		return product, err
	}

	if product.IDProduct == 0 {
		return product, errors.New("no product found on that id")
	}

	product.Name = input.Name
	product.Product = input.Product
	product.Price = input.Price

	updateProduct, err := s.repository.Update(product)

	if err != nil {
		return updateProduct, err
	}

	return updateProduct, nil
}

func (s *service) DeleteProduct(input GetProductDetailInput) (Product, error) {

	product, err := s.repository.FindById(input.IDProduct)

	if err != nil {
		return product, err
	}

	if product.IDProduct == 0 {
		return product, errors.New("no product found on that id")
	}

	_, err = s.repository.Delete(input.IDProduct)

	if err != nil {
		return product, err
	}

	return product, nil
}
