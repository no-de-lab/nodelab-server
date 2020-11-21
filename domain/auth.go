package domain

import (
	"context"
	am "github.com/no-de-lab/nodelab-server/auth/model"
	"github.com/no-de-lab/nodelab-server/user/model"
)

type AuthService interface {
	Login(ctx context.Context, form *am.LoginModel) (err error)
	Signup(ctx context.Context, user *model.CreateUserModel) (err error)
	// TODO: add detail
	SocialLogin() (err error)
}
