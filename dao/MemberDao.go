package dao

import (
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
