package errgo

import (
	"github.com/lib/pq"
	"github.com/magicalbanana/errorx"
	"github.com/streadway/amqp"
	"google.golang.org/api/googleapi"

	"github.com/nats-io/nuid"
	"go.uber.org/zap/zapcore"
)

// Error ...
type Error struct {
	Errorx         *errorx.Error   `json:"errorx,omitempty"`
	ID             string          `json:"error_id,omitempty"`
	Code           string          `json:"code,omitempty"`
	Message        string          `json:"message,omitempty"`
	Details        *Details        `json:"details,omitempty"`
	PQError        *PQError        `json:"pq_error,omitempty"`
	GoogleAPIError *GoogleAPIError `json:"google_api_error,omitempty"`
	AMQPError      *AMQPError      `json:"amqp_error,omitempty"`
}

// MarshalLogObject allows the Error to be passed as an object to the
// paralog.Logger interface.
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
	e := &Error{
		ID:     nuid.Next(),
		Errorx: errorx.New(err),
		Details: &Details{
			Details: make(map[string]string),
		},
	}

	pqErr, ok := err.(*pq.Error)
	if ok {
		e.AddPQError(pqErr)
	}
	googleAPIErr, ok := err.(*googleapi.Error)
	if ok {
		e.AddGoogleAPIError(googleAPIErr)
	}
	amqpErr, ok := err.(*amqp.Error)
	if ok {
		e.AddAMQError(amqpErr)
	}
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
