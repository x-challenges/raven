package errors

var (
	ErrInternal      = New("internal", AsInternal)
	ErrUnimplemented = New("unimplemented", AsInternal)
)

var (
	ErrUnauthenticated = New("unauthenticated", WithCodeOption(UnauthenticatedError), AsClient)
	ErrForbidden       = New("forbidden", WithCodeOption(ForbiddenError), AsClient)
)

var (
	ErrNotFound       = New("not found", WithCodeOption(NotFoundError), AsClient)
	ErrAlreadyExists  = New("already exist", WithCodeOption(AlreadyExistError), AsClient)
	ErrInvalidRequest = New("invalid request", WithCodeOption(InvalidRequestError), AsClient)
	ErrToManyRequest  = New("too many request", WithCodeOption(TooManyRequestError), AsClient)
)
