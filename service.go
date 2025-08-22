package sdk

type Service interface {
	RequestReply(options TaskOptions, method string, input any) Response
	Send(options TaskOptions, method string, input any) error
}

type ServiceBuilder interface {
	WithTenantId(tenantId string) ServiceBuilder
	WithPartitionKey(partitionKey string) ServiceBuilder
	Get() Service
}
