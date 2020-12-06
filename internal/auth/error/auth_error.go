package error

import error "github.com/no-de-lab/nodelab-server/error"

var (
	ErrUserAlreadyExists = error.NewBusinessError("user already exists")
	ErrUserCreate        = error.NewBusinessError("user create error")
	ErrUserNotFound      = error.NewBusinessError("No members found")
)
