package controller

import (
	"net/http"
	"time"

	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
)

type healthController struct{}

type HealthController interface {
	GetServiceHealth(c Context)
}

func NewHealthController() HealthController {
	return &healthController{}
}

func (hc *healthController) GetServiceHealth(c Context) {
	startTime := time.Now()
	uptime := time.Since(startTime)
	health := model.Health{Uptime: uptime, StatusCode: http.StatusOK}
	c.JSON(http.StatusOK, health)
}
