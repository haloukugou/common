package middleware

import (
	"dj/bootstrap"
	"dj/common"
	"dj/constants"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("admintoken")
		if token == "" {
			common.Fail(c, "登录后再进行操作")
			c.Abort()
			return
		}
		redis := bootstrap.RedisConnect()
		key := constants.AdminLoginKey + token
		info, _ := redis.Get(c, key).Result()
		bootstrap.Log.Info("后台鉴权", zap.Any("info", info))
		if info == "" {
			common.Fail(c, "登录信息已过期,请重新登录")
			c.Abort()
			return
		}
		redis.Expire(c, key, constants.RedisTtl*time.Second)
		c.Next()

	}
}
