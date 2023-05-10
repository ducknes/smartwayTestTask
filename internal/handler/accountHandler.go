package handler

import (
	"log"
	"net/http"
)

func (h *handler) AddAccount(w http.ResponseWriter, r *http.Request) {
	if err := h.service.AddAccount(r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("new account created")
}

func (h *handler) UptateAccountSchema(w http.ResponseWriter, r *http.Request) {}

func (h *handler) DeleteAccountById(w http.ResponseWriter, r *http.Request) {}

func (h *handler) GetAirlinesByAccountId(w http.ResponseWriter, r *http.Request) {}
