package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb_gin/param"
	"goweb_gin/service"
	"goweb_gin/tool"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendCode", mc.sendSmsCode)
	engine.POST("/api/loginSms", mc.smsLogin)
	engine.GET("/api/captcha", mc.captcha)
	engine.GET("/api/verifyCaptcha", mc.verifyCaptcha)
}

func (mc *MemberController) sendSmsCode(context *gin.Context) {
	phone, exist := context.GetQuery("phone")
	if !exist {
		tool.Fail(context, "参数解析失败")
		return
	}
	service := service.MemberService{}
	isSend := service.SendCode(phone)
	if isSend {
		tool.Success(context, "发送成功")
		return
	}
	tool.Fail(context, "发送失败")
}

func (*MemberController) smsLogin(context *gin.Context) {
	var smsLoginParam param.SmsLoginParam
	err := tool.Decode(context.Request.Body, &smsLoginParam)
	if err != nil {
		tool.Fail(context, "参数解析失败")
		return
	}

	service := service.MemberService{}
	member := service.SmsLogin(smsLoginParam)
	if member != nil {
		tool.Success(context, member)
		return
	}
	tool.Fail(context, "登录失败")
}

// 生成验证码
func (*MemberController) captcha(context *gin.Context) {
	captcha := tool.GenerateCaptcha(context)
	tool.Success(context, gin.H{
		"captcha_result": captcha,
	})
}

// 验证验证码
func (*MemberController) verifyCaptcha(context *gin.Context) {
	var captcha tool.CaptchaResult
	err := tool.Decode(context.Request.Body, &captcha)
	if err != nil {
		tool.Fail(context, "参数解析失败")
		return
	}
	result := tool.VerifyCaptcha(captcha.Id, captcha.VertifyValue)
	if result {
		fmt.Println("验证通过")
	} else {
		fmt.Println("验证失败")
	}
}
