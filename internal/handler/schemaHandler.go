package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *handler) AddSchema(w http.ResponseWriter, r *http.Request) {
	if err := h.service.AddSchema(r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("new schema created")
}

func (h *handler) FindSchemaByName(w http.ResponseWriter, r *http.Request) {}

func (h *handler) UpdateSchema(w http.ResponseWriter, r *http.Request) {}

func (h *handler) SaveDeleteSchema(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := h.service.SaveDeleteSchema(vars["id"]); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Success delete schema with id %s", vars["id"])))
}
