package main

import (
	"github.com/gin-gonic/gin"
	"metrics/config/server"
	"metrics/internal/server/controller"
	"metrics/internal/server/service"
	"metrics/internal/server/storage"
)

func main() {
	server.ParseFlagsServer()
	router := gin.Default()

	repo := storage.NewStorage()
	sr := service.NewService(&repo)
	h := controller.NewHandler(sr)

	router.POST("/update/:metric/:name/:value", h.UpdateMap)
	router.LoadHTMLGlob("internal/templates/all_metrics.html")
	router.GET("/", h.GetAllMetrics)
	router.GET("/value/:metric/:name", h.GetValueFromMetricName)
	err := router.Run(server.FlagRunAddr)
	if err != nil {
		return
	}

}
