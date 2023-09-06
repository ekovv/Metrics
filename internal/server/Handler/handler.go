package Handler

import (
	"github.com/gorilla/mux"
	"metric/internal/server/Service"
	"net/http"
	"strconv"
)

type Handler struct {
	logic Service.Service
}

func NewHandler(logic Service.Service) *Handler {
	return &Handler{logic: logic}
}

func (l *Handler) UpdateMap(res http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	metric := vars["metric"]
	name := vars["name"]
	value := vars["value"]
	val, _ := strconv.ParseFloat(value, 64)
	err := l.logic.SetMetric(metric, name, val)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
}
