package clientRequest

import (
	"dj/app/validate"
)

type Register struct {
	Mobile string `json:"mobile" binding:"required,mobile"`
	Code   string `json:"code" binding:"required"`
}

func (r *Register) GetMessages() validate.ValidatorMessages {
	return validate.ValidatorMessages{
		"mobile.required": "请输入手机号",
		"mobile.mobile":   "请输入正确的手机号",
		"code.required":   "请输入手机验证码",
	}
}
