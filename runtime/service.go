package runtime

import "github.com/cloudimpl/polycode-sdk-go"

type Service interface {
	GetName() string
	GetDescription(method string) (string, error)
	GetInputType(method string) (any, error)
	GetOutputType(method string) (any, error)
	IsWorkflow(method string) bool
	ExecuteService(ctx polycode.ServiceContext, method string, input any) (any, error)
	ExecuteWorkflow(ctx polycode.WorkflowContext, method string, input any) (any, error)
}
