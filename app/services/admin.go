package services

import (
	"dj/app/common"
	models2 "dj/app/models"
	"dj/bootstrap"
	"dj/constants"
	"dj/request"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
	"strings"
	"time"
)

type LoginResult struct {
	Token string `json:"token"`
}

type ApkList struct {
	Page     uint64        `json:"page"`
	PageSize uint64        `json:"pageSize"`
	MaxPage  uint64        `json:"maxPage"`
	Total    int64         `json:"total"`
	List     []ApkListData `json:"list"`
}

type ApkListData struct {
	Id      uint   `json:"id"`
	Version string `json:"version"`
	File    string `json:"file"`
	IsForce int    `json:"is_force"`
}

func (l *LoginResult) Login(c *gin.Context, params request.AdminLoginParams) error {
	admin := new(models2.Admin)
	bootstrap.Db.Where("account = ?", params.Account).First(&admin)
	if admin.Id <= 0 {
		return fmt.Errorf("管理账号不存")
	}
	if params.Password != admin.Password {
		return fmt.Errorf("密码错误")
	}
	randStr := common.RandStr(6)
	l.Token = common.CreateToken(admin.Account, randStr)

	// token存redis
	data := make(map[string]interface{}, 2)
	data["id"] = admin.Id
	data["name"] = admin.Name
	str, err := json.Marshal(data)
	if err != nil {
		return err
	}
	rErr := bootstrap.Redis.Set(
		c,
		constants.AdminLoginKey+l.Token,
		string(str),
		time.Duration(constants.RedisTtl)*time.Second,
	).Err()
	if rErr != nil {
		return rErr
	}
	return nil
}

func (a *ApkList) ApkList(c *gin.Context, params request.ApkListParams) error {
	a.Page = params.Page
	a.PageSize = params.PageSize

	l := new(models2.Apk)

	bootstrap.Db.Model(&l).Where("id > 0").Count(&a.Total)

	o := float64(a.Total) / float64(params.PageSize)
	a.MaxPage = uint64(math.Ceil(o))
	if params.Page > a.MaxPage {
		return nil
	}

	list := make([]ApkListData, 0)
	offset := (params.Page - 1) * params.PageSize
	err := bootstrap.Db.Table("apk").Select("id,file,version,is_force").Where("id >0").Order("id desc").Offset(int(offset)).Limit(int(a.PageSize)).Scan(&list).Error
	if err != nil {
		return err
	}
	a.List = list
	return nil
}

func (n *Nul) Release(c *gin.Context, params request.ReleaseParams) error {
	// 判断版本号格式
	version := strings.Split(params.Version, ".")
	l := len(version)
	if 3 != l {
		return fmt.Errorf("版本号格式错误-1")
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
	// 查询上个版本的发布信息
	apk := new(models2.Apk)
	bootstrap.Db.Where("id>0").Order("id desc").First(&apk)
	if apk.Id > 0 {
		oldVersion := strings.Split(apk.Version, ".")
		status := false
		for i, n := 0, 3; i < n; i++ {
			if version[i] > oldVersion[i] {
				status = true
				break
			}
		}
		if status == false {
			return fmt.Errorf("当前版本必须高于上个版本号")
		}
	}
	newApk := new(models2.Apk)
	newApk.File = params.File
	newApk.Version = params.Version
	newApk.IsForce = params.IsForce
	bootstrap.Db.Create(newApk)
	return nil
}
