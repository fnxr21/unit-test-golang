package service

import (
	"errors"
	"unit-test/mock/entity"
	"unit-test/mock/repository"
)

type CategoryService struct {
	// sample jika ini adalah connect ke database
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}
