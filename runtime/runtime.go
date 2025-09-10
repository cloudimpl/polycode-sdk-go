package runtime

import (
	"github.com/gin-gonic/gin"
)

var CurrentRuntime Runtime

type Runtime interface {
	RegisterService(service Service) error
	RegisterApi(httpHandler *gin.Engine) error
	Start() error
}
