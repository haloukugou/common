package logic

import (
	"dj/bootstrap"
	"dj/common"
	"dj/constants"
	"dj/model"
	"dj/request"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type Nul struct {
}

type TokenResult struct {
	Token string `json:"token"`
}

type UserInfo struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

// HandleRegister 注册
func (n *Nul) HandleRegister(ctx *gin.Context, params request.RegisterParams) error {
	switch params.TypeString {
	case "account":
		// 简易判断
		if params.Account == "" {
			return fmt.Errorf("账号不允许为空")
		}
		accountLen := len(params.Account)
		if accountLen > 16 || accountLen < 6 {
			return fmt.Errorf("账号长度有误")
		}
		if params.Password == "" {
			return fmt.Errorf("密码不允许为空")
		}
		if params.Rpassword == "" {
			return fmt.Errorf("请确认密码")
		}
		if params.Password != params.Rpassword {
			return fmt.Errorf("两次输入密码不一致")
		}
		return accountRegister(params)
	case "mail":
		if params.Mail == "" {
			return fmt.Errorf("请填写邮箱")
		}
		if false == common.VerifyEmailFormat(params.Mail) {
			return fmt.Errorf("邮箱格式错误")
		}
		if params.Code == "" {
			return fmt.Errorf("请填写邮箱验证码")
		}
		return mailRegister(ctx, params)
	default:
		return fmt.Errorf("注册类型错误")
	}
}

// 账号注册
func accountRegister(params request.RegisterParams) error {
	// 查询数据库是否有相同账号
	user := new(model.User)
	bootstrap.Db.Where("account =?", params.Account).Where("source = ?", params.Source).First(&user)
	if user.Id > 0 {
		return fmt.Errorf("账号已存在")
	}
	// 组装需要添加的数据
	salt := common.RandStr(6)
	user.Account = params.Account
	user.Name = params.Account
	user.Salt = salt
	user.Password = common.CreateMd5Str(salt, params.Password)
	if user.Password == "" {
		return fmt.Errorf("密码生成错误")
	}
	user.CreatedAt = time.Now()
	user.Source = params.Source

	// 添加数据
	bootstrap.Db.Create(user)
	return nil
}

func mailRegister(ctx *gin.Context, params request.RegisterParams) error {
	key := constants.SendMailTypeLock + "register:" + params.Mail
	re, e := bootstrap.Redis.Get(ctx, key).Result()
	if e != nil {
		return e
	}
	if re != params.Code {
		return fmt.Errorf("验证码错误")
	}
	// 查询数据库是否有相同邮箱
	user := new(model.User)
	bootstrap.Db.Where("mail =?", params.Mail).Where("source=?", params.Source).First(&user)
	if user.Id > 0 {
		return fmt.Errorf("邮箱已存在")
	}
	// 组装需要添加的数据
	user.Mail = params.Mail
	user.Name = params.Mail
	user.Account = params.Mail
	user.CreatedAt = time.Now()
	user.Source = params.Source

	// 添加数据
	e = bootstrap.Db.Create(user).Error
	if e != nil {
		return e
	}

	// 清空redis缓存的code
	bootstrap.Redis.Del(ctx, key)
	return nil
}

// HandleLogin 登录
func (t *TokenResult) HandleLogin(c *gin.Context, params request.LoginParams) error {
	switch params.TypeString {
	case "account":
		// 查询数据库是否有相同账号
		user := new(model.User)
		bootstrap.Db.Where("account =?", params.Account).Where("source = ?", params.Source).First(&user)
		if user.Id <= 0 {
			return fmt.Errorf("用户不存在,请先注册")
		}
		loginPwd := common.CreateMd5Str(user.Salt, params.Password)
		if loginPwd != user.Password {
			return fmt.Errorf("密码错误")
		}
		// 生成token
		t.Token = common.CreateToken(user.Account, user.Salt)

		// 放入redis
		rErr := loginInfoToRedis(c, user, t.Token)
		if rErr != nil {
			return rErr
		}
		return nil
	case "mail":
		if false == common.VerifyEmailFormat(params.Account) {
			return fmt.Errorf("邮箱格式错误")
		}
		if params.Code == "" {
			return fmt.Errorf("请输入邮箱验证码")
		}
		// 查询账号是否存在
		user := new(model.User)
		bootstrap.Db.Where("mail=?", params.Account).Where("source = ?", params.Source).First(&user)
		if user.Id <= 0 {
			return fmt.Errorf("该邮箱尚未注册")
		}

		// 验证code
		key := constants.SendMailTypeLock + "login:" + params.Account
		re, e := bootstrap.Redis.Get(c, key).Result()
		if e != nil {
			return e
		}
		if re != params.Code {
			return fmt.Errorf("验证码错误")
		}
		t.Token = common.CreateToken(params.Account, params.Code)

		// redis存信息
		rErr := loginInfoToRedis(c, user, t.Token)
		if rErr != nil {
			return rErr
		}
		return nil
	default:
		return fmt.Errorf("登录方式错误")
	}
}

// loginInfoToRedis 登录信息放入redis
func loginInfoToRedis(c *gin.Context, user *model.User, token string) error {
	redisData := make(map[string]string, 5)
	redisData["id"] = strconv.FormatUint(user.Id, 10)
	redisData["name"] = user.Name
	redisData["account"] = user.Account
	redisData["created_at"] = user.CreatedAt.Format("2006-01-02 15:04:05")
	redisData["title"] = user.Title
	data, uErr := json.Marshal(redisData)
	if uErr != nil {
		return uErr
	}

	err := bootstrap.Redis.Set(
		c,
		constants.LoginKey+token,
		string(data),
		time.Duration(constants.RedisTtl)*time.Second,
	).Err()
	if err != nil {
		return err
	}
	return nil
}

// HandleLoginOut 退出登录
func (n *Nul) HandleLoginOut(ctx *gin.Context) error {
	token := ctx.GetHeader("token")
	bootstrap.Redis.Del(ctx, constants.LoginKey+token)
	return nil
}

// UserInfo 用户详情
func (u *UserInfo) UserInfo(ctx *gin.Context) error {
	// 拿缓存信息
	info, e := redisUserInfo(ctx)
	if e != nil {
		return e
	}
	// 转类型
	i := make(map[string]string)
	jErr := json.Unmarshal([]byte(info), &i)
	if jErr != nil {
		bootstrap.Log.Warn("查询用户信息,类型转换错误")
		return e
	}
	idInt64, _ := strconv.ParseInt(i["id"], 10, 64)
	u.Id = uint64(idInt64)
	u.Name = i["name"]
	u.Title = i["title"]
	return nil
}

// EditInfo 编辑详情
func (n *Nul) EditInfo(ctx *gin.Context, params request.EditInfo) error {
	// 缓存的用户信息
	info, e := redisUserInfo(ctx)
	if e != nil {
		return e
	}
	// 转成map
	i := make(map[string]string)
	jErr := json.Unmarshal([]byte(info), &i)
	if jErr != nil {
		bootstrap.Log.Warn("查询用户信息,类型转换错误")
		return e
	}
	idInt64, _ := strconv.ParseInt(i["id"], 10, 64)
	userId := uint64(idInt64)
	// 查询数据库用户信息
	user := new(model.User)
	bootstrap.Db.Where("id = ?", userId).First(&user)

	// 编辑用户详情
	d := make(map[string]interface{})
	d["name"] = params.Name
	d["title"] = params.Title
	dbErr := bootstrap.Db.Model(&user).Updates(d).Error
	if dbErr != nil {
		bootstrap.Log.Error("编辑用户信息失败,msg=" + dbErr.Error())
		return dbErr
	}

	// 更新redis里面的信息
	redisData := make(map[string]string, 5)
	redisData["id"] = strconv.FormatUint(user.Id, 10)
	redisData["name"] = user.Name
	redisData["account"] = user.Account
	redisData["created_at"] = user.CreatedAt.Format("2006-01-02 15:04:05")
	redisData["title"] = user.Title
	data, uErr := json.Marshal(redisData)
	if uErr != nil {
		return uErr
	}

	err := bootstrap.Redis.Set(
		ctx,
		constants.LoginKey+ctx.GetHeader("token"),
		string(data),
		constants.RedisTtl*time.Second,
	).Err()
	if err != nil {
		return err
	}
	return nil
}

func redisUserInfo(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader("token")
	info, e := bootstrap.Redis.Get(ctx, constants.LoginKey+token).Result()
	if e != nil {
		bootstrap.Log.Warn("查询用户信息,缓存信息不存在", zap.Any("param", token))
		return "", e
	}
	return info, nil
}

// EditPwd 修改密码
func (n *Nul) EditPwd(ctx *gin.Context, params request.EditPwd) error {
	// 缓存的用户信息
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

	// 查数据库用户信息
	user := new(model.User)
	bootstrap.Db.Where("id = ?", userId).First(&user)
	if user.Salt != "" && user.Password != "" {
		if params.Password == "" {
			return fmt.Errorf("请输入旧密码")
		}
		// 对比密码
		p := common.CreateMd5Str(user.Salt, params.Password)
		if p != user.Password {
			return fmt.Errorf("旧密码错误,请重新输入")
		}
	}
	if params.NewPassword != params.NewRpassword {
		return fmt.Errorf("新密码两次输入不一致")
	}

	// 新密码和盐
	salt := common.RandStr(6)
	np := common.CreateMd5Str(salt, params.NewPassword)

	// 编辑密码
	dbErr := bootstrap.Db.Model(&user).Updates(model.User{Salt: salt, Password: np}).Error
	if dbErr != nil {
		bootstrap.Log.Error("编辑密码信息失败,msg=" + dbErr.Error())
		return dbErr
	}

	// 登录信息失效
	bootstrap.Redis.Del(ctx, constants.LoginKey+ctx.GetHeader("token"))

	return nil
}

func (n *Nul) BindMail(ctx *gin.Context, params request.BindMail) error {
	user := new(model.User)
	// 查询邮件是否被绑定
	bootstrap.Db.Where("mail=?", params.Mail).Where("source = ?", params.Source).First(&user)
	if user.Id > 0 {
		return fmt.Errorf("邮箱已被绑定")
	}

	// 验证code是否正确
	token := ctx.GetHeader("token")
	key := constants.SendMailTypeLock + "bind:" + token
	rInfo, _ := bootstrap.Redis.Get(ctx, key).Result()
	if rInfo != params.Code {
		return fmt.Errorf("验证码错误")
	}

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
	bootstrap.Db.Where("id=?", userId).First(&user)
	rErr := bootstrap.Db.Model(&user).Updates(model.User{Mail: params.Mail}).Error
	if rErr != nil {
		return rErr
	}

	// 清空验证码信息
	bootstrap.Redis.Del(ctx, key)
	return nil
}

func (n *Nul) RetrievePwd(ctx *gin.Context, params request.RetrievePwd) error {
	// 查询账号是否存在
	user := new(model.User)
	bootstrap.Db.Where("mail=?", params.Mail).Where("source = ?", params.Source).First(&user)
	if user.Id <= 0 {
		return fmt.Errorf("账号信息不存在")
	}
	// 验证code
	key := constants.SendMailTypeLock + "pwd:" + params.Mail
	re, e := bootstrap.Redis.Get(ctx, key).Result()
	if e != nil {
		return e
	}
	if re != params.Code {
		return fmt.Errorf("邮箱验证码错误")
	}
	if params.NewPassword != params.NewRpassword {
		return fmt.Errorf("两次输入的密码不一致")
	}

	// 更新账号密码
	salt := common.RandStr(6)
	pwd := common.CreateMd5Str(salt, params.NewPassword)
	dErr := bootstrap.Db.Model(&user).Updates(model.User{Salt: salt, Password: pwd}).Error
	if dErr != nil {
		return dErr
	}
	return nil
}
