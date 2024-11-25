package errors

import "errors"

const (
	DefaultMessage = "something went wrong"
)

type Error struct {
	InnerErr error `json:"-"`

	Message string `json:"message"`
	Code    *Code  `json:"code"`
	Level   *Level `json:"level"`
	Fields  Fields `json:"fields,omitempty"`
}

func GetMessage(err error) string {
	var (
		message = DefaultMessage
	)

	for err != nil {
		if e, ok := err.(*Error); ok {
			if e.Message != "" {
				message = e.Message
				break
			}
		}

		err = Unwrap(err)
	}
	return message
}

// ErrorOptions type.
type ErrorOptions func(*Error)

func WithFieldsOption(fields Fields) ErrorOptions {
	return func(ce *Error) {
		ce.Fields = fields
	}
}

// WithCodeOption options for Error construct.
func WithCodeOption(code Code) ErrorOptions {
	return func(ce *Error) {
		ce.Code = &code
	}
}

// WithLevelOption options for Error construct.
func WithLevelOption(level Level) ErrorOptions {
	return func(ce *Error) {
		ce.Level = &level
	}
}

func applyErrorOptions(err *Error, opts ...ErrorOptions) *Error {
	for _, opt := range opts {
		opt(err)
	}
	return err
}

// New error construct.
func New(message string, opts ...ErrorOptions) error {
	return applyErrorOptions(
		&Error{
			InnerErr: errors.New(message),
			Message:  message,
			Level:    nil,
			Code:     nil,
			Fields:   Fields{},
		},
		opts...,
	)
}

func (e *Error) Error() string {
	return e.InnerErr.Error()
}

func (e *Error) Unwrap() error {
	return e.InnerErr
}

func WithMessage(err error, message string) error {
	return &Error{
		InnerErr: err,
		Message:  message,
	}
}

func WithFields(err error, fields ...Field) error {
	return &Error{
		InnerErr: err,
		Fields:   fields,
	}
}
