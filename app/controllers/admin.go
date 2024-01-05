package controllers

import (
	"dj/app/common"
	services2 "dj/app/services"
	"dj/bootstrap"
	"dj/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AdminLogin
//
//	@Summary		管理员登录
//	@Description
//	@Tags			管理员
//	@Accept			application/json
//	@Produce		application/json
//
// @Param			params		body		request.AdminLoginParams		false		"请求参数"
// @Success			200			{object}	utils.Res
// @Router			/admin/login [post]
func AdminLogin(c *gin.Context) {
	params := request.AdminLoginParams{}
	if e := c.ShouldBindJSON(&params); e != nil {
		bootstrap.Log.Info("管理员登录失败", zap.Any("params", params))
		common.Fail(c, "参数错误")
		return
	}
	result := services2.LoginResult{}
	err := result.Login(c, params)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}
	common.Success(c, result)
}

// ApkList
//
//	@Summary		app列表
//	@Description
//	@Tags			管理员
//	@Accept			application/json
//	@Produce		application/json
//
// @Security 		admintoken
// @Success			200			{object}	utils.Res
// @Router			/admin/apkList [post]
func ApkList(c *gin.Context) {
	params := request.ApkListParams{}
	if e := c.ShouldBindJSON(&params); e != nil {
		bootstrap.Log.Info("查询apk信息列表失败", zap.Any("params", params))
		common.Fail(c, "参数错误")
		return
	}
	re := services2.ApkList{}
	err := re.ApkList(c, params)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}
	common.Success(c, re)
}

// Upload
//
//	@Summary		上传apk文件
//	@Description
//	@Tags			管理员
//	@Accept			application/json
//	@Produce		application/json
//
// @Success			200			{object}	utils.Res
// @Router			/admin/upload [post]
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		common.Fail(c, err.Error())
		return
	}
	if file.Size > 10*1024*1024 {
		common.Fail(c, "文件大小超出")
		return
	}
	types := file.Header.Get("Content-Type")
	if types != "application/vnd.android.package-archive" {
		common.Fail(c, "上传类型错误")
		return
	}

	newName := common.CreateMd5Str(common.RandStr(6), file.Filename) + ".apk"
	e := c.SaveUploadedFile(file, "./storage/file/"+newName)
	if e != nil {
		common.Fail(c, "上传错误")
		return
	}
	var res struct {
		File string
	}
	res.File = newName
	common.Success(c, res)
}

// Release
//
//	@Summary		发布app
//	@Description
//	@Tags			管理员
//	@Accept			application/json
//	@Produce		application/json
//
// @Security 		admintoken
// @Success			200			{object}	utils.Res
// @Router			/admin/release [post]
func Release(c *gin.Context) {
	params := request.ReleaseParams{}
	if e := c.ShouldBindJSON(&params); e != nil {
		bootstrap.Log.Info("发布失败", zap.Any("params", params))
		common.Fail(c, "参数错误")
		return
	}
	re := services2.Nul{}
	e := re.Release(c, params)
	if e != nil {
		common.Fail(c, e.Error())
		return
	}
	common.Ok(c)
}
