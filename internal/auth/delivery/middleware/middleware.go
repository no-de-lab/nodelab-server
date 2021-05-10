package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/no-de-lab/nodelab-server/internal/auth/util"
)

const (
	authTypeBearer = "bearer"
	authHeaderKey  = "Authorization"
)

// UserPayloadCtxKey is the key for payload passed along context
const UserPayloadCtxKey = "UserPayloadCtxKey"

// Authorize reads the authorization header and sets the user payload
func Authorize(jwtMaker util.JWTMaker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header[authHeaderKey]

			if len(authHeader) == 0 {
				return next(c)
			}

			fields := strings.Fields(authHeader[0])
			if len(fields) < 2 {
				err := errors.New("invalid authorization header")
				return err
			}

			authType := strings.ToLower(fields[0])
			if authType != authTypeBearer {
				err := fmt.Errorf("unsupported authorization type: %s", authType)
				return err
			}

			token := fields[1]
			payload, err := jwtMaker.VerifyToken(token)
			if err != nil {
				err := fmt.Errorf("invalid token")
				return err
			}

			c.Set(UserPayloadCtxKey, payload)
			return next(c)
		}
	}
}
