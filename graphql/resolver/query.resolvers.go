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

func (r *queryResolver) Studies(ctx context.Context) (*gqlschema.StudyConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Study(ctx context.Context, id string) (*gqlschema.Study, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*gqlschema.User, error) {
	return r.ur.User(ctx, id)
}

func (r *queryResolver) Me(ctx context.Context) (*gqlschema.User, error) {
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

	return r.ur.Me(ctx, payload.Email)
}

// Query returns gqlschema.QueryResolver implementation.
func (r *Resolver) Query() gqlschema.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
