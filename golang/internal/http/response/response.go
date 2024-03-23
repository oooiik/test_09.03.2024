package response

type Error struct {
	Message string `json:"message"`
}

func Err(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}

type Meta struct {
	Total   uint `json:"total"`
	Removed uint `json:"removed"`
	Limit   uint `json:"limit"`
	Offset  uint `json:"offset"`
}
