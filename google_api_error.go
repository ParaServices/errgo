package errgo

import (
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/api/googleapi"
)

type HTTPHeader http.Header

func (h HTTPHeader) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if len(h) < 1 {
		return nil
	}
	for k, v := range h {
		zap.Strings(k, v)
	}

	return nil

}

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
	if header := g.Header; header != nil && len(header) > 0 {
		enc.AddObject("header", HTTPHeader(header))
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
func (e *Error) AddGoogleAPIError(apiErr *googleapi.Error) {
	if apiErr == nil {
		return
	}
	e.AddDetail("google_api_error", &GoogleAPIError{apiErr})
}

// SetGoogleAPIError ...
func (e *Error) SetGoogleAPIError(apiErr *googleapi.Error) {
	if apiErr == nil {
		return
	}
	e.AddDetail("google_api_error", &GoogleAPIError{apiErr})
}
