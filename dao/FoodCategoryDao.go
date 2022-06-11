package dao

import (
	"goweb_gin/model"
	"goweb_gin/tool"
)

type FoodCategoryDao struct {
	*tool.Orm
}

func NewFoodCategoryDao() *FoodCategoryDao {
	return &FoodCategoryDao{tool.DBEngine}
}

func (fcd *FoodCategoryDao) QueryCategories() ([]model.FoodCategory, error) {
	var categories []model.FoodCategory
	err := fcd.Find(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
