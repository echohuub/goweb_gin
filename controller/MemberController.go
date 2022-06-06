package controller

import (
	"goweb_gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendCode", mc.sendSmsCode)
}

func (mc *MemberController) sendSmsCode(context *gin.Context) {
	phone, exist := context.GetQuery("phone")
	if !exist {
		context.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数解析失败",
		})
		return
	}
	service := service.MemberService{}
	isSend := service.SendCode(phone)
	if isSend {
		context.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "发送成功",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "发送失败",
	})

}
