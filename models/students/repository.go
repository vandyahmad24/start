package students

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Store(student Students) (Students, error)
	FindByID(ID int) (Students, error)
	Update(student Students) (Students, error)
	Delete(Student Students) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(student Students) (Students, error) {
	err := r.db.Create(&student).Error
	if err != nil {
		return student, err
	}
	return student, nil

}

func (r *repository) FindByID(ID int) (Students, error) {
	var student Students
	err := r.db.Where("id = ?", ID).Find(&student).Error
	if err != nil {
		return student, err
	}
	if student.ID == 0 {
		return student, errors.New("student not found")
	}
	return student, nil

}

func (r *repository) Update(student Students) (Students, error) {
	err := r.db.Save(&student).Error

	if err != nil {
		return student, err
	}

	return student, nil
}

func (r *repository) Delete(student Students) error {
	err := r.db.Delete(&student).Error
	if err != nil {
		return err
	}
	return nil
}
