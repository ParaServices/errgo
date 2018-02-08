package errgo

import (
	"go.uber.org/zap/zapcore"
	"google.golang.org/api/googleapi"
)

// GoogleAPIError ...
type GoogleAPIError struct {
	*googleapi.Error
}

// MarshalLogObject ...
func (g GoogleAPIError) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("error", g.Error.Error())
	enc.AddString("error_body", g.Body)
	enc.AddInt("error_code", g.Code)
	enc.AddString("error_message", g.Message)
	enc.AddArray("errors", zapcore.ArrayMarshalerFunc(func(ae zapcore.ArrayEncoder) error {
		for _, err := range g.Errors {
			ae.AppendObject(zapcore.ObjectMarshalerFunc(func(oe zapcore.ObjectEncoder) error {
				oe.AddString("messsage", err.Message)
				oe.AddString("reason", err.Reason)
				return nil
			}))
		}
		return nil
	}))
	return nil
}

// AddGoogleAPIError ...
func (e *Error) AddGoogleAPIError(err *googleapi.Error) {
	e.GoogleAPIError = &GoogleAPIError{err}
}
