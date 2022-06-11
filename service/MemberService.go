package service

import (
	"fmt"
	"goweb_gin/dao"
	"goweb_gin/model"
	"goweb_gin/param"
	"goweb_gin/tool"
	"math/rand"
	"strconv"
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

func (ms *MemberService) Login(name string, password string) *model.Member {
	// 1.使用用户名查询用户，如果存在就直接返回
	md := dao.MemberDao{tool.DBEngine}
	member := md.Query(name, password)
	if member.Id != 0 {
		return member
	}
	// 2.如果不存在，则注册新用户并返回
	user := model.Member{}
	user.UserName = name
	user.Password = tool.EncoderSha256(password)
	user.RegisterTime = time.Now().Unix()

	result := md.InsertMember(user)
	user.Id = result

	return &user
}

func (ms *MemberService) UploadAvatar(userId int64, fileName string) string {
	memberDao := dao.MemberDao{tool.DBEngine}
	result := memberDao.UpdateMember(userId, fileName)
	if result == 0 {
		return ""
	}
	return fileName
}

func (ms *MemberService) GetUserInfo(userId string) *model.Member {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil
	}
	memberDao := dao.MemberDao{tool.DBEngine}
	return memberDao.QueryMemberById(int64(id))
}
