package errors

import (
	"errors"
	"fmt"
)

var (
	Is     = errors.Is
	As     = errors.As
	Join   = errors.Join
	Wrap   = fmt.Errorf
	Unwrap = errors.Unwrap
)
