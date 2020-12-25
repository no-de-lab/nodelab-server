package service

import (
	"context"

	e "github.com/no-de-lab/nodelab-server/internal/auth/error"
	am "github.com/no-de-lab/nodelab-server/internal/auth/model"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	um "github.com/no-de-lab/nodelab-server/internal/user/model"
)

// AuthService business logic for auth
type AuthService struct {
	userService domain.UserService
}

// NewAuthService return new AuthService instance
func NewAuthService(service domain.UserService) domain.AuthService {
	return &AuthService{
		service,
	}
}

// Login login process & issue jwt
func (a *AuthService) Login(ctx context.Context, form *am.LoginModel) (err error) {
	return nil
}

// Signup user signup process
func (a *AuthService) Signup(ctx context.Context, user *um.CreateUserModel) error {
	existsUser, err := a.userService.FindByEmail(ctx, user.Email)

	if err != nil {
		return e.ErrUserNotFound.SetInternal(err)
	}

	if existsUser != nil {
		return e.ErrUserAlreadyExists
	}

	err = a.userService.CreateUser(ctx, user)

	if err != nil {
		return e.ErrUserCreate.SetInternal(err)
	}

	return nil
}

// SocialLogin signup or login for social user
func (a *AuthService) SocialLogin() (err error) {
	return nil
}
