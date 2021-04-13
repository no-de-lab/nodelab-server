// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/no-de-lab/nodelab-server/api/healthcheck"
	"github.com/no-de-lab/nodelab-server/config"
	"github.com/no-de-lab/nodelab-server/container"
	"github.com/no-de-lab/nodelab-server/db"
	"github.com/no-de-lab/nodelab-server/graphql/resolver"
	"github.com/no-de-lab/nodelab-server/internal/auth"
	"github.com/no-de-lab/nodelab-server/internal/user"
)

// MainSet all instance set (service, resolver, handler ... etc)
var MainSet = wire.NewSet(healthcheck.NewHealthCheckHandler, config.LoadConfig, db.NewDatabase, auth.AuthSet, user.UserSet, container.NewDIContainer, resolver.NewResolver)

// InitializeDIContainer return instance bean container
func InitializeDIContainer() *container.DIContainer {
	wire.Build(MainSet)
	return nil
}

// InitializeResolver return root resolver
func InitializeResolver() *resolver.Resolver {
	wire.Build(MainSet)
	return nil
}
