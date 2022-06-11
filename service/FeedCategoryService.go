package service

import (
	"goweb_gin/dao"
	"goweb_gin/model"
)

type FoodCategoryService struct {
}

func (s *FoodCategoryService) Categories() ([]model.FoodCategory, error) {
	categoryDao := dao.NewFoodCategoryDao()
	return categoryDao.QueryCategories()
}
