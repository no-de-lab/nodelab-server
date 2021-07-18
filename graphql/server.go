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
	// e.Use(EchoContextToContextMiddleware)
	e.Use(CORSMiddleware)
	// e.Use(CORSMiddlewareWrapper)

	// e.OPTIONS(graphqlEndpoint, addHeaders)
	e.POST(graphqlEndpoint, func(c echo.Context) error {
		cc := c.(*EchoContext)
		graphqlHandler.ServeHTTP(cc.Response(), cc.Request())
		return nil
	})
	e.GET(graphqlEndpoint, func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	e.OPTIONS(graphqlEndpoint, echo.MethodNotAllowedHandler)
}

// func addHeaders(c echo.Context) error {
// 	headers := c.Response().Header()
// 	headers.Add("Access-Control-Allow-Origin", "*")
// 	headers.Add("Vary", "Origin")
// 	headers.Add("Vary", "Access-Control-Request-Method")
// 	headers.Add("Vary", "Access-Control-Request-Headers")
// 	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
// 	headers.Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT")
// 	c.Response().WriteHeader(http.StatusOK)
// 	return nil
// }
