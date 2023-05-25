package api

import (
	"aki/aki"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type ApiHandler struct {
	app aki.Aki
}

func New(db *gorm.DB) *ApiHandler {
	return &ApiHandler{
		app: *aki.New(db),
	}
}

func (h *ApiHandler) Me(c echo.Context) error {
	return nil
}

func (h *ApiHandler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
