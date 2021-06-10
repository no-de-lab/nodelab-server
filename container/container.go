package container

import (
	"github.com/no-de-lab/nodelab-server/api"
	"github.com/no-de-lab/nodelab-server/api/healthcheck"
	"github.com/no-de-lab/nodelab-server/config"
	ah "github.com/no-de-lab/nodelab-server/internal/auth/delivery/http"
	sh "github.com/no-de-lab/nodelab-server/internal/study/delivery/http"
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
	studyHander *sh.StudyHandler,
	healthcheckHandler *healthcheck.HealthCheckHandler,
) *DIContainer {
	Container = &DIContainer{
		configuration,
		[]api.ApiHandler{userHandler, authHandler, healthcheckHandler},
	}
	return Container
}
