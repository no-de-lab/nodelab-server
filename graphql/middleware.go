package graphql

import (
	"context"

	"github.com/no-de-lab/nodelab-server/graphql/resolver"

	"github.com/labstack/echo/v4"
)

// EchoContextToContextMiddleware wraps echo context to standard context
func EchoContextToContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), resolver.EchoCtxKey{}, c)
		c.SetRequest(c.Request().WithContext(ctx))
		cc := &EchoContext{c, ctx}
		return next(cc)
	}
}
