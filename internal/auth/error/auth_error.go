package error

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var (
// ErrUserAlreadyExists = errors.NewBusinessError("user already exists")
// ErrUserCreate        = errors.NewBusinessError("user create error")
// ErrUserNotFound      = errors.NewBusinessError("No members found")
)

func NewGraphqlError(ctx context.Context, message string, code int) error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: message,
		Extensions: map[string]interface{}{
			"code": code,
		},
	}
}

func ErrUserAlreadyExist(ctx context.Context, err error) error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: "User already exists",
		Extensions: map[string]interface{}{
			"code": http.StatusConflict,
		},
	}
}

func ErrInvalidUserInput(ctx context.Context, err error) error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: fmt.Sprintf("User input invalid: %v", err),
		Extensions: map[string]interface{}{
			"code": http.StatusBadRequest,
		},
	}
}

func ErrInternalServer(ctx context.Context, err error) error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: fmt.Sprintf("Internal server error: %v", err),
		Extensions: map[string]interface{}{
			"code": http.StatusInternalServerError,
		},
	}
}
