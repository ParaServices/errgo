package errgo

import (
	"github.com/magicalbanana/errorx"

	"github.com/nats-io/nuid"
	"go.uber.org/zap/zapcore"
)

// Error ...
type Error struct {
	Errorx         *errorx.Error   `json:"errorx"`
	ID             string          `json:"error_id"`
	Code           string          `json:"code"`
	Message        string          `json:"message"`
	Details        Details         `json:"details"`
	PQError        *PQError        `json:"pq_error,omitempty"`
	GoogleAPIError *GoogleAPIError `json:"google_api_error,omitempty"`
	AMQPError      *AMQPError      `json:"amqp_error,omitempty"`
}

// MarshalLogObject ...
func (e Error) MarshalLogObject(kv zapcore.ObjectEncoder) error {
	kv.AddString("error_id", e.ID)
	if len(e.Code) != 0 {
		kv.AddString("code", e.Code)
	}
	if len(e.Message) != 0 {
		kv.AddString("message", e.Message)
	}
	kv.AddString("error", e.Error())
	if len(e.Details.Details) > 0 {
		kv.AddObject("details", e.Details)
	}
	if e.PQError != nil {
		kv.AddObject("pq_error", e.PQError)
	}
	if e.GoogleAPIError != nil {
		kv.AddObject("google_api_error", e.GoogleAPIError)
	}
	if e.AMQPError != nil {
		kv.AddObject("amqp_error", e.AMQPError)
	}
	kv.AddByteString("error_stack", e.Errorx.Stack())
	return nil
}

// New ...
func New(err error) *Error {
	if err == nil {
		return nil
	}
	e := &Error{}
	e.ID = nuid.Next()
	e.Errorx = errorx.New(err)
	e.Details = Details{Details: make(map[string]string)}
	e.PQError = (*PQError)(nil)
	e.GoogleAPIError = (*GoogleAPIError)(nil)
	e.AMQPError = (*AMQPError)(nil)
	return e
}

// Error ...
func (e *Error) Error() string {
	return e.Errorx.Error()
}

// Cause ...
func (e *Error) Cause() error {
	return e.Errorx.Cause
}
