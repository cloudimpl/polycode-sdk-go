package apicontext

import (
	"context"
	"github.com/cloudimpl/polycode-sdk-go"
)

func FromContext(ctx context.Context) (polycode.ApiContext, error) {
	value := ctx.Value("polycode.context")
	if value == nil {
		return nil, polycode.ErrContextNotFound
	}

	return value.(polycode.ApiContext), nil
}
