package controller

import (
	"github.com/gin-gonic/gin"
	"goweb_gin/service"
	"goweb_gin/tool"
)

type FoodCategoryController struct {
}

func (fcc *FoodCategoryController) Router(engine *gin.Engine) {
	engine.GET("/api/foodCategory", fcc.foodCategory)
}

func (fcc *FoodCategoryController) foodCategory(context *gin.Context) {
	foodCategoryService := service.FoodCategoryService{}
	categories, err := foodCategoryService.Categories()
	if err != nil {
		tool.Fail(context, "请求失败")
		return
	}
	tool.Success(context, categories)
}
