package runtime

import (
	"github.com/cloudimpl/polycode-sdk-go"
	"github.com/gin-gonic/gin"
)

var CurrentRuntime Runtime

type Runtime interface {
	RegisterService(service Service) error
	RegisterApi(httpHandler *gin.Engine) error
	RegisterValidator(validator polycode.Validator) error
	GetValidator() polycode.Validator
	Start() error
}

func RegisterService(service Service) error {
	return CurrentRuntime.RegisterService(service)
}

func RegisterApi(httpHandler *gin.Engine) error {
	return CurrentRuntime.RegisterApi(httpHandler)
}

func RegisterValidator(validator polycode.Validator) error {
	return CurrentRuntime.RegisterValidator(validator)
}

func GetValidator() polycode.Validator {
	return CurrentRuntime.GetValidator()
}

func Start() error {
	return CurrentRuntime.Start()
}
