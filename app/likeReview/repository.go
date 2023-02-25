package likereview

import "gorm.io/gorm"

type Repository interface {
	Like(likeReview LikeReview) (LikeReview, error)
	DisLike(idReview int, idMember int) (LikeReview, error)
}

type repositoryDB struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryDB {
	return &repositoryDB{db}
}

func (r *repositoryDB) Like(likeReview LikeReview) (LikeReview, error) {
	err := r.db.Create(&likeReview).Error
	if err != nil {
		return likeReview, err
	}

	return likeReview, nil
}

func (r *repositoryDB) DisLike(idReview int, idMember int) (LikeReview, error) {
	var lReview LikeReview
	err := r.db.Where("id_review = ?", idReview).Where("id_member = ?", idMember).Delete(&lReview).Error

	if err != nil {
		return lReview, err
	}
	return lReview, nil
}
