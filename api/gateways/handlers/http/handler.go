package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Aldiwildan77/rust-notifier-api/api/domains"
	"github.com/Aldiwildan77/rust-notifier-api/api/gateways/presenters"
	"github.com/Aldiwildan77/rust-notifier-api/config"
	"github.com/ReneKroon/ttlcache/v2"
	"github.com/labstack/echo/v4"
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

	if val, err := hh.Cache.Get("old-data"); err != ttlcache.ErrNotFound {
		return c.JSON(http.StatusOK, presenters.GameResponse{
			UpdaterResponse: val.(presenters.UpdaterResponse),
		})
	}

	dm := domains.NewUsecase()
	res, err := dm.GetListPayment()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, presenters.GameResponse{})
	}

	hh.Cache.SetTTL(time.Duration(5 * time.Minute))
	hh.Cache.Set("old-data", res)

	return c.JSON(http.StatusOK, presenters.GameResponse{
		UpdaterResponse: res,
	})
}
