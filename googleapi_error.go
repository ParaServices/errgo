package errgo

import (
	"strconv"

	"google.golang.org/api/googleapi"
)

// AddGoogleAPIError ...
func (e Error) AddGoogleAPIError(err *googleapi.Error) {
	setDetails("google_api_error_error", err.Error(), &e.Details)
	setDetails("google_api_error_body", err.Body, &e.Details)
	setDetails("google_api_error_code", strconv.Itoa(err.Code), &e.Details)
	setDetails("google_api_error_message", err.Message, &e.Details)
	// for k, v := range err.Errors {
	//
	// }
}
