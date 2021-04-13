package auth

import (
	"github.com/google/wire"
	"github.com/no-de-lab/nodelab-server/internal/auth/delivery/graphql"
	"github.com/no-de-lab/nodelab-server/internal/auth/delivery/http"
	"github.com/no-de-lab/nodelab-server/internal/auth/model"
	"github.com/no-de-lab/nodelab-server/internal/auth/repository"
	"github.com/no-de-lab/nodelab-server/internal/auth/service"
	"github.com/no-de-lab/nodelab-server/internal/auth/util"
)

var AuthSet = wire.NewSet(util.NewJWTMaker, repository.NewAuthRepository, service.NewAuthService, http.NewAuthHandler, model.NewValidator, graphql.NewAuthResolver)
