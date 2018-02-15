package errgo

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap/zapcore"
)

// AMQPError ...
type AMQPError struct {
	*amqp.Error
}

// MarshalLogObject ...
func (a AMQPError) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("code", a.Code)
	enc.AddString("reason", a.Reason)
	enc.AddBool("server", a.Server)
	enc.AddBool("recover", a.Recover)
	return nil
}

// AddAMQError ...
func (e *Error) AddAMQError(err *amqp.Error) {
	e.AMQPError = &AMQPError{err}
}
