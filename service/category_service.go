package service

import (
	"errors"
	"unit-test/entity"
	"unit-test/repository"
)

type CategoryService struct {
	// sample jika ini adalah connect ke database
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
