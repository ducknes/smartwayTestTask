package handler

import (
	"smartwayTestTAsk/internal/service"

	"github.com/gorilla/mux"
)

type Handler interface {
	Register(router *mux.Router)
}

type handler struct {
	service service.Service
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/providers-service/add_airline", h.AddAirline).Methods("POST")
	router.HandleFunc("/providers-service/delete_airline", h.DeleteAirlineByCode).Methods("DELETE")
	router.HandleFunc("/providers-service/add_provider", h.AddProvider).Methods("POST")
	router.HandleFunc("/providers-service/delete_provider", h.DeleteProviderById).Methods("DELETE")
	router.HandleFunc("/providers-service/update_providers_list", h.UpdateProvidersList).Methods("PUT")
	router.HandleFunc("/providers-service/add_schema", h.AddSchema).Methods("POST")
	router.HandleFunc("/providers-service/find_schema", h.FindSchemaByName).Methods("POST")
	router.HandleFunc("/providers-service/update_schema", h.UpdateSchema).Methods("PUT")
	router.HandleFunc("/providers-service/delete_schema", h.SaveDeleteSchema).Methods("DELETE")
	router.HandleFunc("/providers-service/add_account", h.AddAccount).Methods("POST")
	router.HandleFunc("/providers-service/update_acc_sch", h.UptateAccountSchema).Methods("PUT")
	router.HandleFunc("/providers-service/delete_account", h.DeleteAccountById).Methods("DELETE")
	router.HandleFunc("/providers-service/:accountId/account-airlines", h.GetAirlinesByAccountId).Methods("GET")
	router.HandleFunc("/providers-service/:providerId/provider-airlines", h.GetAirlinesByProviderId).Methods("GET")
}

func NewHandler(service service.Service) Handler {
	return &handler{
		service: service,
	}
}
