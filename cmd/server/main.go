package main

import (
	"github.com/gorilla/mux"
	"log"
	"metrics/internal/server/my_handler"
	"metrics/internal/server/my_storage"
	"metrics/internal/server/service"
	"net/http"
)

func main() {
	repo := my_storage.NewStorage()
	sr := service.NewService(repo)
	h := my_handler.NewHandler(sr)
	router := mux.NewRouter()
	router.HandleFunc("/update/{metric}/{name}/{value}", h.UpdateMap).Methods(http.MethodPost)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
