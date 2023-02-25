package likereview

type Service interface {
	Like(input LikeReview) (LikeReview, error)
	DisLike(idReview int, idMember int) (LikeReview, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Like(input LikeReview) (LikeReview, error) {
	lReview := LikeReview{
		ID_Review: input.ID_Review,
		ID_Member: input.ID_Member,
	}

	newLReview, err := s.repository.Like(lReview)

	if err != nil {
		return newLReview, err
	}

	return newLReview, nil
}

func (s *service) DisLike(idReview int, idMember int) (LikeReview, error) {
	likeReview, err := s.repository.DisLike(idReview, idMember)

	if err != nil {
		return likeReview, err
	}

	return likeReview, nil
}
