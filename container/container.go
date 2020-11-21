package container

import (
	"github.com/no-de-lab/nodelab-server/api"
	ah "github.com/no-de-lab/nodelab-server/internal/auth/delivery/http"
	uh "github.com/no-de-lab/nodelab-server/internal/user/delivery/http"
)

type DIContainer struct {
	Handlers []api.ApiHandler
}

// singleton expose
var Container *DIContainer

func NewDIContainer(
	userHandler *uh.UserHandler,
	authHandler *ah.AuthHandler,
) *DIContainer {
	Container = &DIContainer{
		[]api.ApiHandler{userHandler, authHandler},
	}
	return Container
}
