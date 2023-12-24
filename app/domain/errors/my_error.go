package errors

type IMyError interface {
	Error() string
	Original() string
	Code() uint16
}

type MyError struct {
	message  string
	original string
	code     uint16
}

func NewMyError(message, original string, code uint16) error {
	return &MyError{
		message:  message,
		original: original,
		code:     code,
	}
}

func (m *MyError) Error() string {
	return m.original
}

func (m *MyError) DisplayError() string {
	return m.message
}

func (m *MyError) Code() uint16 {
	return m.code
}
