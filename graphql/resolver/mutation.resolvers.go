package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
)

func (r *mutationResolver) SignupEmail(ctx context.Context, email string, password string) (*gqlschema.Auth, error) {
	return r.ar.SignupEmail(ctx, email, password)
}

func (r *mutationResolver) LoginEmail(ctx context.Context, email string, password string) (*gqlschema.Auth, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns gqlschema.MutationResolver implementation.
func (r *Resolver) Mutation() gqlschema.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
