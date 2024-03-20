package response

type Error struct {
	Message string `json:"message"`
}

func Err(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}
