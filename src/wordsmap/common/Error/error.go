package Error

import "fmt"

// QNAPError define any error we want to present to response
type Error struct {
	Code    string
	Status  int
	Message string
}

// implement error interface
func (q Error) Error() string {
	return fmt.Sprintf("%s %d %s", q.Code, q.Status, q.Message)
}

// New err and append some useful message
func New(err Error, errMsg string) Error {
	return Error{err.Code, err.Status, fmt.Sprintf("%s -> %s", err.Message, errMsg)}
}