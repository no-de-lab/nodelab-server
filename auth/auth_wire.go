package auth

import (
	"github.com/google/wire"
	"github.com/no-de-lab/nodelab-server/auth/delivery/http"
	"github.com/no-de-lab/nodelab-server/auth/service"
)

var AuthSet = wire.NewSet(service.NewAuthService, http.NewAuthHandler)
