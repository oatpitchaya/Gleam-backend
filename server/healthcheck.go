package server

import (
	"net/http"

	"github.com/Win-TS/gleam-backend.git/pkg/response"
	"github.com/labstack/echo/v4"
)

type healthCheck struct {
	App    string `json:"app"`
	Status string `json:"status"`
}

func (s *server) healthCheckService(c echo.Context) error {
	return response.SuccessResponse(c, http.StatusOK, &healthCheck{
		App:    s.cfg.App.Name,
		Status: "OK",
	})
}