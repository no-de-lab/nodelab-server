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
