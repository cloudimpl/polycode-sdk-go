package polycode

import "github.com/cloudimpl/polycode-sdk-go/errors"

type Signal interface {
	Await() Response
	EmitValue(taskId string, data any) error
	EmitError(taskId string, err errors.Error) error
}
