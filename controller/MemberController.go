package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb_gin/model"
	"goweb_gin/param"
	"goweb_gin/service"
	"goweb_gin/tool"
	"strconv"
	"time"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendCode", mc.sendSmsCode)
	engine.POST("/api/loginSms", mc.smsLogin)
	engine.GET("/api/captcha", mc.captcha)
	engine.GET("/api/verifyCaptcha", mc.verifyCaptcha)
	engine.POST("/api/loginName", mc.nameLogin)
	engine.POST("/api/upload/avatar", mc.uploadAvatar)
	engine.GET("/api/userinfo", mc.GetUserInfo)
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
		jsonStr, _ := json.Marshal(member)
		err := tool.SetSession(context, "user_"+string(member.Id), jsonStr)
		if err != nil {
			tool.Fail(context, "登录失败")
			return
		}

		context.SetCookie(tool.CookieName, strconv.Itoa(int(member.Id)), 10*60, "/", "localhost", true, true)

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
	jsonStr, _ := json.Marshal(member)
	err = tool.SetSession(context, "user_"+string(member.Id), jsonStr)
	if err != nil {
		tool.Fail(context, "登录失败")
		return
	}
	context.SetCookie(tool.CookieName, strconv.Itoa(int(member.Id)), 10*60, "/", "localhost", true, true)
	tool.Success(context, member)
}

func (*MemberController) uploadAvatar(context *gin.Context) {
	// 1.解析上传参数：file、user_id
	userId := context.PostForm("user_id")
	fmt.Println(userId)
	file, err := context.FormFile("avatar")
	if err != nil || userId == "" {
		tool.Fail(context, "参数解析失败")
		return
	}

	// 2.判断user_id对应的用户是否已经登录
	sess := tool.GetSession(context, "user_"+userId)
	if sess == nil {
		tool.Fail(context, "参数不合法")
		return
	}
	var member model.Member
	json.Unmarshal(sess.([]byte), &member)

	// 3.file保存到本地
	fileName := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = context.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.Fail(context, "头像更新失败")
		return
	}

	// 4.将保存后的文件本地路径，保存到用户表
	memberService := service.MemberService{}
	path := memberService.UploadAvatar(member.Id, fileName[1:])
	if path != "" {
		tool.Success(context, "http://localhost:8090"+path)
		return
	}
	// 5.返回结果
	tool.Fail(context, "上传失败")
}

func (*MemberController) GetUserInfo(context *gin.Context) {
	cookie, err := tool.CookieAuth(context)
	if err != nil {
		context.Abort()
		tool.Fail(context, "未登录")
		return
	}
	memberService := service.MemberService{}
	member := memberService.GetUserInfo(cookie.Value)
	if member == nil {
		tool.Fail(context, "用户不存在")
		return
	}
	tool.Success(context, map[string]interface{}{
		"id":            member.Id,
		"user_name":     member.UserName,
		"mobile":        member.Mobile,
		"register_time": member.RegisterTime,
		"avatar":        member.Avatar,
		"balance":       member.Balance,
		"city":          member.City,
	})
}
