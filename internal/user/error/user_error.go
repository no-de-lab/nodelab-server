package error

import error "github.com/no-de-lab/nodelab-server/error"

var (
	ErrUserNotFound = error.NewBusinessError("No members found")
)
