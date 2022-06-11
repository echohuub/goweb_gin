package dao

import (
	"fmt"
	"goweb_gin/model"
	"goweb_gin/tool"

	"github.com/google/logger"
)

type MemberDao struct {
	*tool.Orm
}

func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		logger.Fatal(err)
	}
	return result
}

func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	var sms model.SmsCode
	if _, err := md.Where(" phone = ? and code = ?", phone, code).Get(&sms); err != nil {
		fmt.Println(err.Error())
	}
	return &sms
}

func (md *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member
	if _, err := md.Where(" mobile = ?", phone).Get(&member); err != nil {
		fmt.Println(err.Error())
	}
	return &member
}

func (md *MemberDao) InsertMember(member model.Member) int64 {
	result, err := md.InsertOne(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

func (md *MemberDao) Query(name string, password string) *model.Member {
	var member = model.Member{}
	password = tool.EncoderSha256(password)
	_, err := md.Where(" user_name = ? and password = ?", name, password).Get(&member)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	return &member
}

func (md *MemberDao) UpdateMember(userId int64, fileName string) int64 {
	member := model.Member{Avatar: fileName}
	result, err := md.Where(" id = ?", userId).Update(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}
