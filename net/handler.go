package net

import (
	"encoding/json"
	"net/http"

	"github.com/sean0427/micro-service-pratice-product-domain/model"
)

type service interface {
	Get() ([]*model.Product, error)
}

type handler struct {
	service service
}

func New(service service) *handler {
	return &handler{
		service: service}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.service.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) InitHandler() *http.ServeMux {
	handler := http.NewServeMux()
	handler.HandleFunc("/products", h.Get)

	return handler
}
