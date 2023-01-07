package net

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sean0427/micro-service-pratice-product-domain/model"
)

type service interface {
	Get() ([]*model.Product, error)
	GetById(context.Context, string) (*model.Product, error)
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

func (h *handler) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	product, err := h.service.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) InitHandler() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Get("/", h.Get)
		r.Get("/:id", h.GetById)
	})

	return r
}
