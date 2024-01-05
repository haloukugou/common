package utils

import (
	"dj/bootstrap"
	"dj/constants"
	"github.com/gin-gonic/gin"
	"time"
)

func Set(ctx *gin.Context, key, value string) bool {
	err := bootstrap.RedisConnect().Set(ctx, key, value, time.Second*constants.CodeTtl).Err()
	if err != nil {
		return false
	}
	return true
}
