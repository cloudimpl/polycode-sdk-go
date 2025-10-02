package polycode

type Service interface {
	RequestReply(options TaskOptions, method string, input any) Response
	Send(options TaskOptions, method string, input any) error
}

type ServiceBuilder interface {
	WithEnvId(envId string) ServiceBuilder
	Get() Service
}

type ControllerBuilder interface {
	WithEnvId(envId string) ControllerBuilder
	Get() Controller
}
