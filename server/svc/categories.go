package svc

import (
	"gopherss/db"
	"gopherss/model"
)

type CategoriesService struct{}

var categoriesRepository = db.CategoriesRepository{}

func (s CategoriesService) GetMany(input model.ListCategoriesInput) ([]model.Category, error) {
	categories, err := categoriesRepository.GetMany(input)
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (s CategoriesService) Get(id int) (*model.Category, error) {
	category, err := categoriesRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s CategoriesService) Create(input model.AddCategoryInput) error {
	err := categoriesRepository.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

func (s CategoriesService) Update(id int, input model.UpdateCategoryInput) (bool, error) {
	updated, err := categoriesRepository.Update(id, input)
	if err != nil {
		return false, err
	}
	return updated, nil
}

func (s CategoriesService) Delete(id int, keepFeeds bool) error {
	err := categoriesRepository.Delete(id, keepFeeds)
	if err != nil {
		return err
	}
	return nil
}
