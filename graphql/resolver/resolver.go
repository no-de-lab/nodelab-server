package resolver

import (
	ag "github.com/no-de-lab/nodelab-server/internal/auth/delivery/graphql"
	ug "github.com/no-de-lab/nodelab-server/internal/user/delivery/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver a root resolver
type Resolver struct {
	ur *ug.UserResolver
	ar *ag.AuthResolver
}

// NewResolver return new root resolver instance
func NewResolver(
	userResolver *ug.UserResolver,
	authResolver *ag.AuthResolver,
) *Resolver {
	return &Resolver{
		ur: userResolver,
		ar: authResolver,
	}
}
