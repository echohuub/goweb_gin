package controller

import (
	"github.com/gin-gonic/gin"
	"goweb_gin/service"
	"goweb_gin/tool"
)

type ShopController struct {
}

func (sc *ShopController) Router(engine *gin.Engine) {
	engine.GET("/api/shops", sc.GetShopList)
	engine.GET("/api/searchShop", sc.SearchShop)
}

func (sc *ShopController) GetShopList(context *gin.Context) {
	longitude := context.Query("longitude")
	latitude := context.Query("latitude")

	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "116.34"
		latitude = "40.34"
	}

	shopService := service.ShopService{}
	list := shopService.GetShopList(longitude, latitude)
	for _, shop := range list {
		shop.Supports = shopService.GetService(shop.Id)
	}
	tool.Success(context, list)
}

func (sc *ShopController) SearchShop(context *gin.Context) {
	longitude := context.Query("longitude")
	latitude := context.Query("latitude")
	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "116.34"
		latitude = "40.34"
	}

	keyword := context.Query("keyword")
	if keyword == "" {
		tool.Fail(context, "请输入内容")
		return
	}
	shopService := service.ShopService{}
	list := shopService.SearchShop(longitude, latitude, keyword)
	tool.Success(context, list)
}
