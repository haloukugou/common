package clientService

import (
	"dj/app/request/clientRequest"
	"github.com/gin-gonic/gin"
)

type Kong struct {
}

func (k *Kong) UserRegister(ctx *gin.Context, params clientRequest.Register) error {

	return nil
}
