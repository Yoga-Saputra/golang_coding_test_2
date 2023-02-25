package reviewproduct

type Service interface {
	GetReviewProduct() ([]ReviewProduct, error)
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
