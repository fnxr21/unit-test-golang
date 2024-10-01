package repository

import "unit-test/mock/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
