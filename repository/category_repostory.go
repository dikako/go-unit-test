package repository

import "github.com/dikako/go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
