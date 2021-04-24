package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	errors "github.com/no-de-lab/nodelab-server/error"
	am "github.com/no-de-lab/nodelab-server/internal/auth/model"
	"github.com/no-de-lab/nodelab-server/internal/auth/util"
	"github.com/no-de-lab/nodelab-server/internal/domain"

	"gopkg.in/guregu/null.v4"
	"gopkg.in/jeevatkm/go-model.v1"
)

// AuthService business logic for auth
type AuthService struct {
	jwtMaker       util.JWTMaker
	userService    domain.UserService
	authRepository domain.AuthRepository
}

// NewAuthService return new AuthService instance
func NewAuthService(jwtMaker util.JWTMaker, userService domain.UserService, authRepository domain.AuthRepository) domain.AuthService {
	return &AuthService{
		jwtMaker,
		userService,
		authRepository,
	}
}

// Login login process & issue jwt
func (as *AuthService) Login(ctx context.Context, form *am.LoginModel) (err error) {
	return nil
}

// SignupSocial signup a user by social account
func (as *AuthService) SignupSocial(ctx context.Context, user *am.SignupSocialModel) (string, error) {
	return "", nil
}

// SignupEmail signup a user by email and password
func (as *AuthService) SignupEmail(ctx context.Context, user *am.SignupEmailModel) (string, error) {
	userAcc, err := as.authRepository.FindAccountByEmail(ctx, user.Email)
	if userAcc != nil {
		return "", errors.NewBusinessError("User already exists", fmt.Errorf("User already exists"), http.StatusConflict)
	}

	if err != nil {
		return "", errors.NewInternalError("Failed to create user", err, http.StatusInternalServerError)
	}

	var userAccountModel domain.UserAccount
	errs := model.Copy(&userAccountModel, user)
	if errs != nil {
		return "", errors.NewInternalError("Failed to copy account model", errs[0], http.StatusInternalServerError)
	}

	hashedPassword, err := util.HashedPassword(user.Password.String)
	if err != nil {
		return "", errors.NewInternalError("Failed to has password", err, http.StatusInternalServerError)
	}
	userAccountModel.Password = null.NewString(hashedPassword, true)

	err = as.authRepository.CreateUserByEmail(ctx, &userAccountModel)
	if err != nil {
		return "", errors.NewInternalError("Failed to create user in DB", err, http.StatusInternalServerError)
	}

	token, err := as.jwtMaker.CreateToken(user.Email, 168*time.Hour)
	if err != nil {
		return "", fmt.Errorf("Failed to make jwt token")
	}

	return token, nil
}

// LoginSocial logins social user
func (as *AuthService) LoginSocial(ctx context.Context, email string) (string, error) {
	return "", nil
}

// LoginEmail logins email user and returns token
func (as *AuthService) LoginEmail(ctx context.Context, email, password string) (string, error) {
	user, err := as.authRepository.FindAccountByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", err
	}

	err = util.CheckPassword(password, user.Password.String)
	if err != nil {
		return "", err
	}

	token, err := as.jwtMaker.CreateToken(email, 168*time.Hour)
	if err != nil {
		return "", err
	}

	return token, nil
}
