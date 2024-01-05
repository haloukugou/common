package request

type LatestParams struct {
	ClientVersion string `json:"client_version" binding:"required"`
}
