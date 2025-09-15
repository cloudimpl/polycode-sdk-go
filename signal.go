package sdk

import "github.com/cloudimpl/byte-sdk-go/errors"

type Signal interface {
	Await() Response
	EmitValue(taskId string, data any) error
	EmitError(taskId string, err errors.Error) error
}
