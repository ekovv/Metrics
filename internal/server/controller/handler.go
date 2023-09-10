package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"metrics/internal/server/service"
	"net/http"
	"strconv"
)

type Handler struct {
	logic service.Service
}

func NewHandler(logic service.Service) *Handler {
	return &Handler{logic: logic}
}

func (l *Handler) UpdateMap(c *gin.Context) {
	metric := c.Param("metric")
	name := c.Param("name")
	value := c.Param("value")
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("Invalid request"))
		return
	}

	err = l.logic.SetMetric(metric, name, val)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("Invalid request"))
		return
	}
	c.String(http.StatusOK, "OK")
}

func (l *Handler) GetAllMetrics(c *gin.Context) {
	c.HTML(http.StatusOK, "all_metrics.html", gin.H{
		"title": l.logic.GetAllMetrics(),
	})
}

func (l *Handler) GetValueFromMetricName(c *gin.Context) {
	name := c.Param("name")
	s, err := l.logic.GetValueFromM(name)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, errors.New(err.Error()))
	}
	c.HTML(http.StatusOK, "all_metrics.html", gin.H{
		"title": s,
	})
}
