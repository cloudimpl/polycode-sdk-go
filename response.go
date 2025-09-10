package sdk

type Response interface {
	IsError() bool
	HasResult() bool
	Get(ret any) error
	GetAny() (any, error)
}
