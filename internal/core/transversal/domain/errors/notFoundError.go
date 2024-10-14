package domain

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(s string) error {
	return &NotFoundError{
		Message: s,
	}
}
