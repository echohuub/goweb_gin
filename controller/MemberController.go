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
	engine.POST("/api/loginName", mc.nameLogin)
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

// 用户名+密码，验证码登录
func (*MemberController) nameLogin(context *gin.Context) {
	// 1.解析用户登录传递参数
	var loginParam param.LoginParam
	err := tool.Decode(context.Request.Body, &loginParam)
	if err != nil {
		tool.Fail(context, "参数解析失败")
		return
	}

	// 2.验证验证码
	validate := tool.VerifyCaptcha(loginParam.Id, loginParam.Value)
	if !validate {
		tool.Fail(context, "验证码不正确，请重新验证")
		return
	}

	// 3.登录
	ms := service.MemberService{}
	member := ms.Login(loginParam.Name, loginParam.Password)
	if member.Id == 0 {
		tool.Fail(context, "登录失败")
		return
	}
	tool.Success(context, member)
}
