package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

const uidCookieName = "user"

func (h *ApiHandler) AuthMiddleware(db *gorm.DB, loginURL string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			uid, err := c.Request().Cookie(uidCookieName)
			if err != nil {
				return err
			}

			if !h.checkUIDCookie(uid) {
				path := c.Request().URL.Path
				return c.Redirect(
					http.StatusTemporaryRedirect,
					fmt.Sprintf("%s?next=%s", loginURL, url.QueryEscape(path)),
				)
			}

			return next(c)
		}
	}
}

func (h *ApiHandler) checkUIDCookie(uid *http.Cookie) bool {
	if uid == nil {
		return false
	}

	u, err := h.app.GetUserByUID(uid.Value)
	if err != nil {
		log.Errorf("getting user: %w", err)

		return false
	}

	if u == nil {
		return false
	}

	return true
}
