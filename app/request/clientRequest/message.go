package clientRequest

type SendMsg struct {
	MsgType int    `json:"msg_type" binding:"required,oneof=1 2"`
	Mobile  string `json:"mobile" binding:"required,mobile"`
}
