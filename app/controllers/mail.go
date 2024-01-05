package controllers

import (
	"dj/app/common"
	"dj/app/services"
	"dj/bootstrap"
	"dj/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var typeString = []string{"bind", "register", "login", "pwd"}

// SendMail
//
//	@Summary		发送邮件
//	@Description
//	@Tags			邮件
//	@Accept			application/json
//	@Produce		application/json
//
// @Param			params		body		request.SendMail		false		"请求参数"
// @Success			200			{object}	utils.Res
// @Router			/service/sendMail [post]
func SendMail(c *gin.Context) {
	mail := request.SendMail{}
	// type 绑定账号=bind 注册=register 登录=login 找回密码=pwd
	if e := c.ShouldBindJSON(&mail); e != nil {
		bootstrap.Log.Info("发送邮件,参数错误", zap.Any("params", mail))
		common.Fail(c, "参数错误")
		return
	}
	in := common.IsContainStr(typeString, mail.TypeStr)
	if false == in {
		bootstrap.Log.Info("发送邮件,参数错误-1", zap.Any("params", mail))
		common.Fail(c, "参数错误-1")
		return
	}
	r := &services.Kong{}
	err := r.SendMail(c, mail)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}
	common.Success(c, r)
}
