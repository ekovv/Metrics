package main

import (
	"github.com/gorilla/mux"
	"log"
	"metrics/internal/server/Handler"
	"metrics/internal/server/Service"
	"metrics/internal/server/Storage"
	"net/http"
)

func main() {
	repo := Storage.NewStorage()
	sr := Service.NewService(repo)
	h := Handler.NewHandler(sr)
	router := mux.NewRouter()
	router.HandleFunc("/update/{metric}/{name}/{value}", h.UpdateMap).Methods(http.MethodPost)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
