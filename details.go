package errgo

import "go.uber.org/zap/zapcore"

// Details represents the additional details for a given error (metadata)
type Details struct {
	Details map[string]interface{} `json:"details,omitempty"`
}

// Add adds a detail to the details
func (d Details) Add(key string, value interface{}) {
	if d.Details == nil {
		d.Details = make(map[string]interface{})
	}
	d.Details[key] = value
}

// MarshalLogObject ...
func (d Details) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	for k, v := range d.Details {
		enc.AddReflected(k, v)
	}
	return nil
}
