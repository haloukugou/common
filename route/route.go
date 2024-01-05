package route

import (
	"dj/app/controllers/client"
	"dj/app/middleware"
	_ "dj/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	route := gin.Default()
	route.Use(middleware.Cors())

	// swag文档路由
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	c := route.Group("/c")
	{
		user := c.Group("/user")
		{
			user.POST("/register", client.Register)
		}
		message := c.Group("/message")
		{
			message.POST("/send", client.Send)
		}
	}

	//user := route.Group("/user")
	//{
	//	// 注册
	//	user.POST("/register", controllers.Register)
	//	// 登录
	//	user.POST("/login", controllers.Login)
	//	// 退出登录
	//	user.POST("/loginOut", middleware.Auth(), controllers.LoginOut)
	//	// 用户信息
	//	user.POST("/userInfo", middleware.Auth(), controllers.UserInfo)
	//	// 编辑信息
	//	user.POST("/editInfo", middleware.Auth(), controllers.EditInfo)
	//	// 修改密码
	//	user.POST("/editPwd", middleware.Auth(), controllers.EditPwd)
	//	// 绑定邮箱
	//	user.POST("/bindMail", middleware.Auth(), controllers.BindMail)
	//	// 找回密码
	//	user.POST("/retrievePwd", controllers.RetrievePwd)
	//	// 关注
	//	user.POST("/follow", middleware.Auth(), controllers.Follow)
	//	// 取消关注
	//	user.POST("/cancelFollow", middleware.Auth(), controllers.CancelFollow)
	//	// 关注列表
	//	user.POST("/followList", middleware.Auth(), controllers.FollowList)
	//	// 粉丝列表
	//	user.POST("/fansList", middleware.Auth(), controllers.FansList)
	//}
	//service := route.Group("/service")
	//{
	//	// 发送邮件
	//	service.POST("/sendMail", controllers.SendMail)
	//}
	//
	//replace := route.Group("/replace")
	//{
	//	// 最新版本信息
	//	replace.POST("/latest", controllers.Latest)
	//}
	//
	//admin := route.Group("/admin")
	//{
	//	// 管理员登录
	//	admin.POST("/login", controllers.AdminLogin)
	//	// 查看apk列表
	//	admin.POST("/apkList", middleware.AdminAuth(), controllers.ApkList)
	//	// 上传文件
	//	admin.POST("/upload", middleware.AdminAuth(), controllers.Upload)
	//	// 发布app
	//	admin.POST("/release", middleware.AdminAuth(), controllers.Release)
	//}
	return route
}
