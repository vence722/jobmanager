package errors

type Error struct {
	ErrCode    int
	ErrMessage string
	RawError   error
}

func NewError(errCode int, errMessage string, rawError error) *Error {
	return &Error{
		ErrCode:    errCode,
		ErrMessage: errMessage,
		RawError:   rawError,
	}
}
