// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/no-de-lab/nodelab-server/auth"
	"github.com/no-de-lab/nodelab-server/config"
	"github.com/no-de-lab/nodelab-server/container"
	"github.com/no-de-lab/nodelab-server/db"
	"github.com/no-de-lab/nodelab-server/user"
)

var MainSet = wire.NewSet(auth.AuthSet, user.UserSet, config.LoadConfig, db.NewDatabase, container.NewDIContainer)

func InitializeDIContainer() *container.DIContainer {
	wire.Build(MainSet)
	return nil
}
