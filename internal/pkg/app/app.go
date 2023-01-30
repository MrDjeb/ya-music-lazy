package app

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mrdjeb/ya-music-lazy/internal/app/endpoint"
	"github.com/mrdjeb/ya-music-lazy/internal/app/mw"
	"github.com/mrdjeb/ya-music-lazy/internal/app/service"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/", a.e.Status)

	return a, nil

}

func (a *App) Run() error {
	fmt.Print("Server was start!")

	if err := a.echo.Start(":8081"); err != nil {
		return err
	}
	return nil
}
