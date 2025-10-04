package polycode

type Signal interface {
	Await() Response
	EmitValue(taskId string, data any) error
	EmitError(taskId string, err Error) error
}
