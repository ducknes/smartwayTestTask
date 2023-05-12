package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *handler) AddProvider(w http.ResponseWriter, r *http.Request) {
	if err := h.service.AddProvider(r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("new provider created")
}

func (h *handler) DeleteProviderById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := h.service.DeleteProviderById(vars["provider_id"]); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Success delete provider with id = %s", vars["provider_id"])))
}

func (h *handler) GetAirlinesByProviderId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := h.service.GetAirlinesByProviderId(vars["providerId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
