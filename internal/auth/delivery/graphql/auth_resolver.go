package graphql

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
	ae "github.com/no-de-lab/nodelab-server/internal/auth/error"
	am "github.com/no-de-lab/nodelab-server/internal/auth/model"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	log "github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v4"
)

// AuthResolver auth resolver for graphql
type AuthResolver struct {
	Validator   validator.Validate
	AuthService domain.AuthService
}

// NewAuthResolver return new auth resolver instance
func NewAuthResolver(validator validator.Validate, authService domain.AuthService) *AuthResolver {
	return &AuthResolver{
		Validator:   validator,
		AuthService: authService,
	}
}

// SignupEmail signs up the user by email
func (ar *AuthResolver) SignupEmail(ctx context.Context, email, password string) (*gqlschema.Auth, error) {
	signupModel := &am.SignupEmailModel{
		Email:    email,
		Password: null.NewString(password, true),
	}

	err := ar.Validator.Struct(signupModel)
	if err != nil {
		log.WithError(err).Errorf("Bad request from user")
		return nil, ae.NewGraphqlError(ctx, err.Error(), http.StatusBadRequest)
	}

	token, err := ar.AuthService.SignupEmail(ctx, signupModel)
	if err != nil {
		log.WithError(err).Errorf("Failed to create user by email, password")
		return nil, ae.NewGraphqlError(ctx, err.Error(), 0)
	}

	var gqlAuth gqlschema.Auth
	gqlAuth.Email = email
	gqlAuth.Token = token

	return &gqlAuth, nil
}

// LoginEmail logins the user by email
func (ar *AuthResolver) LoginEmail(ctx context.Context, email, password string) (*gqlschema.Auth, error) {
	token, err := ar.AuthService.LoginEmail(ctx, email, password)
	if err != nil {
		log.Error(err)
		return nil, ae.NewGraphqlError(ctx, err.Error(), http.StatusBadRequest)
	}

	gqlAuth := gqlschema.Auth{
		Email: email,
		Token: token,
	}

	return &gqlAuth, nil
}

// LoginSocial logins the user by social account using email
func (ar *AuthResolver) LoginSocial(ctx context.Context, provider gqlschema.Provider, accessToken string, email string) (*gqlschema.Auth, error) {
	loginModel := &am.LoginSocialModel{
		Provider:    provider,
		Email:       email,
		AccessToken: accessToken,
	}
	token, err := ar.AuthService.LoginSocial(ctx, loginModel)
	if err != nil {
		log.Error(err)
		return nil, ae.NewGraphqlError(ctx, err.Error(), http.StatusBadRequest)
	}

	gqlAuth := gqlschema.Auth{
		Email: email,
		Token: token,
	}

	return &gqlAuth, nil
}
