package logic

import (
	"dj/bootstrap"
	"dj/model"
	"dj/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Latest struct {
	IsForce       int    `json:"is_force"`
	LatestVersion string `json:"latest_version"`
	NeedReplace   bool   `json:"need_replace"`
}

func (l *Latest) LatestInfo(c *gin.Context, params request.LatestParams) error {
	// 判断版本号格式
	version := strings.Split(params.ClientVersion, ".")
	length := len(version)
	if 3 != length {
		return fmt.Errorf("版本号格式错误")
	}
	for _, v := range version {
		if v == "" {
			return fmt.Errorf("版本号格式错误-2")
		}
		i, nErr := strconv.Atoi(v)
		if nErr != nil {
			return fmt.Errorf("版本号格式错误-3")
		}
		if i < 0 {
			return fmt.Errorf("版本号格式错误-4")
		}
	}

	// 查询最新版本信息
	apk := new(model.Apk)
	e := bootstrap.Db.Where("id>0").Order("id desc").First(&apk).Error
	if e != nil {
		return e
	}
	fmt.Println(apk)
	newVersion := strings.Split(apk.Version, ".")
	for i, n := 0, 3; i < n; i++ {
		if version[i] < newVersion[i] {
			l.NeedReplace = true
			break
		}
	}
	l.LatestVersion = apk.Version
	l.IsForce = apk.IsForce
	return nil
}
