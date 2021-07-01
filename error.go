package errgo

import (
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/magicalbanana/errorx"
	"github.com/rs/xid"
	"github.com/streadway/amqp"
	"google.golang.org/api/googleapi"

	"go.uber.org/zap/zapcore"
)

// Error ...
type Error struct {
	Errorx  *errorx.Error `json:"errorx,omitempty"`
	ErrorID string        `json:"error_id,omitempty"`
	Code    string        `json:"code,omitempty"`
	Message string        `json:"message,omitempty"`
	Details Details       `json:"details,omitempty"`
}

func (e *Error) GetErrorx() *errorx.Error {
	return e.Errorx
}

func (e *Error) GetErrorID() string {
	if e.ErrorID == "" {
		e.SetErrorID(xid.New().String())
	}
	return e.ErrorID
}

func (e *Error) GetCode() string {
	return e.Code
}

func (e *Error) GetMessage() string {
	return e.Message
}

func (e *Error) GetDetails() Details {
	return e.Details
}

func (e *Error) HasDetails() bool {
	return (e.Details != nil && len(e.Details) > 0)
}

func (e *Error) GetPQError() *PQError {
	v, ok := e.Details[PQErrorKey]
	if !ok {
		return nil
	}
	pqErr, ok := v.(*PQError)
	if !ok {
		return nil
	}
	return pqErr
}

func (e *Error) GetGoogleAPIError() *GoogleAPIError {
	v, ok := e.Details[GoogleAPIErrorKey]
	if !ok {
		return nil
	}
	googleAPIError, ok := v.(*GoogleAPIError)
	if !ok {
		return nil
	}
	return googleAPIError
}

func (e *Error) GetAMQPError() *AMQPError {
	v, ok := e.Details[AMQPErrorKey]
	if !ok {
		return nil
	}
	amqpError, ok := v.(*AMQPError)
	if !ok {
		return nil
	}
	return amqpError
}

func (e *Error) SetErrorID(id string) {
	e.ErrorID = id
}

func (e *Error) SetCode(code string) {
	e.Code = code
}

func (e *Error) SetMessage(message string) {
	e.Message = message
}

func (e *Error) SetDetails(details Details) {
	if details == nil || len(details) < 1 {
		return
	}

	for k, v := range details {
		e.AddDetail(k, v)
	}
}

func (e *Error) AddDetail(key string, value interface{}) {
	if e.Details == nil {
		e.Details = make(Details)
	}
	e.Details.Add(key, value)
}

// MarshalLogObject allows the Error to be passed as an object to the
// paralog.Logger interface.
func (e Error) MarshalLogObject(kv zapcore.ObjectEncoder) error {
	kv.AddString("error_id", e.GetErrorID())
	if v := e.GetCode(); v != "" {
		kv.AddString("code", v)
	}
	if v := e.GetMessage(); v != "" {
		kv.AddString("message", v)
	}
	if v := e.Error(); v != "" {
		kv.AddString("error", v)
	}
	if e.HasDetails() {
		kv.AddObject("details", e.Details)
	}
	if stack := e.Stack(); stack != nil && len(stack) > 0 {
		kv.AddByteString("stack_trace", stack)
	}
	return nil
}

// New ...
func New(err error) *Error {
	if err == nil {
		return nil
	}

	errx, ok := err.(*Error)
	if ok {
		return errx
	}

	e := &Error{
		ErrorID: xid.New().String(),
		Errorx:  errorx.New(err),
	}

	switch v := err.(type) {
	case *pq.Error:
		e.SetPQError(v)
	case *googleapi.Error:
		e.SetGoogleAPIError(v)
	case *amqp.Error:
		e.SetAMQPError(v)
	}

	return e
}

func NewF(s string, args ...interface{}) *Error {
	if s == "" {
		return nil
	}

	if len(args) < 1 {
		return New(errors.New(s))
	}

	return New(fmt.Errorf(s, args...))
}

// Error ...
func (e *Error) Error() string {
	if e.Errorx != nil {
		return e.Errorx.Error()
	}
	return ""
}

// Cause ...
func (e *Error) Cause() error {
	if e.Errorx != nil {
		return e.Errorx.Cause
	}
	return nil
}

func (e *Error) Stack() []byte {
	if e.Errorx != nil {
		return e.Errorx.Stack()
	}
	return nil
}

var _ error = (*Error)(nil)
var _ ErrorGetter = (*Error)(nil)
var _ ErrorSetter = (*Error)(nil)
var _ ErrorAccessor = (*Error)(nil)

type ErrorGetter interface {
	GetErrorx() *errorx.Error
	GetErrorID() string
	GetCode() string
	GetMessage() string
	GetDetails() Details
	GetPQError() *PQError
	GetGoogleAPIError() *GoogleAPIError
	GetAMQPError() *AMQPError
	Error() string
	Cause() error
	Stack() []byte
}

type ErrorSetter interface {
	SetErrorID(id string)
	SetCode(code string)
	SetMessage(message string)
	SetDetails(details Details)
	AddDetail(key string, value interface{})
	SetPQError(pqErr *pq.Error)
	SetGoogleAPIError(apiErr *googleapi.Error)
	SetAMQPError(amqpErr *amqp.Error)
}

type ErrorAccessor interface {
	ErrorGetter
	ErrorSetter
}
