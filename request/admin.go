package request

type AdminLoginParams struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ApkListParams struct {
	Page     uint64 `json:"page" binding:"required"`
	PageSize uint64 `json:"pageSize" binding:"required"`
}

type ReleaseParams struct {
	File    string `json:"file" binding:"required"`
	Version string `json:"version" binding:"required"`
	IsForce int    `json:"is_force" binding:"required,oneof=0 1"`
}
