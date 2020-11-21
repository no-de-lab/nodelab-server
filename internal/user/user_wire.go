package user

import (
	"github.com/google/wire"
	"github.com/no-de-lab/nodelab-server/internal/user/delivery/http"
	"github.com/no-de-lab/nodelab-server/internal/user/repository"
	"github.com/no-de-lab/nodelab-server/internal/user/service"
)

var UserSet = wire.NewSet(repository.NewUserRepository, service.NewUserService, http.NewUserHandler)
