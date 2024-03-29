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

// CORSMiddleware adds CORS header to requests
func CORSMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
		// c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "http://3.36.166.101")
		c.Response().Header().Set("Vary", "Origin")
		c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		return next(c)
	}
}
