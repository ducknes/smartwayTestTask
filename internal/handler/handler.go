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
	router.HandleFunc("/providers-service/add_airline", h.AddAirline).Methods("POST") // done

	router.HandleFunc("/providers-service/{iata}/delete_airline", h.DeleteAirlineByCode).Methods("GET", "DELETE") // done

	router.HandleFunc("/providers-service/add_provider", h.AddProvider).Methods("POST") // done

	router.HandleFunc("/providers-service/{provider_id}/delete_provider", h.DeleteProviderById).Methods("DELETE") // done

	router.HandleFunc("/providers-service/{airline_code}/update_providers_list/{provider_ids:.*}", h.UpdateProvidersList).Methods("PUT") // done

	router.HandleFunc("/providers-service/add_schema", h.AddSchema).Methods("POST") // done

	router.HandleFunc("/providers-service/find_schema/{name}", h.FindSchemaByName).Methods("GET") // done

	router.HandleFunc("/providers-service/{id}/update_schema_name", h.UpdateSchema).Methods("PUT") // done

	router.HandleFunc("/providers-service/{id}/delete_schema", h.SaveDeleteSchema).Methods("DELETE") // done

	router.HandleFunc("/providers-service/add_account", h.AddAccount).Methods("POST") // done

	router.HandleFunc("/providers-service/{account_id}/update_acc_sch/{schema_id}", h.UptateAccountSchema).Methods("PUT") // done

	router.HandleFunc("/providers-service/{id}/delete_account", h.DeleteAccountById).Methods("DELETE") // done

	router.HandleFunc("/providers-service/{accountId}/account-airlines", h.GetAirlinesByAccountId).Methods("GET") 

	router.HandleFunc("/providers-service/{providerId}/provider-airlines", h.GetAirlinesByProviderId).Methods("GET")
}

func NewHandler(service service.Service) Handler {
	return &handler{
		service: service,
	}
}
