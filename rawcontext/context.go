package rawcontext

import (
	"context"
	"github.com/cloudimpl/polycode-sdk-go"
)

func FromContext(ctx context.Context) (polycode.RawContext, error) {
	rawCtx, ok := ctx.(polycode.RawContext)
	if ok {
		return rawCtx, nil
	}

	value := ctx.Value("polycode.context")
	if value == nil {
		return nil, polycode.ErrContextNotFound
	}

	return value.(polycode.RawContext), nil
}
