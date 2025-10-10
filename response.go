package polycode

type Response interface {
	Get(ret any) error
	GetAny() (any, error)

	HasResult() bool
	IsError() bool
	Output() any
	Error() Error
}
