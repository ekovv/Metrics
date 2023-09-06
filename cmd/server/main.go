package main

import (
	"github.com/gorilla/mux"
	"log"
	"metrics/internal/server/Handler"
	"net/http"
)

func main() {
	h := Handler.Handler{}
	router := mux.NewRouter()
	router.HandleFunc("/update/{metric}/{name}/{value}", h.UpdateMap).Methods(http.MethodPost)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
