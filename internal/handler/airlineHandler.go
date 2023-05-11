package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *handler) AddAirline(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if err := h.service.AddAirline(r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("new airline created")
}

func (h *handler) DeleteAirlineByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if errDelete := h.service.DeleteAirlineByCode(vars["iata"]); errDelete != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errDelete.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Success delete airline with iata = %s", vars["iata"])))
}

func (h *handler) UpdateProvidersList(w http.ResponseWriter, r *http.Request) {}
