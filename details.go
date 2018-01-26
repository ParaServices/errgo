package errgo

import "go.uber.org/zap/zapcore"

// Details ...
type Details struct {
	Details map[string]string `json:"details"`
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
