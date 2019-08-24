package errgo

import "go.uber.org/zap/zapcore"

// Details represents the additional details for a given error (metadata)
type Details struct {
	Details map[string]string `json:"details,omitempty"`
}

// Add adds a detail to the details
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
