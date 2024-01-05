package test

//
//import (
//	"dj/app/common"
//	"dj/bootstrap"
//	"dj/config"
//	"dj/request"
//	"dj/route"
//	"encoding/json"
//	"github.com/gin-gonic/gin"
//	"go.uber.org/zap"
//	"os"
//	"testing"
//)
//
//var r *gin.Engine
//
//func Init() {
//	_ = os.Chdir("../")
//	// 加载配置项
//	config.InitConfig()
//
//	// 拿数据库连接
//	bootstrap.Db = bootstrap.Connect()
//	_, e := bootstrap.Db.DB()
//	if e != nil {
//		panic(e)
//	}
//
//	// 拿redis连接
//	bootstrap.Redis = bootstrap.RedisConnect()
//
//	// log日志
//	bootstrap.Log = bootstrap.InitializeLog()
//
//	r = route.Router()
//}
//
//func TestRegister(t *testing.T) {
//	Init()
//	str := common.RandStr(6)
//	p := request.RegisterParams{
//		TypeString: "account",
//		Account:    str,
//		Password:   str,
//		Rpassword:  str,
//		Mail:       "18280189749@163.com",
//		Code:       "login",
//	}
//	mar, _ := json.Marshal(p)
//	MarshalParamsAndRequest(mar, t, "/user/register", "")
//}
//
//func TestLogin(t *testing.T) {
//	Init()
//	p := request.LoginParams{
//		Account:    "dongjie",
//		Password:   "123456",
//		TypeString: "account",
//		Code:       "",
//	}
//	jsonStr, _ := json.Marshal(p)
//	MarshalParamsAndRequest(jsonStr, t, "/user/login", "")
//}
//
//func TestLoginOut(t *testing.T) {
//	Init()
//	token := common.RandStr(6)
//	MarshalParamsAndRequest([]byte{}, t, "/user/loginOut", token)
//}
//
//func TestUserInfo(t *testing.T) {
//	Init()
//	token := "e92f305b21f8e80b56332ce3a86fdd7c"
//	MarshalParamsAndRequest([]byte{}, t, "/user/userInfo", token)
//}
//
//func TestEditInfo(t *testing.T) {
//	Init()
//	token := "e92f305b21f8e80b56332ce3a86fdd7c"
//	p := request.EditInfo{
//		Name:  common.RandStr(6),
//		Title: common.RandStr(20),
//	}
//	jsonStr, _ := json.Marshal(p)
//	MarshalParamsAndRequest(jsonStr, t, "/user/editInfo", token)
//}
//
//func TestEditPwd(t *testing.T) {
//	Init()
//	token := "e92f305b21f8e80b56332ce3a86fdd7c"
//	newPwd := common.RandStr(6)
//	p := request.EditPwd{
//		Password:     common.RandStr(6),
//		NewPassword:  newPwd,
//		NewRpassword: newPwd,
//	}
//	jsonStr, _ := json.Marshal(p)
//	MarshalParamsAndRequest(jsonStr, t, "/user/editPwd", token)
//}
//
//func TestBindMail(t *testing.T) {
//	Init()
//	token := "e92f305b21f8e80b56332ce3a86fdd7c"
//	p := request.BindMail{
//		Mail: "491838375@qq.com",
//		Code: common.RandStr(4),
//	}
//	jsonStr, _ := json.Marshal(p)
//	MarshalParamsAndRequest(jsonStr, t, "/user/bindMal", token)
//}
//
//func TestRetrievePwd(t *testing.T) {
//	Init()
//	token := "e92f305b21f8e80b56332ce3a86fdd7c"
//	newPwd := common.RandStr(6)
//	p := request.RetrievePwd{
//		Mail:         "491838375@qq.com",
//		Code:         common.RandStr(4),
//		NewPassword:  newPwd,
//		NewRpassword: newPwd,
//	}
//	jsonStr, _ := json.Marshal(p)
//	MarshalParamsAndRequest(jsonStr, t, "/user/retrievePwd", token)
//}
//
//func TestFollow(t *testing.T) {
//	Init()
//	token := "e92f305b21f8e80b56332ce3a86fdd7c"
//	p := request.Follow{
//		FollowedPerson: 61,
//	}
//	jsonStr, _ := json.Marshal(p)
//	MarshalParamsAndRequest(jsonStr, t, "/user/follow", token)
//}
//
//func TestCancelFollow(t *testing.T) {
//	Init()
//	token := "e92f305b21f8e80b56332ce3a86fdd7c"
//	p := request.Follow{
//		FollowedPerson: 61,
//	}
//	jsonStr, _ := json.Marshal(p)
//	MarshalParamsAndRequest(jsonStr, t, "/user/cancelFollow", token)
//}
//
//func TestFollowList(t *testing.T) {
//	Init()
//	token := "e92f305b21f8e80b56332ce3a86fdd7c"
//	p := request.FollowList{
//		Page:     1,
//		PageSize: 5,
//	}
//	jsonStr, _ := json.Marshal(p)
//	MarshalParamsAndRequest(jsonStr, t, "/user/fansList", token)
//}
//
//func TestFansList(t *testing.T) {
//	Init()
//	token := "e92f305b21f8e80b56332ce3a86fdd7c"
//	p := request.FollowList{
//		Page:     1,
//		PageSize: 5,
//	}
//	jsonStr, _ := json.Marshal(p)
//	MarshalParamsAndRequest(jsonStr, t, "/user/followList", token)
//}
//
//func MarshalParamsAndRequest(p []byte, t *testing.T, path, token string) {
//	res := common.CurlPost(path, string(p), r, token)
//	bootstrap.Log.Info("测试接口:"+path, zap.Any("result", res))
//	re, _ := common.JsonToMap(res)
//	if re["state"] != true {
//		t.Error(re["msg"])
//	}
//}
