package handler

import (
	"log"
	"net/http"
	"smartwayTestTAsk/internal/service"

	"github.com/gorilla/mux"
)

type Handler interface {
	Register(router *mux.Router)
}

type handler struct {
	service *service.Service
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/add_airline", h.AddAirline).Methods("POST")
	router.HandleFunc("#2", h.DeleteAirlineByCode).Methods("DELETE")
	router.HandleFunc("#3", h.AddProvider).Methods("POST")
	router.HandleFunc("#4", h.DeleteProviderById).Methods("DELETE")
	router.HandleFunc("#5", h.UpdateProvidersList).Methods("PUT")
	router.HandleFunc("#6", h.AddSchema).Methods("POST")
	router.HandleFunc("#7", h.FindSchemaByName).Methods("POST")
	router.HandleFunc("#8", h.UpdateSchema).Methods("PUT")
	router.HandleFunc("#9", h.SaveDeleteSchema).Methods("DELETE")
	router.HandleFunc("#10", h.AddAccount).Methods("POST")
	router.HandleFunc("#11", h.UptateAccountSchema).Methods("PUT")
	router.HandleFunc("#12", h.DeleteAccountById).Methods("DELETE")
	router.HandleFunc("#13", h.GetAirlinesByAccountId).Methods("GET")
	router.HandleFunc("#14", h.GetAirlinesByProviderId).Methods("GET")
}

func NewHandler(service *service.Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) AddAirline(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("add airline")); err != nil {
		log.Println(err)
	}
}

func (h *handler) DeleteAirlineByCode(w http.ResponseWriter, r *http.Request) {}

func (h *handler) AddProvider(w http.ResponseWriter, r *http.Request) {}

func (h *handler) DeleteProviderById(w http.ResponseWriter, r *http.Request) {}

func (h *handler) UpdateProvidersList(w http.ResponseWriter, r *http.Request) {}

func (h *handler) AddSchema(w http.ResponseWriter, r *http.Request) {}

func (h *handler) FindSchemaByName(w http.ResponseWriter, r *http.Request) {}

func (h *handler) UpdateSchema(w http.ResponseWriter, r *http.Request) {}

func (h *handler) SaveDeleteSchema(w http.ResponseWriter, r *http.Request) {}

func (h *handler) AddAccount(w http.ResponseWriter, r *http.Request) {}

func (h *handler) UptateAccountSchema(w http.ResponseWriter, r *http.Request) {}

func (h *handler) DeleteAccountById(w http.ResponseWriter, r *http.Request) {}

func (h *handler) GetAirlinesByAccountId(w http.ResponseWriter, r *http.Request) {}

func (h *handler) GetAirlinesByProviderId(w http.ResponseWriter, r *http.Request) {}
