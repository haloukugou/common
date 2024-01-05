package request

type RegisterParams struct {
	Account    string `json:"account"`
	Password   string `json:"password"`
	Rpassword  string `json:"rpassword"`
	TypeString string `json:"typeString" binding:"required"`
	Mail       string `json:"mail"`
	Code       string `json:"code"`
	Source     int    `json:"source" binding:"required"`
}

type LoginParams struct {
	Account    string `json:"account" binding:"required"`
	Password   string `json:"password"`
	TypeString string `json:"typeString" binding:"required"`
	Code       string `json:"code"`
	Source     int    `json:"source" binding:"required"`
}

type EditInfo struct {
	Name  string `json:"name" binding:"required"`
	Title string `json:"title"`
}

type EditPwd struct {
	Password     string `json:"password"`
	NewPassword  string `json:"newPassword" binding:"required"`
	NewRpassword string `json:"newRpassword" binding:"required"`
}

type BindMail struct {
	Mail   string `json:"mail" binding:"required,email"`
	Code   string `json:"code" binding:"required"`
	Source int    `json:"source" binding:"required"`
}

type RetrievePwd struct {
	Mail         string `json:"mail" binding:"required,email"`
	Code         string `json:"code" binding:"required"`
	NewPassword  string `json:"newPassword" binding:"required"`
	NewRpassword string `json:"newRpassword" binding:"required"`
	Source       int    `json:"source" binding:"required"`
}

type Follow struct {
	FollowedPerson uint64 `json:"followedPerson" binding:"required"`
}

type FollowList struct {
	Page     uint64 `json:"page" binding:"required"`
	PageSize uint64 `json:"pageSize" binding:"required"`
}
