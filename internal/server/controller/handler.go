package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
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
		c.Status(http.StatusBadRequest)
		return
	}

	err = l.logic.SetMetric(metric, name, val)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}

func (l *Handler) GetAllMetrics(c *gin.Context) {
	c.HTML(http.StatusOK, "all_metrics.html", gin.H{
		"metrics": l.logic.GetAllMetrics(),
	})
}

func (l *Handler) GetMetricValue(c *gin.Context) {
	_ = c.Param("metric")
	name := c.Param("name")
	s, err := l.logic.GetVal(name)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Status(http.StatusOK)
	c.String(http.StatusOK, strconv.FormatFloat(s, 'f', -1, 64))
}

func (l *Handler) UpdateByJSON(c *gin.Context) {
	var metric Metrics
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &metric)
	if err != nil {
		fmt.Println("JSON NOT GOOD")
		c.Status(http.StatusBadRequest)
		return
	}
	if metric.Delta != nil {
		delta := float64(*metric.Delta)
		err = l.logic.SetMetric(metric.MType, metric.ID, delta)
		if err != nil {
			fmt.Println(err)
		}
		delta, err = l.logic.GetVal(metric.ID)
		if err != nil {
			fmt.Println(err)
			return
		}
		newDel := int64(delta)
		metric.Delta = &newDel
	} else {
		err = l.logic.SetMetric(metric.MType, metric.ID, *metric.Value)
		if err != nil {
			fmt.Println(err)
			return
		}
		val, err := l.logic.GetVal(metric.ID)
		if err != nil {
			fmt.Println(err)
			return
		}
		metric.Value = &val
	}

	bytes, err := json.MarshalIndent(metric, "", "    ")
	if err != nil {
		fmt.Println("JSON NOT GOOD")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Header("Content-Type", "application/json")
	c.Writer.Write(bytes)

}

func (l *Handler) GetMetricValueByJSON(c *gin.Context) {
	var metric Metrics
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &metric)
	if err != nil {
		fmt.Println("JSON NOT GOOD")
		c.Status(http.StatusBadRequest)
		return
	}
	val, err := l.logic.GetVal(metric.ID)
	if err != nil {
		fmt.Println(err)
	}
	if metric.MType == "counter" {
		delta := int64(val)
		metric.Delta = &delta
	}
	if metric.MType == "gauge" {
		metric.Value = &val
	}
	bytes, err := json.MarshalIndent(metric, "", "    ")
	if err != nil {
		fmt.Println("JSON NOT GOOD")
		c.Status(http.StatusNotFound)
		return
	}
	c.Header("Content-Type", "application/json")
	c.Writer.Write(bytes)

}
