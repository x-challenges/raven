package errors

type Code = string

const (
	InternalError      Code = "Internal"
	UnimplementedError Code = "Unimplemented"
)

const (
	UnauthenticatedError Code = "Unauthenticated"
	ForbiddenError       Code = "Forbidden"
)

const (
	InvalidRequestError Code = "InvalidRequest"
	NotFoundError       Code = "NotFound"
	AlreadyExistError   Code = "AlreadyExist"
	TooManyRequestError Code = "TooManyRequest"
)

var (
	AsInternal = WithCodeOption(InternalError)
)

func GetCode(err error) Code {
	var (
		code = InternalError
	)

	for err != nil {
		if e, ok := err.(*Error); ok {
			if e.Code != nil {
				code = *e.Code
				break
			}
		}

		err = Unwrap(err)
	}
	return code
}
