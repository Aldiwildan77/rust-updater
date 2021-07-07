package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Aldiwildan77/rust-notifier-api/config"
	"github.com/Aldiwildan77/rust-notifier-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.Register(e)

	go func() {
		if err := e.Start(":" + fmt.Sprint(config.Cfg.Port)); err != nil {
			e.Logger.Info("Shutting down the server.")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
