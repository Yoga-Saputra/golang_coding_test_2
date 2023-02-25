package products

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Product, error)
	Insert(product Product) (Product, error)
	FindById(idProduct int) (Product, error)
	Update(product Product) (Product, error)
	Delete(idproduct int) (Product, error)
}

type repositoryDB struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryDB {
	return &repositoryDB{db}
}

// insert data to table Product

func (r *repositoryDB) Insert(product Product) (Product, error) {
	err := r.db.Create(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repositoryDB) FindById(idProduct int) (Product, error) {
	var product Product

	err := r.db.Where("id_product = ?", idProduct).Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repositoryDB) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repositoryDB) FindAll() ([]Product, error) {
	var allProduct []Product
	err := r.db.Find(&allProduct).Error

	if err != nil {
		return allProduct, err
	}

	return allProduct, nil
}

func (r *repositoryDB) Delete(idProduct int) (Product, error) {
	var product Product

	err := r.db.Where("id_product = ?", idProduct).Delete(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
