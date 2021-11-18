package controller

import (
	"net/http"
	"time"

	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
)

type healthController struct{}

// HealthController - Holds the abstraction of controller methods.
type HealthController interface {
	// GetServiceHealth - Calculate the uptime and return an instance of Health.
	GetServiceHealth(c Context)
}

// NewHealthController - Create and returns an instance of healthController.
func NewHealthController() *healthController {
	return &healthController{}
}

// GetServiceHealth - Calculate the uptime and return an instance of Health.
func (hc *healthController) GetServiceHealth(c Context) {
	startTime := time.Now()
	uptime := time.Since(startTime)
	health := model.Health{Uptime: uptime, StatusCode: http.StatusOK}
	c.JSON(http.StatusOK, health)
}
