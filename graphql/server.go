package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/no-de-lab/nodelab-server/config"
	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
	"github.com/no-de-lab/nodelab-server/graphql/resolver"
	"github.com/no-de-lab/nodelab-server/internal/auth/delivery/middleware"
	"github.com/no-de-lab/nodelab-server/internal/auth/util"
)

const graphqlEndpoint = "/graphql"

// EchoContext is a custom context to wrap echo context with standard context
type EchoContext struct {
	echo.Context
	ctx context.Context
}

// SetupGraphQL setup graphql handler, playground & root resolver
func SetupGraphQL(e *echo.Echo, resolver *resolver.Resolver, cfg *config.Configuration) {

	graphqlHandler := handler.NewDefaultServer(gqlschema.NewExecutableSchema(gqlschema.Config{Resolvers: resolver}))
	playgroundHandler := playground.Handler("GraphQL", graphqlEndpoint)

	e.Use(middleware.Authorize(util.NewJWTMaker(cfg)))
	e.Use(EchoContextToContextMiddleware)
	e.Use(middleware.AddCORS)

	e.POST(graphqlEndpoint, func(c echo.Context) error {
		cc := c.(*EchoContext)
		graphqlHandler.ServeHTTP(cc.Response(), cc.Request())
		return nil
	})
	e.GET(graphqlEndpoint, func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}
