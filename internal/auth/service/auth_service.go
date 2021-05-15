package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	e "github.com/no-de-lab/nodelab-server/error"
	am "github.com/no-de-lab/nodelab-server/internal/auth/model"
	"github.com/no-de-lab/nodelab-server/internal/auth/provider"
	"github.com/no-de-lab/nodelab-server/internal/auth/util"
	"github.com/no-de-lab/nodelab-server/internal/domain"

	"gopkg.in/guregu/null.v4"
	"gopkg.in/jeevatkm/go-model.v1"

	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
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

// SignupSocial signup a user by social account in LoginSocial
func (as *AuthService) signupSocial(ctx context.Context, user *am.LoginSocialModel) (string, error) {

	var userAccountModel domain.UserAccount
	errs := model.Copy(&userAccountModel, user)
	if errs != nil {
		return "", e.NewInternalError("failed to copy account model", errs[0], http.StatusInternalServerError)
	}

	err := as.authRepository.CreateUserBySocial(ctx, &userAccountModel)
	if err != nil {
		return "", e.NewInternalError("failed to create user in DB", err, http.StatusInternalServerError)
	}

	token, err := as.jwtMaker.CreateToken(user.Email, 168*time.Hour)
	if err != nil {
		return "", fmt.Errorf("failed to make jwt token")
	}

	return token, nil
}

// SignupEmail signup a user by email and password
func (as *AuthService) SignupEmail(ctx context.Context, user *am.SignupEmailModel) (string, error) {
	userAcc, err := as.authRepository.FindAccountByEmail(ctx, user.Email)
	if userAcc != nil {
		return "", e.NewBusinessError("user already exists", fmt.Errorf("user already exists"), http.StatusConflict)
	}

	if !errors.Is(err, sql.ErrNoRows) && err != nil {
		return "", e.NewInternalError("failed to create user", err, http.StatusInternalServerError)
	}

	var userAccountModel domain.UserAccount
	errs := model.Copy(&userAccountModel, user)
	if errs != nil {
		return "", e.NewInternalError("failed to copy account model", errs[0], http.StatusInternalServerError)
	}

	hashedPassword, err := util.HashedPassword(user.Password.String)
	if err != nil {
		return "", e.NewInternalError("failed to has password", err, http.StatusInternalServerError)
	}
	userAccountModel.Password = null.NewString(hashedPassword, true)

	err = as.authRepository.CreateUserByEmail(ctx, &userAccountModel)
	if err != nil {
		return "", e.NewInternalError("failed to create user in DB", err, http.StatusInternalServerError)
	}

	token, err := as.jwtMaker.CreateToken(user.Email, 168*time.Hour)
	if err != nil {
		return "", fmt.Errorf("failed to make jwt token")
	}

	return token, nil
}

// LoginSocial logins social user
func (as *AuthService) LoginSocial(ctx context.Context, user *am.LoginSocialModel) (string, error) {
	userAcc, err := as.authRepository.FindAccountByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}

	var providerID string
	switch user.Provider {
	case gqlschema.ProviderKakao:
		kakaoId, err := provider.LoginKakao(user.AccessToken)
		if err != nil {
			return "", err
		}
		providerID = kakaoId
	case gqlschema.ProviderGoogle:
		tokenInfo, err := provider.LoginGoogle(user.AccessToken)

		if err != nil {
			return "", err
		}

		providerID = tokenInfo.Email
	default:
		return "", fmt.Errorf("invalid Provider error")
	}

	// no user -> make account
	if userAcc == nil {
		return as.signupSocial(ctx, user)
	}

	if userAcc.ProviderID.String != providerID && gqlschema.Provider(userAcc.Provider.String) != user.Provider {
		return "", fmt.Errorf("Login Failed")
	}

	token, err := as.jwtMaker.CreateToken(user.Email, 168*time.Hour)
	if err != nil {
		return "", err
	}

	return token, nil
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
