package common

import (
	"crypto/md5"
	"crypto/rand"
	"dj/constants"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	time2 "time"
)

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Fail 返回失败信息
func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Res{Code: constants.FAIL_CODE, Msg: msg, Data: map[string]interface{}{}})
}

// Success 返回成功信息
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Res{Code: constants.SUCCESS_CODE, Msg: "success", Data: data})
}

// Ok 返回没有data的成功信息
func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, Res{Code: constants.SUCCESS_CODE, Msg: "success", Data: map[string]interface{}{}})
}

// RandStr 随机字符串
func RandStr(length int) string {
	baseStr := "1234567890qwertyuiopasdfghjklzxcvbnm"
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	for i, b := range bytes {
		bytes[i] = baseStr[b%byte(len(baseStr))]
	}
	return string(bytes)
}

// CreateMd5Str 生成md5字符串
func CreateMd5Str(str1, str2 string) string {
	m := md5.New()
	_, err := io.WriteString(m, str1+str2)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", m.Sum(nil))
}

// CreateToken 生成token
func CreateToken(account, randStr string) string {
	str := RandStr(6)
	time := time2.Now().UnixMicro()
	timeStr := fmt.Sprintf("%d", time)
	return CreateMd5Str(account+timeStr, str+randStr)
}

// JsonToMap json转map
func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	return mapResult, err
}

// MapToJson map转json
func MapToJson(data map[string]string) (string, error) {
	jsonStr, err := json.Marshal(data)
	return string(jsonStr), err
}

// CurlPost curl的post请求
func CurlPost(path, params string, r *gin.Engine, header string) string {
	req := httptest.NewRequest("POST", path, strings.NewReader(params))
	if header != "" {
		req.Header.Set("token", header)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w.Body.String()
}

// PathExists 路劲是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// IsContainStr 字符串是否存在数组中
func IsContainStr(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// IsContainInt 数字是否存在数组中
func IsContainInt(items []int, item int) bool {
	for _, value := range items {
		if value == item {
			return true
		}
	}
	return false
}

// VerifyEmailFormat 匹配电子邮箱
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// ValidateMobile 校验手机号
func ValidateMobile(fl validator.FieldLevel) bool {
	fmt.Println("123123123")
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}
