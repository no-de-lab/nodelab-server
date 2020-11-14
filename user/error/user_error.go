package error

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserCreate        = errors.New("user create error")
	ErrUserNotFound      = errors.New("No members found")
)
