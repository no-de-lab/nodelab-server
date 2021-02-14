package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
	"github.com/no-de-lab/nodelab-server/graphql/resolver"
)

const graphqlEndpoint = "/graphql"

// SetupGraphQL setup graphql handler, playground & root resolver
func SetupGraphQL(e *echo.Echo, resolver *resolver.Resolver) {

	graphqlHandler := handler.NewDefaultServer(gqlschema.NewExecutableSchema(gqlschema.Config{Resolvers: resolver}))
	playgroundHandler := playground.Handler("GraphQL", graphqlEndpoint)

	e.POST(graphqlEndpoint, func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	e.GET(graphqlEndpoint, func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}
