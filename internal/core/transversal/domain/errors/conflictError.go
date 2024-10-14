package domain

type ConflictError struct {
	Message string
}

func (e *ConflictError) Error() string {
	return e.Message
}

func NewConflictError(s string) error {
	return &ConflictError{
		Message: s,
	}
}
