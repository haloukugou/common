package services

import (
	"dj/app/common"
	"dj/bootstrap"
	"dj/config"
	"dj/constants"
	"dj/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strconv"
	"time"
)

type Kong struct {
}

func (k *Kong) SendMail(ctx *gin.Context, mail request.SendMail) error {
	key := constants.SendMailTypeLock + mail.TypeStr + ":" + mail.Mail
	r, e := bootstrap.Redis.Get(ctx, key).Result()
	if e == nil && r != "" {
		return fmt.Errorf("请勿重复发送")
	}

	code := common.RandStr(4)
	s := send(mail.Mail, code, mail.TypeStr)
	if false == s {
		return fmt.Errorf("发送邮件失败")
	}

	// 记录redis
	bootstrap.Redis.SetEX(ctx, key, code, 300*time.Second)
	return nil
}

func send(mail, code, typeStr string) bool {
	e := email.NewEmail()
	e.From = config.Config.Mail.From
	e.To = []string{mail}
	e.Subject = "邮箱验证码来喽"
	t := title(typeStr)
	e.Text = []byte(t + code + ".有效期为5分钟.")
	err := e.Send(config.Config.Mail.Host+":"+strconv.FormatInt(config.Config.Mail.Port, 10), smtp.PlainAuth("", config.Config.Mail.From, config.Config.Mail.Secret, config.Config.Mail.Host))
	if err != nil {
		bootstrap.Log.Error("发送邮件失败,msg=" + err.Error())
		return false
	}
	return true
}

func title(typeStr string) string {
	str := map[string]string{
		"bind":     "你的绑定邮箱验证码为:",
		"register": "你的注册验证码为:",
		"login":    "你的登录验证码为:",
		"pwd":      "你的修改密码验证码为:",
	}
	return str[typeStr]
}
