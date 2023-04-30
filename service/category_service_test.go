package service

import (
	"fmt"
	"github.com/dikako/go-unit-test/entity"
	"github.com/dikako/go-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	fmt.Println(categoryService.Get("1")) // <nil> Category Not Found

	category, err := categoryService.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Dika",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	fmt.Println(categoryService.Get("2")) // &{2 Dika} <nil>

	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
