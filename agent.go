package polycode

type Agent interface {
	Call(options TaskOptions, input AgentInput) Response
}

type AgentBuilder interface {
	WithTenantId(tenantId string) AgentBuilder
	Get() Agent
}
