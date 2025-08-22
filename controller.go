package sdk

type Controller interface {
	Call(options TaskOptions, path string, apiReq ApiRequest) (ApiResponse, error)
}
