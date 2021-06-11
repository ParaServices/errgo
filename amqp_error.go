package errgo

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap/zapcore"
)

// AMQPError represents the error from the amqp package
type AMQPError struct {
	*amqp.Error
}

// MarshalLogObject ...
func (a AMQPError) MarshalLogObject(enc zapcore.ObjectEncoder) error {
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

// AddAMQError ...
func (e *Error) AddAMQError(err *amqp.Error) {
	if err == nil {
		return
	}
	e.AMQPError = &AMQPError{err}
}

// SetAMQPError ...
func (e *Error) SetAMQPError(amqpErr *amqp.Error) {
	if amqpErr == nil {
		return
	}
	e.AMQPError = &AMQPError{amqpErr}
}
