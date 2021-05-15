package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
)

func (r *userResolver) Studies(ctx context.Context, obj *gqlschema.User) (*gqlschema.StudyConnection, error) {
	return &gqlschema.StudyConnection{}, nil
}

// User returns gqlschema.UserResolver implementation.
func (r *Resolver) User() gqlschema.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
