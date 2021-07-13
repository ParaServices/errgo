package errgo

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap/zapcore"
)

const AMQPErrorKey = "amqp_error"

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

// AMQPError represents the error from the amqp package
type AMQPError struct {
	*amqp.Error
}

// MarshalLogObject ...
func (a AMQPError) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if a.Error == nil {
		return nil
	}
	if code := a.Code; code != 0 {
		enc.AddInt("code", a.Code)
	}
	if reason := a.Reason; reason != "" {
		enc.AddString("reason", reason)
	}
	enc.AddBool("server", a.Server)
	enc.AddBool("recover", a.Recover)
	return nil
}

// SetAMQPError ...
func (e *Error) SetAMQPError(amqpErr *amqp.Error) {
	if amqpErr == nil {
		return
	}
	e.AddDetail(AMQPErrorKey, &AMQPError{amqpErr})
}
