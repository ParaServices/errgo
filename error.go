package errgo

import (
	"github.com/magicalbanana/errorx"

	"github.com/nats-io/nuid"
	"go.uber.org/zap/zapcore"
)

// Error ...
type Error struct {
	Errorx  *errorx.Error `json:"errorx"`
	ID      string        `json:"error_id"`
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Details Details       `json:"details"`
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
	kv.AddObject("details", e.Details)
	kv.AddString("error_stack", string(e.Errorx.Stack()))
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
	return e
}

func (e *Error) Error() string {
	return e.Errorx.Error()
}

func (e *Error) Cause() error {
	return e.Errorx.Cause
}
