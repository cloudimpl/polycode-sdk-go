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

func RegisterService(service Service) error {
	return CurrentRuntime.RegisterService(service)
}

func RegisterApi(httpHandler *gin.Engine) error {
	return CurrentRuntime.RegisterApi(httpHandler)
}

func Start() error {
	return CurrentRuntime.Start()
}
