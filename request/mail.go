package request

type SendMail struct {
	Mail    string `json:"mail" binding:"required,email"`
	TypeStr string `json:"type" binding:"required"`
}
