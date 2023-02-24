package members

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Member, error)
	Insert(member Member) (Member, error)
	FindById(idMember int) (Member, error)
	Update(member Member) (Member, error)
	Delete(idMember int) (Member, error)
}

type repositoryDB struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryDB {
	return &repositoryDB{db}
}

// insert data to table members

func (r *repositoryDB) Insert(member Member) (Member, error) {
	err := r.db.Create(&member).Error

	if err != nil {
		return member, err
	}
	// config.Loggers("info", member)
	fmt.Println(member)
	return member, nil
}

func (r *repositoryDB) FindById(idMember int) (Member, error) {
	var member Member

	err := r.db.Where("id_member = ?", idMember).Find(&member).Error
	if err != nil {
		return member, err
	}

	return member, nil
}

func (r *repositoryDB) Update(member Member) (Member, error) {
	err := r.db.Save(&member).Error

	if err != nil {
		return member, err
	}

	return member, nil
}

func (r *repositoryDB) FindAll() ([]Member, error) {
	var allMember []Member
	err := r.db.Find(&allMember).Error

	if err != nil {
		return allMember, err
	}

	return allMember, nil
}

func (r *repositoryDB) Delete(idMember int) (Member, error) {
	var member Member

	// err := r.db.Where("id_member = ?", idMember).Find(&member).Error
	// if err != nil {
	// 	return member, err
	// }

	err := r.db.Where("id_member = ?", idMember).Delete(&member).Error
	if err != nil {
		return member, err
	}

	return member, nil
}
