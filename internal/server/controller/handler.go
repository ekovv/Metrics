package controller

import (
	"github.com/gorilla/mux"
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

func (l *Handler) UpdateMap(res http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	metric := vars["metric"]
	name := vars["name"]
	value := vars["value"]
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	err = l.logic.SetMetric(metric, name, val)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusOK)
}
