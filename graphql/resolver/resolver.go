package resolver

import "github.com/no-de-lab/nodelab-server/internal/user/delivery/graphql"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver a root resolver
type Resolver struct {
	ur *graphql.UserResolver
}

// NewResolver return new root resolver instance
func NewResolver(
	userResolver *graphql.UserResolver,
) *Resolver {
	return &Resolver{
		ur: userResolver,
	}
}
