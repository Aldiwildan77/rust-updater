package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Aldiwildan77/rust-notifier-api/api/domains"
	"github.com/Aldiwildan77/rust-notifier-api/api/gateways/presenters"
	"github.com/Aldiwildan77/rust-notifier-api/config"
	"github.com/ReneKroon/ttlcache/v2"
	"github.com/labstack/echo/v4"
)

var (
	ttl = 5 * time.Minute
	cn  = "old-data"
)

type HandlerProto interface {
	GetUpdater(c echo.Context) error
}

type handler struct {
	Conn  *config.Config
	Cache ttlcache.SimpleCache
}

func NewHandler(cfg *config.Config, cache ttlcache.SimpleCache) HandlerProto {
	return &handler{
		Conn:  cfg,
		Cache: cache,
	}
}

func (hh *handler) GetUpdater(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

	if val, err := hh.Cache.Get(cn); err != ttlcache.ErrNotFound {
		return c.JSON(http.StatusOK, presenters.GameResponse{
			UpdaterResponse: val.(presenters.UpdaterResponse),
		})
	}

	dm := domains.NewUsecase(context.Background())
	res, err := dm.GetListUpdates()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, presenters.GameResponse{})
	}

	hh.Cache.SetTTL(time.Duration(ttl))
	hh.Cache.Set(cn, res)

	return c.JSON(http.StatusOK, presenters.GameResponse{
		UpdaterResponse: res,
	})
}
