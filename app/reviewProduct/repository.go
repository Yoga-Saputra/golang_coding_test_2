package reviewproduct

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]ReviewProduct, error)
	Insert(reviewProduct ReviewProduct) (ReviewProduct, error)
	FindById(id int, fieldName string) (ReviewProduct, error)
}

type repositoryDB struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryDB {
	return &repositoryDB{db}
}

func (r *repositoryDB) FindAll() ([]ReviewProduct, error) {
	var rProduct []ReviewProduct
	err := r.db.Debug().Preload("Members").Preload("ReviewProductMembers").Find(&rProduct).Error

	if err != nil {
		return rProduct, err
	}
	fmt.Println(rProduct)

	return rProduct, nil
}

func (r *repositoryDB) FindById(id int, fieldName string) (ReviewProduct, error) {
	var reviewProduct ReviewProduct

	field := fmt.Sprintf("%s = ?", fieldName)
	err := r.db.Where(field, id).Find(&reviewProduct).Error
	if err != nil {
		return reviewProduct, err
	}

	return reviewProduct, nil
}

func (r *repositoryDB) Insert(reviewProduct ReviewProduct) (ReviewProduct, error) {
	err := r.db.Create(&reviewProduct).Error

	if err != nil {
		return reviewProduct, err
	}

	return reviewProduct, nil
}
