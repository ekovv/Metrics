package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"metrics/config/server"
	"metrics/internal/server/controller"
	log "metrics/internal/server/logger"
	"metrics/internal/server/service"
	"metrics/internal/server/storage"
)

func main() {
	config := server.New()
	router := gin.Default()
	router.Use(log.HTTPLogger())
	repo := storage.NewStorage()
	sr := service.NewService(&repo)
	h := controller.NewHandler(sr)

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	log.Sugar = *logger.Sugar()

	router.POST("/update/:metric/:name/:value", h.UpdateMap)
	router.LoadHTMLGlob("internal/templates/all_metrics.html")
	router.GET("/", h.GetAllMetrics)
	router.GET("/value/:metric/:name", h.GetMetricValue)

	err = router.Run(config.Host)
	if err != nil {
		return
	}

}
