package service

import (
	"context"

	am "github.com/no-de-lab/nodelab-server/internal/auth/model"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	e "github.com/no-de-lab/nodelab-server/internal/user/error"
	um "github.com/no-de-lab/nodelab-server/internal/user/model"
	log "github.com/sirupsen/logrus"
)

type AuthService struct {
	userService domain.UserService
}

func NewAuthService(service domain.UserService) domain.AuthService {
	return &AuthService{
		service,
	}
}

func (a *AuthService) Login(ctx context.Context, form *am.LoginModel) (err error) {
	return nil
}

func (a *AuthService) Signup(ctx context.Context, user *um.CreateUserModel) error {
	existsUser, err := a.userService.FindByEmail(ctx, user.Email)
	if err != nil {
		log.Error(err)
		return err
	}

	if existsUser != nil {
		return e.ErrUserAlreadyExists
	}

	err = a.userService.CreateUser(ctx, user)

	if err != nil {
		log.Error(err)
		return e.ErrUserCreate
	}

	return nil
}

func (a *AuthService) SocialLogin() (err error) {
	return nil
}
