package service

import (
	"fmt"
	"goweb_gin/dao"
	"goweb_gin/model"
	"goweb_gin/param"
	"goweb_gin/tool"
	"math/rand"
	"time"
)

type MemberService struct {
}

func (*MemberService) SmsLogin(smsLoginParam param.SmsLoginParam) *model.Member {
	// 1.获取到手机号和验证码
	// 2.验证手机号+验证码是否正确
	memberDao := dao.MemberDao{tool.DBEngine}
	sms := memberDao.ValidateSmsCode(smsLoginParam.Phone, smsLoginParam.Code)
	if sms.Id == 0 {
		return nil
	}
	// 3.根据手机号member表中查询记录
	member := memberDao.QueryByPhone(smsLoginParam.Phone)
	if member.Id != 0 {
		return member
	}
	// 4.新创建一个member记录，并保存
	user := model.Member{}
	user.UserName = smsLoginParam.Phone
	user.Mobile = smsLoginParam.Phone
	user.RegisterTime = time.Now().Unix()
	user.Id = memberDao.InsertMember(user)

	return &user
}

func (ms *MemberService) SendCode(phone string) bool {
	// 产生一个验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	fmt.Println("验证码：" + code)
	// todo 调用阿里云sdk完成发送

	smsCode := model.SmsCode{Phone: phone, Code: code, BizId: "", CreateTime: time.Now().Unix()}

	memberDao := dao.MemberDao{tool.DBEngine}
	result := memberDao.InsertCode(smsCode)
	return result > 0
}
