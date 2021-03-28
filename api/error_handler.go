package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	e "github.com/no-de-lab/nodelab-server/error"
	log "github.com/sirupsen/logrus"
)

func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if echoError, ok := err.(*echo.HTTPError); ok {
		code = echoError.Code
		if echoError.Internal != nil {
			log.Error(echoError.Internal)
		}
		_ = c.JSON(code, echoError)
		return
	}

	if businessError, ok := err.(*e.BusinessError); ok {
		log.Error(businessError.Internal)
		_ = c.JSON(code, err)
		return
	}

	log.Error(err)
	_ = c.JSON(code, echo.ErrInternalServerError.Message)
}
