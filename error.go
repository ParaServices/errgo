package errgo

import (
	"github.com/magicalbanana/errorx"

	"github.com/nats-io/nuid"
	"go.uber.org/zap/zapcore"
)

// Error ...
type Error struct {
	errorx.Error
	ID      string
	Code    string
	Message string
	Details Details
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
	kv.AddString("error", e.Error.Cause.Error())
	kv.AddObject("details", e.Details)
	kv.AddString("error_stack", string(e.Error.Stack()))
	return nil
}

// New ...
func New(err error) *Error {
	e := &Error{}
	e.ID = nuid.Next()
	e.Error = *errorx.New(err)
	e.Details = Details{Details: make(map[string]string)}
	return e
}

// Details ...
type Details struct {
	Details map[string]string
}

// Add ...
func (d Details) Add(key, value string) {
	d.Details[key] = value
}

// MarshalLogObject ...
func (d Details) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	for k, v := range d.Details {
		enc.AddString(k, v)
	}
	return nil
}
