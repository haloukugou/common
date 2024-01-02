package test

import (
	"dj/request"
	"encoding/json"
	"testing"
)

func TestSendMail(t *testing.T) {
	Init()
	p := request.SendMail{
		Mail:    "18280189749@163.com",
		TypeStr: "login",
	}
	mar, _ := json.Marshal(p)
	MarshalParamsAndRequest(mar, t, "/service/sendMail", "")
}
