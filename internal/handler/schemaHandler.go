package handler

import (
	"log"
	"net/http"
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

func (h *handler) SaveDeleteSchema(w http.ResponseWriter, r *http.Request) {}
