package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
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
	panic(fmt.Errorf("not implemented"))
}

// Query returns gqlschema.QueryResolver implementation.
func (r *Resolver) Query() gqlschema.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Users(ctx context.Context) ([]*gqlschema.User, error) {
	panic(fmt.Errorf("not implemented"))
}
