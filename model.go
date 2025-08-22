package sdk

import (
	"github.com/cloudimpl/byte-os/sdk/errors"
	"time"
)

type LogLevel string

type BackoffStrategy struct {
	InitialInterval time.Duration `json:"initialInterval"`
	MaxInterval     time.Duration `json:"maxInterval"`
	Multiplier      float64       `json:"multiplier"`
}

type TaskOptions struct {
	Timeout         time.Duration   `json:"timeout"`
	Retries         int             `json:"retries"`
	RetryOnFail     bool            `json:"retryOnFail"`
	BackoffStrategy BackoffStrategy `json:"backoffStrategy"`
}

func (t TaskOptions) WithTimeout(timeout time.Duration) TaskOptions {
	t.Timeout = timeout
	return t
}

type AuthContext struct {
	Claims map[string]interface{} `json:"claims"`
}

type ApiRequest struct {
	Id              string            `json:"id"`
	Host            string            `json:"host"`
	Method          string            `json:"method"`
	Path            string            `json:"path"`
	Query           map[string]string `json:"query"`
	Header          map[string]string `json:"header"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
}

type ApiResponse struct {
	StatusCode      int               `json:"statusCode"`
	Header          map[string]string `json:"header"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
}

type AgentInput struct {
	SessionKey string            `json:"sessionKey"`
	TaskId     int64             `json:"taskId"`
	LLMInput   LLMInput          `json:"llmInput"`
	Labels     map[string]string `json:"labels"`
	ChannelId  string            `json:"channelId"`
}

type LLMInput struct {
	Text string `json:"text"`
}

type HandlerContextMeta struct {
	OrgId        string            `json:"orgId"`
	EnvId        string            `json:"envId"`
	AppName      string            `json:"appName"`
	AppId        string            `json:"appId"`
	TenantId     string            `json:"tenantId"`
	PartitionKey string            `json:"partitionKey"`
	TaskGroup    string            `json:"taskGroup"`
	TaskName     string            `json:"taskName"`
	TaskId       string            `json:"taskId"`
	ParentId     string            `json:"parentId"`
	TraceId      string            `json:"traceId"`
	InputId      string            `json:"inputId"`
	Caller       CallerContextMeta `json:"caller"`
}

type CallerContextMeta struct {
	OrgId        string `json:"orgId"`
	EnvId        string `json:"envId"`
	AppName      string `json:"appName"`
	AppId        string `json:"appId"`
	TenantId     string `json:"tenantId"`
	PartitionKey string `json:"partitionKey"`
	TaskGroup    string `json:"taskGroup"`
	TaskName     string `json:"taskName"`
	TaskId       string `json:"taskId"`
}

type Response struct {
	output  any
	isError bool
	error   errors.Error
}

func (r Response) IsError() bool {
	return r.isError
}

func (r Response) HasResult() bool {
	return r.output != nil
}

func (r Response) Get(ret any) error {
	if r.isError {
		return r.error
	}

	return ConvertType(r.output, ret)
}

func (r Response) GetAny() (any, error) {
	if r.isError {
		return nil, r.error
	} else {
		return r.output, nil
	}
}
