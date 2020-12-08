package healthcheck

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthCheckHandler struct { }

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) SetupRoutes(e *echo.Echo) {
	e.GET("/l7/monitor", h.HealthCheck)
}

func (h *HealthCheckHandler) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}