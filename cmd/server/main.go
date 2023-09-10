package main

import (
	"github.com/gin-gonic/gin"
	"metrics/internal/server/controller"
	"metrics/internal/server/service"
	"metrics/internal/server/storage"
)

func main() {
	router := gin.Default()

	repo := storage.NewStorage()
	sr := service.NewService(&repo)
	h := controller.NewHandler(sr)
	//router := mux.NewRouter()
	//router.HandleFunc("/update/{metric}/{name}/{value}", h.UpdateMap).Methods(http.MethodPost)
	//http.Handle("/", router)
	//log.Fatal(http.ListenAndServe("localhost:8080", router))

	router.POST("/update/:metric/:name/:value", h.UpdateMap)
	router.LoadHTMLGlob("internal/templates/all_metrics.html")

	router.GET("/", h.GetAllMetrics)
	router.GET("/value/:metric/:name", h.GetValueFromMetricName)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}

}
