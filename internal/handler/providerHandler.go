package handler

import (
	"log"
	"net/http"
)

func (h *handler) AddProvider(w http.ResponseWriter, r *http.Request) {
	if err := h.service.AddProvider(r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("new provider created")
}

func (h *handler) DeleteProviderById(w http.ResponseWriter, r *http.Request) {}

func (h *handler) GetAirlinesByProviderId(w http.ResponseWriter, r *http.Request) {}
