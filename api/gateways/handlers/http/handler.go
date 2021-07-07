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
		return c.JSON(http.StatusOK, presenters.Response{
			Status:  true,
			Data:    val,
			Message: "Notifier Fetched",
		})
	}

	dm := domains.NewUsecase()
	res, err := dm.GetListPayment()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, presenters.Response{
			Status:  false,
			Message: "Failed to get data",
		})
	}

	hh.Cache.SetTTL(time.Duration(3 * time.Second))
	hh.Cache.Set("old-data", res)

	return c.JSON(http.StatusOK, presenters.Response{
		Status:  true,
		Data:    res,
		Message: "Notifier Fetched",
	})
}
