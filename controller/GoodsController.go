package controller

import (
	"github.com/gin-gonic/gin"
	"goweb_gin/service"
	"goweb_gin/tool"
	"strconv"
)

type GoodsController struct {
}

func (gc *GoodsController) Router(app *gin.Engine) {
	app.GET("/api/goods", gc.getGoods)
}

func (gc *GoodsController) getGoods(context *gin.Context) {
	shopId, exist := context.GetQuery("shop_id")
	if !exist {
		tool.Fail(context, "参数错误")
		return
	}

	id, err := strconv.Atoi(shopId)
	if err != nil {
		tool.Fail(context, "参数错误")
		return
	}
	goodsService := service.GoodsService{}
	goods := goodsService.GetGoods(int64(id))
	tool.Success(context, goods)
}
