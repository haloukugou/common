package route

import (
	"dj/controller"
	_ "dj/docs"
	"dj/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	route := gin.Default()
	route.Use(middleware.Cors())

	// swag文档路由
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := route.Group("/user")
	{
		// 注册
		user.POST("/register", controller.Register)
		// 登录
		user.POST("/login", controller.Login)
		// 退出登录
		user.POST("/loginOut", middleware.Auth(), controller.LoginOut)
		// 用户信息
		user.POST("/userInfo", middleware.Auth(), controller.UserInfo)
		// 编辑信息
		user.POST("/editInfo", middleware.Auth(), controller.EditInfo)
		// 修改密码
		user.POST("/editPwd", middleware.Auth(), controller.EditPwd)
		// 绑定邮箱
		user.POST("/bindMail", middleware.Auth(), controller.BindMail)
		// 找回密码
		user.POST("/retrievePwd", controller.RetrievePwd)
		// 关注
		user.POST("/follow", middleware.Auth(), controller.Follow)
		// 取消关注
		user.POST("/cancelFollow", middleware.Auth(), controller.CancelFollow)
		// 关注列表
		user.POST("/followList", middleware.Auth(), controller.FollowList)
		// 粉丝列表
		user.POST("/fansList", middleware.Auth(), controller.FansList)
	}
	service := route.Group("/service")
	{
		// 发送邮件
		service.POST("/sendMail", controller.SendMail)
	}

	replace := route.Group("/replace")
	{
		// 最新版本信息
		replace.POST("/latest", controller.Latest)
	}

	admin := route.Group("/admin")
	{
		// 管理员登录
		admin.POST("/login", controller.AdminLogin)
		// 查看apk列表
		admin.POST("/apkList", middleware.AdminAuth(), controller.ApkList)
		// 上传文件
		admin.POST("/upload", middleware.AdminAuth(), controller.Upload)
		// 发布app
		admin.POST("/release", middleware.AdminAuth(), controller.Release)
	}
	return route
}
