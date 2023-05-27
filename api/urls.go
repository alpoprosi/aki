package api

import "github.com/labstack/echo"

type ApiURL struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

func (h *ApiHandler) PrivateURLs() []ApiURL {
	return []ApiURL{
		{
			Method:  "GET",
			Path:    "me",
			Handler: h.Me,
		},
	}
}

func (h *ApiHandler) PublicURLs() []ApiURL {
	return []ApiURL{
		{
			Method:  "GET",
			Path:    "health-check",
			Handler: h.Health,
		},
	}
}

func (h *ApiHandler) TxURLs() []ApiURL {
	return []ApiURL{}
}
