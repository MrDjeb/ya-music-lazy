package endpoint

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	WhatTime() string
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	tmStr := e.s.WhatTime()

	err := ctx.String(http.StatusOK, "Status Hey~!: "+tmStr)
	if err != nil {
		return err
	}

	return nil

}
