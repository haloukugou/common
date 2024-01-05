package main

import (
	"database/sql"
	"dj/bootstrap"
	"dj/config"
	"dj/route"
)

func initProjectSetting() {
	// 加载配置项
	config.InitConfig()
}

// @title 简单的社交系统
// @version 1.0
// @description dj的go接口文档
// @host 127.0.0.1:8888
// @securityDefinitions.apikey  Token
// @in                          header
// @name                        token
func main() {
	initProjectSetting()

	// 拿数据库连接
	bootstrap.Db = bootstrap.Connect()
	db, e := bootstrap.Db.DB()
	if e != nil {
		panic(e)
	}
	// 关闭数据库
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// 拿redis连接
	bootstrap.Redis = bootstrap.RedisConnect()

	// log日志
	bootstrap.Log = bootstrap.InitializeLog()

	// 初始化验证器
	bootstrap.InitializeValidator()

	s := route.Router()
	err := s.Run(config.Config.Server.Port)
	if err != nil {
		panic("服务启动错误")
	}
}
