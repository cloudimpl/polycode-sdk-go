package polycode

type Response interface {
	HasResult() bool
	IsError() bool
	Get(ret any) error
	GetAny() any
	Error() Error
}
