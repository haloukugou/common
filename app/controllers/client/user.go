package client

import (
	"dj/app/common"
	"dj/app/request/clientRequest"
	"dj/app/services/clientService"
	"dj/app/validate"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	params := clientRequest.Register{}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		common.Fail(ctx, validate.GetErrorMsg(params, err))
		return
	}
	reg := clientService.Kong{}
	e := reg.UserRegister(ctx, params)
	if e != nil {
		common.Fail(ctx, e.Error())
		return
	}
	common.Ok(ctx)
}
