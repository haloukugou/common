package bootstrap

import (
	"dj/app/common"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func InitializeValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		e := v.RegisterValidation("mobile", common.ValidateMobile)
		if e != nil {
			panic(e.Error())
		}

		// 注册自定义 json tag 函数
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
