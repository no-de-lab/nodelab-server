package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	"github.com/no-de-lab/nodelab-server/internal/user/model"
)

type AuthHandler struct {
	authService domain.AuthService
}

func NewAuthHandler(service domain.AuthService) *AuthHandler {
	return &AuthHandler{
		service,
	}
}

func (a *AuthHandler) SetupRoutes(e *echo.Echo) {
	authGroup := e.Group("/auth")
	authGroup.POST("/signup", a.Signup)
}

func (a *AuthHandler) Signup(c echo.Context) (err error) {
	var user model.CreateUser

	// context := c.Request().Context()

	if err = c.Bind(&user); err != nil {
		return
	}

	// err = a.authService.SignupSocial(context, &user)

	if err != nil {
		// if errors.Is(err, e.ErrUserAlreadyExists) {
		// 	return echo.NewHTTPError(http.StatusConflict, e.ErrUserAlreadyExists.Error())
		// }
		return
	}

	return c.NoContent(http.StatusOK)
}
