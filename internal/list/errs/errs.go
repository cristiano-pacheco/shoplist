package errs

import "errors"

var ErrResourceDoesNotBelongToUser = errors.New("the resource does not belong to the user")
