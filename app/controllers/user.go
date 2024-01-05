package controllers

import (
	"dj/app/common"
	"dj/app/request"
	services2 "dj/app/services"
	"dj/bootstrap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Register
//
//	@Summary		用户注册接口
//	@Description
//	@Tags			用户
//	@Accept			application/json
//	@Produce		application/json
//
// @Param			params		body		request.RegisterParams		false		"请求参数"
// @Success			200			{object}	utils.Res
// @Router			/user/register [post]
func Register(c *gin.Context) {
	// 获取参数
	params := request.RegisterParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		bootstrap.Log.Warn("用户注册,请求参数错误", zap.Any("params", params))
		common.Fail(c, "参数错误-1")
		return
	}
	sourceArr := []int{1, 2, 3}
	if false == common.IsContainInt(sourceArr, params.Source) {
		common.Fail(c, "来源信息不存在")
		return
	}

	reg := &services2.Nul{}
	rErr := reg.HandleRegister(c, params)
	if rErr != nil {
		bootstrap.Log.Warn("用户注册,失败", zap.Any("error", rErr.Error()))
		common.Fail(c, rErr.Error())
		return
	}

	common.Ok(c)
}

// Login
//
//	@Summary		用户登录接口
//	@Description
//	@Tags			用户
//	@Accept			application/json
//	@Produce		application/json
//
// @Param			params		body		request.LoginParams		false		"请求参数"
// @Success			200			{object}	utils.Res
// @Router			/user/login [post]
func Login(c *gin.Context) {
	// 获取参数
	params := request.LoginParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		bootstrap.Log.Warn("用户登录,参数错误", zap.Any("params", params))
		common.Fail(c, "参数错误")
		return
	}
	sourceArr := []int{1, 2, 3}
	if false == common.IsContainInt(sourceArr, params.Source) {
		common.Fail(c, "来源信息不存在")
		return
	}
	tokenResult := &services2.TokenResult{}
	rErr := tokenResult.HandleLogin(c, params)
	if rErr != nil {
		bootstrap.Log.Error("用户登录,失败", zap.Any("error", rErr.Error()))
		common.Fail(c, rErr.Error())
		return
	}
	common.Success(c, tokenResult)
}

// LoginOut
//
//	@Summary		退出登录接口
//	@Description
//	@Tags			用户
//	@Accept			application/json
//	@Produce		application/json
//
// @Security 		Token
// @Success			200			{object}	utils.Res
// @Router			/user/loginOut [post]
func LoginOut(c *gin.Context) {
	r := &services2.Nul{}
	e := r.HandleLoginOut(c)
	if e != nil {
		bootstrap.Log.Error("用户退出登录,失败", zap.Any("error", e.Error()))
		common.Fail(c, e.Error())
		return
	}
	common.Ok(c)
}

// UserInfo
//
//	@Summary		用户详情
//	@Description
//	@Tags			用户
//	@Accept			application/json
//	@Produce		application/json
//
// @Security 		Token
// @Success			200			{object}	utils.Res
// @Router			/user/userInfo [post]
func UserInfo(c *gin.Context) {
	r := &services2.UserInfo{}
	e := r.UserInfo(c)
	if e != nil {
		common.Fail(c, e.Error())
		return
	}
	common.Success(c, r)
}

// EditInfo
//
//	@Summary		编辑信息
//	@Description
//	@Tags			用户
//	@Accept			application/json
//	@Produce		application/json
//
// @Security 		Token
// @Param			params		body		request.EditInfo		false		"请求参数"
// @Success			200			{object}	utils.Res
// @Router			/user/editInfo [post]
func EditInfo(c *gin.Context) {
	params := request.EditInfo{}
	if e := c.ShouldBindJSON(&params); e != nil {
		bootstrap.Log.Info("编辑信息,参数错误", zap.Any("params", params))
		common.Fail(c, "参数错误")
		return
	}
	r := &services2.Nul{}
	e := r.EditInfo(c, params)
	if e != nil {
		common.Fail(c, e.Error())
		return
	}
	common.Success(c, r)
}

// EditPwd
//
//	@Summary		修改密码
//	@Description
//	@Tags			用户
//	@Accept			application/json
//	@Produce		application/json
//
// @Security 		Token
// @Param			params		body		request.EditPwd		false		"请求参数"
// @Success			200			{object}	utils.Res
// @Router			/user/editPwd [post]
func EditPwd(c *gin.Context) {
	params := request.EditPwd{}
	if e := c.ShouldBindJSON(&params); e != nil {
		bootstrap.Log.Info("编辑密码,参数错误", zap.Any("params", params))
		common.Fail(c, "参数错误")
		return
	}
	r := &services2.Nul{}
	e := r.EditPwd(c, params)
	if e != nil {
		common.Fail(c, e.Error())
		return
	}
	common.Success(c, r)
}

// BindMail
//
//	@Summary		绑定邮箱
//	@Description
//	@Tags			用户
//	@Accept			application/json
//	@Produce		application/json
//
// @Security 		Token
// @Param			params		body		request.BindMail		false		"请求参数"
// @Success			200			{object}	utils.Res
// @Router			/user/bindMail [post]
func BindMail(c *gin.Context) {
	mail := request.BindMail{}
	if e := c.ShouldBindJSON(&mail); e != nil {
		bootstrap.Log.Info("绑定邮箱,参数错误", zap.Any("params", mail))
		common.Fail(c, "参数错误")
		return
	}
	sourceArr := []int{1, 2, 3}
	if false == common.IsContainInt(sourceArr, mail.Source) {
		common.Fail(c, "来源信息不存在")
		return
	}
	r := &services2.Nul{}
	err := r.BindMail(c, mail)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}
	common.Success(c, r)
}

// RetrievePwd
//
//	@Summary		找回密码
//	@Description
//	@Tags			用户
//	@Accept			application/json
//	@Produce		application/json
//
// @Security 		Token
// @Param			params		body		request.RetrievePwd		false		"请求参数"
// @Success			200			{object}	utils.Res
// @Router			/user/retrievePwd [post]
func RetrievePwd(c *gin.Context) {
	params := request.RetrievePwd{}
	if e := c.ShouldBindJSON(&params); e != nil {
		bootstrap.Log.Info("找回密码,参数错误", zap.Any("params", params))
		common.Fail(c, "参数错误")
		return
	}
	sourceArr := []int{1, 2, 3}
	if false == common.IsContainInt(sourceArr, params.Source) {
		common.Fail(c, "来源信息不存在")
		return
	}
	r := services2.Nul{}
	err := r.RetrievePwd(c, params)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}
	common.Success(c, r)
}

// Follow
//
//	@Summary		关注
//	@Description
//	@Tags			用户
//	@Accept			application/json
//	@Produce		application/json
//
// @Security 		Token
// @Param			params		body		request.Follow		false		"请求参数"
// @Success			200			{object}	utils.Res
// @Router			/user/follow [post]
func Follow(c *gin.Context) {
	params := request.Follow{}
	if e := c.ShouldBindJSON(&params); e != nil {
		bootstrap.Log.Info("关注用户,参数错误", zap.Any("params", params))
		common.Fail(c, "参数错误")
		return
	}
	r := services2.Nul{}
	err := r.Follow(c, params)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}
	common.Success(c, r)
}

// CancelFollow
//
//	@Summary		取消关注
//	@Description
//	@Tags			用户
//	@Accept			application/json
//	@Produce		application/json
//
// @Security 		Token
// @Param			params		body		request.Follow		false		"请求参数"
// @Success			200			{object}	utils.Res
// @Router			/user/cancelFollow [post]
func CancelFollow(c *gin.Context) {
	params := request.Follow{}
	if e := c.ShouldBindJSON(&params); e != nil {
		bootstrap.Log.Info("取消关注用户,参数错误", zap.Any("params", params))
		common.Fail(c, "参数错误")
		return
	}
	r := services2.Nul{}
	err := r.CancelFollow(c, params)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}
	common.Success(c, r)
}

// FollowList
//
//	@Summary		关注列表
//	@Description
//	@Tags			用户
//	@Accept			application/json
//	@Produce		application/json
//
// @Security 		Token
// @Param			params		body		request.FollowList		false		"请求参数"
// @Success			200			{object}	utils.Res
// @Router			/user/followList [post]
func FollowList(c *gin.Context) {
	params := request.FollowList{}
	if e := c.ShouldBindJSON(&params); e != nil {
		bootstrap.Log.Info("查询关注列表,参数错误-1", zap.Any("params", params))
		common.Fail(c, "参数错误-1")
		return
	}
	f := services2.FollowList{}
	err := f.FollowList(c, params)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}
	common.Success(c, f)
}

// FansList
//
//	@Summary		粉丝列表
//	@Description
//	@Tags			用户
//	@Accept			application/json
//	@Produce		application/json
//
// @Security 		Token
// @Param			params		body		request.FollowList		false		"请求参数"
// @Success			200			{object}	utils.Res
// @Router			/user/fansList [post]
func FansList(c *gin.Context) {
	params := request.FollowList{}
	if e := c.ShouldBindJSON(&params); e != nil {
		bootstrap.Log.Info("查询粉丝列表,参数错误-1", zap.Any("params", params))
		common.Fail(c, "参数错误-1")
		return
	}
	f := services2.FollowList{}
	err := f.FansList(c, params)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}
	common.Success(c, f)
}
