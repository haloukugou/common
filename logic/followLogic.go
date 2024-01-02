package logic

import (
	"dj/bootstrap"
	"dj/model"
	"dj/request"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
)

type FollowList struct {
	Page     uint64     `json:"page"`
	PageSize uint64     `json:"pageSize"`
	MaxPage  uint64     `json:"maxPage"`
	Total    int64      `json:"total"`
	List     []ListData `json:"list"`
}

type ListData struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Mail string `json:"mail"`
}

func (n *Nul) Follow(ctx *gin.Context, params request.Follow) error {
	info, e := redisUserInfo(ctx)
	if e != nil {
		return e
	}
	i := make(map[string]string)
	jErr := json.Unmarshal([]byte(info), &i)
	if jErr != nil {
		bootstrap.Log.Warn("查询用户信息,类型转换错误")
		return jErr
	}
	idInt64, _ := strconv.ParseInt(i["id"], 10, 64)
	userId := uint64(idInt64)
	loginUser := new(model.User)
	bootstrap.Db.Where("id=?", userId).First(&loginUser)

	user := new(model.User)
	bootstrap.Db.Where("id=?", params.FollowedPerson).Where("source = ?", loginUser.Source).First(&user)
	if user.Id <= 0 {
		return fmt.Errorf("关注用户不存在")
	}
	if userId == user.Id {
		return fmt.Errorf("不可以关注自己")
	}
	follow := new(model.Follow)
	// 查询是否存在关注信息
	bootstrap.Db.Where("follow_person=?", userId).Where("followed_person=?", user.Id).First(&follow)
	if follow.Id > 0 {
		return fmt.Errorf("不可重复关注")
	}

	follow.FollowPerson = userId
	follow.FollowedPerson = user.Id
	dErr := bootstrap.Db.Create(follow).Error
	if dErr != nil {
		return dErr
	}
	return nil
}

func (n *Nul) CancelFollow(ctx *gin.Context, params request.Follow) error {
	info, e := redisUserInfo(ctx)
	if e != nil {
		return e
	}
	i := make(map[string]string)
	jErr := json.Unmarshal([]byte(info), &i)
	if jErr != nil {
		bootstrap.Log.Warn("查询用户信息,类型转换错误")
		return jErr
	}
	idInt64, _ := strconv.ParseInt(i["id"], 10, 64)
	userId := uint64(idInt64)

	follow := new(model.Follow)
	// 查询是否存在关注信息
	bootstrap.Db.Where("follow_person=?", userId).Where("followed_person=?", params.FollowedPerson).First(&follow)
	if follow.Id <= 0 {
		return fmt.Errorf("关注信息不存在")
	}
	dErr := bootstrap.Db.Delete(&follow).Error
	if dErr != nil {
		return dErr
	}
	return nil
}

func (f *FollowList) FollowList(ctx *gin.Context, params request.FollowList) error {
	info, e := redisUserInfo(ctx)
	if e != nil {
		return e
	}
	i := make(map[string]string)
	jErr := json.Unmarshal([]byte(info), &i)
	if jErr != nil {
		bootstrap.Log.Warn("查询用户信息,类型转换错误")
		return jErr
	}
	idInt64, _ := strconv.ParseInt(i["id"], 10, 64)
	userId := uint64(idInt64)

	f.Page = params.Page
	f.PageSize = params.PageSize

	fo := new(model.Follow)
	// 查询总数
	bootstrap.Db.Model(&fo).Where("follow_person=?", userId).Count(&f.Total)

	// 偏移量
	a := float64(f.Total) / float64(params.PageSize)
	f.MaxPage = uint64(math.Ceil(a))
	if params.Page > f.MaxPage {
		return nil
	}

	list := make([]ListData, 0)
	offset := (params.Page - 1) * params.PageSize
	// 查询数据
	dErr := bootstrap.Db.Table("follow").Select("user.id,user.name,user.mail").Joins("left join user on follow.follow_person=user.id").Where("follow.follow_person=?", userId).Order("follow.created_at desc").Offset(int(offset)).Limit(int(params.PageSize)).Scan(&list).Error
	if dErr != nil {
		return dErr
	}
	f.List = list
	return nil
}

func (f *FollowList) FansList(ctx *gin.Context, params request.FollowList) error {
	info, e := redisUserInfo(ctx)
	if e != nil {
		return e
	}
	i := make(map[string]string)
	jErr := json.Unmarshal([]byte(info), &i)
	if jErr != nil {
		bootstrap.Log.Warn("查询用户信息,类型转换错误")
		return jErr
	}
	idInt64, _ := strconv.ParseInt(i["id"], 10, 64)
	userId := uint64(idInt64)
	f.Page = params.Page
	f.PageSize = params.PageSize

	fo := new(model.Follow)
	// 查询总数
	bootstrap.Db.Model(&fo).Where("followed_person=?", userId).Count(&f.Total)

	// 偏移量
	a := float64(f.Total) / float64(params.PageSize)
	f.MaxPage = uint64(math.Ceil(a))
	if params.Page > f.MaxPage {
		return nil
	}

	list := make([]ListData, 0)
	offset := (params.Page - 1) * params.PageSize
	// 查询数据
	dErr := bootstrap.Db.Table("follow").Select("user.id,user.name,user.mail").Joins("left join user on follow.follow_person=user.id").Where("follow.followed_person=?", userId).Order("follow.created_at desc").Offset(int(offset)).Limit(int(params.PageSize)).Scan(&list).Error
	if dErr != nil {
		return dErr
	}
	f.List = list
	return nil
}
