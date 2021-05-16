package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	echo "github.com/labstack/echo/v4"
	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
	"github.com/no-de-lab/nodelab-server/internal/auth/delivery/middleware"
	"github.com/no-de-lab/nodelab-server/internal/auth/util"
)

func (r *mutationResolver) SignupEmail(ctx context.Context, email string, password string) (*gqlschema.Auth, error) {
	return r.ar.SignupEmail(ctx, email, password)
}

func (r *mutationResolver) LoginEmail(ctx context.Context, email string, password string) (*gqlschema.Auth, error) {
	return r.ar.LoginEmail(ctx, email, password)
}

func (r *mutationResolver) LoginSocial(ctx context.Context, provider gqlschema.Provider, accessToken string, email string) (*gqlschema.Auth, error) {
	return r.ar.LoginSocial(ctx, provider, accessToken, email)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, email string, input *gqlschema.UpdateUserInput) (*gqlschema.User, error) {
	c := ctx.Value(EchoCtxKey{})

	ec, ok := c.(echo.Context)
	if !ok {
		err := fmt.Errorf("failed to convert context to echo context")
		return nil, err
	}

	payload, ok := ec.Get(middleware.UserPayloadCtxKey).(*util.Payload)
	if !ok {
		err := fmt.Errorf("failed to get user payload from context")
		return nil, err
	}

	return r.ur.UpdateUser(ctx, payload.Email, input)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, email string) (string, error) {
	return r.ur.DeleteUser(ctx, email)
}

// Mutation returns gqlschema.MutationResolver implementation.
func (r *Resolver) Mutation() gqlschema.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
