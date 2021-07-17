package resolver

import (
	ag "github.com/no-de-lab/nodelab-server/internal/auth/delivery/graphql"
	sg "github.com/no-de-lab/nodelab-server/internal/study/delivery/graphql"
	ug "github.com/no-de-lab/nodelab-server/internal/user/delivery/graphql"
)

// EchoCtxKey is the key to retrieve echo context from resolvers
type EchoCtxKey struct{}

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver a root resolver
type Resolver struct {
	ur *ug.UserResolver
	ar *ag.AuthResolver
	sr *sg.StudyResolver
}

// NewResolver return new root resolver instance
func NewResolver(
	userResolver *ug.UserResolver,
	authResolver *ag.AuthResolver,
	studyResolver *sg.StudyResolver,
) *Resolver {
	return &Resolver{
		ur: userResolver,
		ar: authResolver,
		sr: studyResolver,
	}
}
