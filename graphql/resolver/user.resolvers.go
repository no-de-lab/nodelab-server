package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
)

func (r *userResolver) Studies(ctx context.Context, obj *gqlschema.User) (*gqlschema.StudyConnection, error) {
	fmt.Print("Test")
	panic(fmt.Errorf("not implemented"))
}

// User returns gqlschema.UserResolver implementation.
func (r *Resolver) User() gqlschema.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
