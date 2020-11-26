package api

import (
	"github.com/labstack/echo/v4"
	e "github.com/no-de-lab/nodelab-server/error"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if echoError, ok := err.(*echo.HTTPError); ok {
		code = echoError.Code
		if echoError.Internal != nil {
			logrus.Error(echoError.Internal)
		}
		_ = c.JSON(code, echoError)
		return
	}

	if businessError, ok := err.(*e.BusinessError); ok {
		logrus.Error(businessError.Internal)
		_ = c.JSON(code, err)
		return
	}

	logrus.Error(err)
	_ = c.JSON(code, echo.ErrInternalServerError.Message)
}
