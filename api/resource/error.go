package resource

import "fmt"

type ValidationError struct {
	Message string
	Err     error
}

func (r *ValidationError) Error() string {
	return fmt.Sprintf("%s", r.Message)
}
