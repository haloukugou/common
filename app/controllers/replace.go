package controllers

import (
	"dj/app/common"
	"dj/app/services"
	"dj/bootstrap"
	"dj/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Latest
//
//	@Summary		最新版本信息
//	@Description
//	@Tags			版本信息
//	@Accept			application/json
//	@Produce		application/json
//
// @Param			params		body		request.LatestParams		false		"请求参数"
// @Success			200			{object}	utils.Res
// @Router			/replace/latest [post]
func Latest(c *gin.Context) {
	params := request.LatestParams{}
	if e := c.ShouldBindJSON(&params); e != nil {
		bootstrap.Log.Info("查询最新版本失败", zap.Any("params", params))
		common.Fail(c, "参数错误")
		return
	}
	latest := services.Latest{}
	err := latest.LatestInfo(c, params)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}
	common.Success(c, latest)
}
