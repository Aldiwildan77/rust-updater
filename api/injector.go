package api

import (
	"github.com/Aldiwildan77/rust-notifier-api/api/gateways/handlers/http"
	"github.com/Aldiwildan77/rust-notifier-api/config"
	"github.com/ReneKroon/ttlcache/v2"
)

func ProvideHandler() http.HandlerProto {
	return http.NewHandler(&config.Cfg, ttlcache.NewCache())
}
