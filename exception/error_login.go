package exception

type ErrorLogin struct {
	Message string
}

func NewErrorLogin(message string) *ErrorLogin {
	return &ErrorLogin{Message: message}
}

func (e ErrorLogin) Error() string {
	return e.Message
}
