package errgo

import (
	"github.com/lib/pq"
	"go.uber.org/zap/zapcore"
)

// PQError represents the error from the pq package
type PQError struct {
	*pq.Error
}

// MarshalLogObject ...
func (p PQError) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if p.Error == nil {
		return nil
	}
	if code := string(p.Code); code != "" {
		enc.AddString("code", code)
		if name := string(p.Code.Name()); name != "" {
			enc.AddString("code_name", name)
		}
		if class := string(p.Code.Class()); class != "" {
			enc.AddString("code_class", class)
		}
		if name := string(p.Code.Class().Name()); name != "" {
			enc.AddString("code_class_name", name)
		}
		enc.AddString("code_name", p.Code.Name())
	}
	if column := p.Column; column != "" {
		enc.AddString("column", column)
	}
	if constraint := p.Constraint; constraint != "" {
		enc.AddString("constraint", constraint)
	}
	if typeName := p.DataTypeName; typeName != "" {
		enc.AddString("data_type_name", typeName)
	}
	if detail := p.Detail; detail != "" {
		enc.AddString("detail", detail)
	}
	if file := p.File; file != "" {
		enc.AddString("file", file)
	}
	if hint := p.Hint; hint != "" {
		enc.AddString("hint", hint)
	}
	if line := p.Line; line != "" {
		enc.AddString("line", line)
	}
	if position := p.InternalPosition; position != "" {
		enc.AddString("internal_position", position)
	}
	if query := p.InternalQuery; query != "" {
		enc.AddString("internal_query", query)
	}
	if routine := p.Routine; routine != "" {
		enc.AddString("routine", routine)
	}
	if schema := p.Schema; schema != "" {
		enc.AddString("schema", schema)
	}
	if severity := p.Severity; severity != "" {
		enc.AddString("severity", severity)
	}
	if message := p.Message; message != "" {
		enc.AddString("message", message)
	}
	if position := p.Position; position != "" {
		enc.AddString("position", position)
	}
	if where := p.Where; where != "" {
		enc.AddString("where", where)
	}
	if table := p.Table; table != "" {
		enc.AddString("table", table)
	}
	return nil
}

// AddPQError ...
func (e *Error) AddPQError(pqError *pq.Error) {
	if pqError == nil {
		return
	}
	e.PQError = &PQError{pqError}
}

// SetPQError ...
func (e *Error) SetPQError(pqErr *pq.Error) {
	if pqErr == nil {
		return
	}
	e.PQError = &PQError{pqErr}
}
