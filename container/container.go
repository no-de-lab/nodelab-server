package container

import (
	"github.com/no-de-lab/nodelab-server/api"
	"github.com/no-de-lab/nodelab-server/config"
	ah "github.com/no-de-lab/nodelab-server/internal/auth/delivery/http"
	uh "github.com/no-de-lab/nodelab-server/internal/user/delivery/http"
)

type DIContainer struct {
	Config   *config.Configuration
	Handlers []api.ApiHandler
}

// singleton expose
var Container *DIContainer

func NewDIContainer(
	configuration *config.Configuration,
	userHandler *uh.UserHandler,
	authHandler *ah.AuthHandler,
) *DIContainer {
	Container = &DIContainer{
		configuration,
		[]api.ApiHandler{userHandler, authHandler},
	}
	return Container
}
