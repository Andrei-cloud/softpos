package softpos

import "errors"

var (
	ErrSuccess        error = errors.New("success")
	ErrCreated        error = errors.New("created")
	ErrIncorrect      error = errors.New("incorrect properties")
	ErrEntityNotFound error = errors.New("entity not found")
	ErrAcqNotExist    error = errors.New("acquirer does not exists or incorrect properties")
	ErrIvalidToken    error = errors.New("authentication token validation error")
	ErrNoPermission   error = errors.New("do not have permission")
	ErrConflict       error = errors.New("conflict")
	ErrUnknown        error = errors.New("unknown error")
)
