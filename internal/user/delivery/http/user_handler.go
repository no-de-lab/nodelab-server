package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	um "github.com/no-de-lab/nodelab-server/internal/user/model"
	"gopkg.in/jeevatkm/go-model.v1"
)

type UserHandler struct {
	UserService domain.UserService
}

func NewUserHandler(us domain.UserService) *UserHandler {
	return &UserHandler{
		UserService: us,
	}
}

func (h *UserHandler) SetupRoutes(e *echo.Echo) {
	userRouter := e.Group("/users")
	userRouter.GET("/:id", h.GetUserInfo)
}

func (h *UserHandler) GetUserInfo(c echo.Context) error {
	var user *domain.User
	var userInfo um.UserInfo

	context := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		return echo.ErrBadRequest
	}

	if user, err = h.UserService.FindByID(context, int(id)); err != nil {
		return err
	}

	if user == nil {
		// return userError.ErrUserNotFound
		return nil
	}

	errs := model.Copy(&userInfo, user)

	if errs != nil {
		return errs[0]
	}

	return c.JSON(http.StatusOK, user)
}
