package reviewproduct

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]ReviewProduct, error)
}

type repositoryDB struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryDB {
	return &repositoryDB{db}
}

func (r *repositoryDB) FindAll() ([]ReviewProduct, error) {
	var rProduct []ReviewProduct
	err := r.db.Debug().Preload("Members").Find(&rProduct).Error

	if err != nil {
		return rProduct, err
	}
	fmt.Println(rProduct)

	return rProduct, nil
}
