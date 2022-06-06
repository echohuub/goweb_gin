package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func (mc *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", mc.hello)
}

func (mc *HelloController) hello(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "world",
	})
}
