package polycode

import (
	"context"
)

// unexported key type to avoid collisions in context values
type ctxKey struct{}

// WithContext attaches your *ContextImpl to a parent context.
// Call this once at the edge (HTTP middleware, gRPC interceptor, Lambda handler, etc.)
func WithContext(parent context.Context, impl BaseContext) context.Context {
	return context.WithValue(parent, ctxKey{}, impl)
}

// getImpl is the internal helper that extracts the *ContextImpl from a context.
// It should never be exported; callers should use typed accessors like ApiContextFrom or RawContextFrom.
func getImpl(ctx context.Context) (PolycodeContext, bool) {
	v := ctx.Value(ctxKey{})
	if v == nil {
		return nil, false
	}
	impl, ok := v.(PolycodeContext)
	return impl, ok
}

type BaseContext interface {
	context.Context
	Meta() HandlerContextMeta
	AuthContext() AuthContext
	Logger() Logger
	FileStore() Folder
	TempFileStore() Folder
}

type ServiceContext interface {
	BaseContext
	Db() DataStoreBuilder
}

type WorkflowContext interface {
	BaseContext
	ReadOnlyDb() ReadOnlyDataStoreBuilder
	Service(service string) ServiceBuilder
	Agent(agent string) AgentBuilder
	ServiceEx(envId string, service string) ServiceBuilder
	AgentEx(envId string, agent string) AgentBuilder
	App(appName string) Service
	AppEx(envId string, appName string) Service
	Controller(controller string) Controller
	ControllerEx(envId string, controller string) Controller
	Memo(getter func() (any, error)) Response
	Signal(signalName string) Signal
	ClientChannel(channelName string) ClientChannel
	Lock(key string) Lock
}

type ApiContext interface {
	WorkflowContext
}

type RawContext interface {
	BaseContext
	GetMeta(group string, typeName string, key string) (map[string]interface{}, error)
}

type PolycodeContext interface {
	BaseContext
	ServiceContext
	WorkflowContext
	ApiContext
	RawContext
}

func ApiContextFrom(ctx context.Context) (ApiContext, bool) {
	value := ctx.Value("polycode.context")
	if value == nil {
		return nil, false
	}

	return value.(ApiContext), true
}

func RawContextFrom(ctx context.Context) (ApiContext, bool) {
	value := ctx.Value("polycode.context")
	if value == nil {
		return nil, false
	}

	return value.(ApiContext), true
}
