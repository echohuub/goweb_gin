package service

import (
	"fmt"
	"goweb_gin/dao"
	"goweb_gin/model"
	"goweb_gin/tool"
	"math/rand"
	"time"
)

type MemberService struct {
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
