package routes

import (
	"github.com/Aldiwildan77/rust-notifier-api/api/gateways/handlers/http"
	"github.com/labstack/echo/v4"
)

type RouterProto interface {
	Init(e *echo.Echo)
}

type Route struct {
	Handler http.HandlerProto
}

func NewRoute(handler http.HandlerProto) RouterProto {
	return &Route{
		Handler: handler,
	}
}

func (ur *Route) Init(e *echo.Echo) {
	g := e.Group("/api")

	r := g.Group("/notifier")
	r.GET("/", ur.Handler.GetUpdater)
}
