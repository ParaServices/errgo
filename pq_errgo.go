package errgo

import "github.com/lib/pq"

func (e Error) AddPQError(pqErr *pq.Error) {
	e.Details.Add("pq_error_code", string(pqErr.Code))
	e.Details.Add("pq_error_code_class", pqErr.Code.Class())
	e.Details.Add("pq_error_code_name", pqErr.Code.Name())
	e.Details.Add("pq_error_column", pqErr.Column)
	e.Details.Add("pq_error_constraint", pqErr.Constraint)
	e.Details.Add("pq_error_data_type_name", pqErr.DataTypeName)
	e.Details.Add("pq_error_detail", pqErr.Detail)
	e.Details.Add("pq_error_file", pqErr.File)
	e.Details.Add("pq_error_hint", pqErr.Hint)
	e.Details.Add("pq_error_internal_line", pqErr.Line)
	e.Details.Add("pq_error_internal_position", pqErr.InternalPosition)
	e.Details.Add("pq_error_internal_query", pqErr.InternalQuery)
	e.Details.Add("pq_error_internal_routine", pqErr.Routine)
	e.Details.Add("pq_error_internal_schema", pqErr.Schema)
	e.Details.Add("pq_error_internal_severity", pqErr.Severity)
	e.Details.Add("pq_error_message", pqErr.Message)
	e.Details.Add("pq_error_detail", pqErr.Detail)
	e.Details.Add("pq_error_position", pqErr.Position)
	e.Details.Add("pq_error_where", pqErr.Where)
	e.Details.Add("pq_error_table", pqErr.Table)
}
