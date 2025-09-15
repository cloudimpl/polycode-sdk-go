package polycode

type ClientChannel interface {
	Emit(data any) error
}
