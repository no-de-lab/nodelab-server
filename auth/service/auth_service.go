package service

import (
	"context"
	authDto "github.com/no-de-lab/nodelab-server/auth/dto"
	"github.com/no-de-lab/nodelab-server/domain"
	userDto "github.com/no-de-lab/nodelab-server/user/dto"
	e "github.com/no-de-lab/nodelab-server/user/error"
)

type AuthService struct {
	userService domain.UserService
}

func NewAuthService(service domain.UserService) domain.AuthService {
	return &AuthService{
		service,
	}
}

func (a *AuthService) Login(ctx context.Context, form *authDto.LoginDto) (err error) {
	return nil
}

func (a *AuthService) Signup(ctx context.Context, user *userDto.CreateUserDto) error {
	existsUser, err := a.userService.FindByEmail(ctx, user.Email.String)

	if existsUser != nil {
		return e.ErrUserAlreadyExists
	}

	err = a.userService.CreateUser(ctx, user)

	if err != nil {
		return e.ErrUserCreate
	}

	return nil
}

func (a *AuthService) SocialLogin() (err error) {
	return nil
}
