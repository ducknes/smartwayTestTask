package handler

import (
	"log"
	"net/http"
)

func (h *handler) AddAirline(w http.ResponseWriter, r *http.Request) {
	if err := h.service.AddAirline(r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("new airline created")
}

func (h *handler) DeleteAirlineByCode(w http.ResponseWriter, r *http.Request) {}

func (h *handler) UpdateProvidersList(w http.ResponseWriter, r *http.Request) {}
