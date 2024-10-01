package repositories

import (
	"unit-test/api/models"

	"gorm.io/gorm"
)

type Penghuni interface {
	ClientGetNumber(number string) (models.Penghuni, error)
}

func (r *repository) ClientGetNumber(number string) (models.Penghuni, error) {
	var user models.Penghuni
	err := r.db.
		Take(&user, "Handphone = ?", number).
		Error

	return user, err
}

type repository struct {
	db *gorm.DB
}

func Repository(db *gorm.DB) *repository {
	return &repository{db}
}
