package errgo

import "github.com/lib/pq"

func (e Error) AddPQError(pqErr *pq.Error) {
	setDetails("pq_error_code", string(pqErr.Code), &e.Details)
	setDetails("pq_error_code_class", string(pqErr.Code.Class()), &e.Details)
	setDetails("pq_error_code_class_name", string(pqErr.Code.Class().Name()), &e.Details)
	setDetails("pq_error_code_name", pqErr.Code.Name(), &e.Details)
	setDetails("pq_error_column", pqErr.Column, &e.Details)
	setDetails("pq_error_constraint", pqErr.Constraint, &e.Details)
	setDetails("pq_error_data_type_name", pqErr.DataTypeName, &e.Details)
	setDetails("pq_error_detail", pqErr.Detail, &e.Details)
	setDetails("pq_error_file", pqErr.File, &e.Details)
	setDetails("pq_error_hint", pqErr.Hint, &e.Details)
	setDetails("pq_error_line", pqErr.Line, &e.Details)
	setDetails("pq_error_internal_position", pqErr.InternalPosition, &e.Details)
	setDetails("pq_error_internal_query", pqErr.InternalQuery, &e.Details)
	setDetails("pq_error_routine", pqErr.Routine, &e.Details)
	setDetails("pq_error_schema", pqErr.Schema, &e.Details)
	setDetails("pq_error_severity", pqErr.Severity, &e.Details)
	setDetails("pq_error_message", pqErr.Message, &e.Details)
	setDetails("pq_error_detail", pqErr.Detail, &e.Details)
	setDetails("pq_error_position", pqErr.Position, &e.Details)
	setDetails("pq_error_where", pqErr.Where, &e.Details)
	setDetails("pq_error_table", pqErr.Table, &e.Details)
}
