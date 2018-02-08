package errgo

import (
	"github.com/lib/pq"
	"go.uber.org/zap/zapcore"
)

// PQError ...
type PQError struct {
	*pq.Error
}

// MarshalLogObject ...
func (p PQError) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("error_code", string(p.Code))
	enc.AddString("error_code_class", string(p.Code.Class()))
	enc.AddString("error_code_class_name", string(p.Code.Class().Name()))
	enc.AddString("error_code_name", p.Code.Name())
	enc.AddString("error_column", p.Column)
	enc.AddString("error_constraint", p.Constraint)
	enc.AddString("error_data_type_name", p.DataTypeName)
	enc.AddString("error_detail", p.Detail)
	enc.AddString("error_file", p.File)
	enc.AddString("error_hint", p.Hint)
	enc.AddString("error_line", p.Line)
	enc.AddString("error_internal_position", p.InternalPosition)
	enc.AddString("error_internal_query", p.InternalQuery)
	enc.AddString("error_routine", p.Routine)
	enc.AddString("error_schema", p.Schema)
	enc.AddString("error_severity", p.Severity)
	enc.AddString("error_message", p.Message)
	enc.AddString("error_detail", p.Detail)
	enc.AddString("error_position", p.Position)
	enc.AddString("error_where", p.Where)
	enc.AddString("error_table", p.Table)
	return nil
}

// AddPQError ...
func (e *Error) AddPQError(pqError *pq.Error) {
	e.PQError = &PQError{pqError}
}
