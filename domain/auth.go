package domain

import (
	"context"
	authDto "github.com/no-de-lab/nodelab-server/auth/dto"
	"github.com/no-de-lab/nodelab-server/user/dto"
)

type AuthService interface {
	Login(ctx context.Context, form *authDto.LoginDto) (err error)
	Signup(ctx context.Context, user *dto.CreateUserDto) (err error)
	// TODO: add detail
	SocialLogin() (err error)
}
