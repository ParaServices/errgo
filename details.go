package errgo

import "go.uber.org/zap/zapcore"

// Details represents the additional details for a given error (metadata)
type Details map[string]interface{}

// Add adds a detail to the details
func (d Details) Add(key string, value interface{}) {
	if d == nil {
		d = make(map[string]interface{})
	}
	d[key] = value
}

func (d Details) Get(key string) interface{} {
	v, ok := d[key]
	if ok {
		return v
	}
	return nil
}

// MarshalLogObject ...
func (d Details) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	defaultFn := func(key string, value interface{}) {
		enc.AddReflected(key, value)
	}
	for k, v := range d {
		switch k {
		case PQErrorKey:
			switch pqErr := v.(type) {
			case *PQError:
				enc.AddObject(PQErrorKey, pqErr)
			case PQError:
				enc.AddObject(PQErrorKey, &pqErr)
			default:
				defaultFn(k, v)
			}
		case AMQPErrorKey:
			switch amqpErr := v.(type) {
			case *AMQPError:
				enc.AddObject(AMQPErrorKey, amqpErr)
			case AMQPError:
				enc.AddObject(AMQPErrorKey, &amqpErr)
			default:
				defaultFn(k, v)
			}
		case GoogleAPIErrorKey:
			switch apiErr := v.(type) {
			case *GoogleAPIError:
				enc.AddObject(GoogleAPIErrorKey, apiErr)
			case GoogleAPIError:
				enc.AddObject(GoogleAPIErrorKey, &apiErr)
			default:
				defaultFn(k, v)
			}
		default:
			defaultFn(k, v)
		}
	}
	return nil
}
