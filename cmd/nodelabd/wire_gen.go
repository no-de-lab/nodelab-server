// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/no-de-lab/nodelab-server/auth"
	http2 "github.com/no-de-lab/nodelab-server/auth/delivery/http"
	service2 "github.com/no-de-lab/nodelab-server/auth/service"
	"github.com/no-de-lab/nodelab-server/config"
	"github.com/no-de-lab/nodelab-server/container"
	"github.com/no-de-lab/nodelab-server/db"
	"github.com/no-de-lab/nodelab-server/user"
	"github.com/no-de-lab/nodelab-server/user/delivery/http"
	"github.com/no-de-lab/nodelab-server/user/repository"
	"github.com/no-de-lab/nodelab-server/user/service"
)

// Injectors from wire.go:

func InitializeDIContainer() *container.DIContainer {
	configuration := config.LoadConfig()
	sqlxDB := db.NewDatabase(configuration)
	userRepository := repository.NewUserRepository(sqlxDB)
	userService := service.NewUserService(userRepository, configuration)
	userHandler := http.NewUserHandler(userService)
	authService := service2.NewAuthService(userService)
	authHandler := http2.NewAuthHandler(authService)
	diContainer := container.NewDIContainer(userHandler, authHandler)
	return diContainer
}

// wire.go:

var MainSet = wire.NewSet(auth.AuthSet, user.UserSet, config.LoadConfig, db.NewDatabase, container.NewDIContainer)
