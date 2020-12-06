package api

import "github.com/labstack/echo/v4"

type ApiHandler interface {
	SetupRoutes(e *echo.Echo)
}
