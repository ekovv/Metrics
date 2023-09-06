package main

import (
	"github.com/gorilla/mux"
	"log"
	"metrics/internal/server/controller"
	"metrics/internal/server/service"
	"metrics/internal/server/storage"
	"net/http"
)

func main() {
	repo := storage.NewStorage()
	sr := service.NewService(repo)
	h := controller.NewHandler(sr)
	router := mux.NewRouter()
	router.HandleFunc("/update/{metric}/{name}/{value}", h.UpdateMap).Methods(http.MethodPost)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
