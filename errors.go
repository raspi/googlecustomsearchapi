package googlecustomsearchapi

import "fmt"

// JSON errors from API

type Error struct {
	Error ErrorError `json:"error"`
}

var _ error = ErrorError{}

type ErrorError struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Errors  []ErrorElement `json:"errors"`
	Status  string         `json:"status"`
}

func (e ErrorError) Error() string {
	// TODO: add more details
	return fmt.Sprintf(`%v %v: %v`, e.Code, e.Status, e.Message)
}

type ErrorElement struct {
	Message string `json:"message"`
	Domain  string `json:"domain"`
	Reason  string `json:"reason"`
}
