package service

import (
	"context"
	"fmt"
	"net/http"

	errors "github.com/no-de-lab/nodelab-server/error"
	am "github.com/no-de-lab/nodelab-server/internal/auth/model"
	"github.com/no-de-lab/nodelab-server/internal/auth/util"
	"github.com/no-de-lab/nodelab-server/internal/domain"

	"gopkg.in/guregu/null.v4"
	"gopkg.in/jeevatkm/go-model.v1"
)

// AuthService business logic for auth
type AuthService struct {
	userService    domain.UserService
	authRepository domain.AuthRepository
}

// NewAuthService return new AuthService instance
func NewAuthService(userService domain.UserService, authRepository domain.AuthRepository) domain.AuthService {
	return &AuthService{
		userService,
		authRepository,
	}
}

// Login login process & issue jwt
func (as *AuthService) Login(ctx context.Context, form *am.LoginModel) (err error) {
	return nil
}

// SignupSocial signup a user by social account
func (as *AuthService) SignupSocial(ctx context.Context, user *am.SignupSocialModel) error {
	return nil
}

// SignupEmail signup a user by email and password
func (as *AuthService) SignupEmail(ctx context.Context, user *am.SignupEmailModel) error {
	userAcc, err := as.authRepository.FindAccountByEmail(ctx, user.Email)
	if userAcc != nil {
		return errors.NewBusinessError("User already exists", fmt.Errorf("User already exists"), http.StatusConflict)
	}

	if err != nil {
		return errors.NewInternalError("Failed to create user", err, http.StatusInternalServerError)
	}

	var userAccountModel domain.UserAccount
	errs := model.Copy(&userAccountModel, user)
	if errs != nil {
		return errors.NewInternalError("Failed to copy account model", errs[0], http.StatusInternalServerError)
	}

	hashedPassword, err := util.HashedPassword(user.Password.String)
	if err != nil {
		return errors.NewInternalError("Failed to has password", err, http.StatusInternalServerError)
	}
	userAccountModel.Password = null.NewString(hashedPassword, true)

	err = as.authRepository.CreateUserByEmail(ctx, &userAccountModel)
	if err != nil {
		return errors.NewInternalError("Failed to create user in DB", err, http.StatusInternalServerError)
	}

	return nil
}

// SocialLogin signup or login for social user
func (a *AuthService) SocialLogin() error {
	return nil
}
