package polycode

type Validator interface {
	Validate(obj any) error
}
