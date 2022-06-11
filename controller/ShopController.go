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
	tool.Success(context, list)
}
