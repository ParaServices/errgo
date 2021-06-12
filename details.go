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
	defaultFn := func(key string, value interface{}) {
		enc.AddReflected(key, value)
	}
	for k, v := range d.Details {
		switch k {
		case "pq_error":
			switch pqErr := v.(type) {
			case *PQError:
				enc.AddObject("pq_error", pqErr)
			case PQError:
				enc.AddObject("pq_error", &pqErr)
			default:
				defaultFn(k, v)
			}
		case "amqp_error":
			switch amqpErr := v.(type) {
			case *AMQPError:
				enc.AddObject("amqp_error", amqpErr)
			case AMQPError:
				enc.AddObject("amqp_error", &amqpErr)
			default:
				defaultFn(k, v)
			}
		case "google_api_error":
			switch apiErr := v.(type) {
			case *GoogleAPIError:
				enc.AddObject("google_api_error", apiErr)
			case GoogleAPIError:
				enc.AddObject("google_api_error", &apiErr)
			default:
				defaultFn(k, v)
			}
		default:
			defaultFn(k, v)
		}
	}
	return nil
}
