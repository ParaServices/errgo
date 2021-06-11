package errgo

import (
	"go.uber.org/zap/zapcore"
	"google.golang.org/api/googleapi"
)

// GoogleAPIError represents the error that is returned from the googleapi
// package
type GoogleAPIError struct {
	*googleapi.Error
}

// MarshalLogObject satisfies the interface for the paralog.Logger interface
func (g GoogleAPIError) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if g.Error == nil {
		return nil
	}
	if err := g.Error.Error(); err != "" {
		enc.AddString("error", err)
	}
	if body := g.Body; body != "" {
		enc.AddString("body", body)
	}
	if code := g.Code; code != 0 {
		enc.AddInt("code", code)
	}
	if message := g.Message; message != "" {
		enc.AddString("message", message)
	}
	if len(g.Errors) > 0 {
		enc.AddArray("errors", zapcore.ArrayMarshalerFunc(func(ae zapcore.ArrayEncoder) error {
			for _, err := range g.Errors {
				ae.AppendObject(zapcore.ObjectMarshalerFunc(func(oe zapcore.ObjectEncoder) error {
					if message := err.Message; message != "" {
						enc.AddString("message", message)
					}
					if reason := err.Reason; reason != "" {
						enc.AddString("reason", reason)
					}
					return nil
				}))
			}
			return nil
		}))
	}
	return nil
}

// AddGoogleAPIError ...
func (e *Error) AddGoogleAPIError(err *googleapi.Error) {
	if err == nil {
		return
	}
	e.GoogleAPIError = &GoogleAPIError{err}
}

// SetGoogleAPIError ...
func (e *Error) SetGoogleAPIError(apiErr *googleapi.Error) {
	if apiErr == nil {
		return
	}
	e.GoogleAPIError = &GoogleAPIError{apiErr}
}
