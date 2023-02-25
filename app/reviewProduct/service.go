package reviewproduct

import "errors"

type Service interface {
	GetReviewProduct() ([]ReviewProduct, error)
	GetReviewProductById(input int, fieldName string) (ReviewProduct, error)
	InsertReviewProduct(input InputReviewProduct) (ReviewProduct, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetReviewProduct() ([]ReviewProduct, error) {
	reviewProd, err := s.repository.FindAll()

	if err != nil {
		return reviewProd, err
	}

	return reviewProd, nil
}

func (s *service) InsertReviewProduct(input InputReviewProduct) (ReviewProduct, error) {
	product := ReviewProduct{
		ID_Product: input.ID_Product,
		Member_ID:  input.ID_Memebr,
		DescReview: input.DescReview,
	}

	newProduct, err := s.repository.Insert(product)

	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) GetReviewProductById(input int, fieldName string) (ReviewProduct, error) {
	rProduct, err := s.repository.FindById(input, fieldName)
	if err != nil {
		return rProduct, err
	}

	if rProduct.ID_Review == 0 {
		return rProduct, errors.New("no review Product found on that id")
	}

	return rProduct, nil
}
