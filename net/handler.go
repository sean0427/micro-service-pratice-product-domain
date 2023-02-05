package net

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sean0427/micro-service-pratice-product-domain/model"
)

type service interface {
	Get(context.Context, *model.GetProductsParams) ([]*model.Product, error)
	GetByID(context.Context, string) (*model.Product, error)
	Create(ctx context.Context, product *model.Product) (string, error)
	Update(ctx context.Context, id string, params *model.Product) (*model.Product, error)
}

type handler struct {
	service service
}

func New(service service) *handler {
	return &handler{
		service: service}
}

func getProductsParams(r *http.Request) (*model.GetProductsParams, error) {
	params := model.GetProductsParams{}

	name := chi.URLParam(r, "name")
	if name != "" {
		params.Name = &name
	}

	manu := chi.URLParam(r, "manufacturer")
	if manu != "" {
		params.ManufacturerID = &manu
	}

	return &params, nil
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params, err := getProductsParams(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	products, err := h.service.Get(r.Context(), params)
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
	product, err := h.service.GetByID(r.Context(), id)
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

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var params *model.Product

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.service.Create(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(id))
	w.WriteHeader(http.StatusCreated)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	var params *model.Product

	id := chi.URLParam(r, "id")
	if id != "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if params.ID.IsZero() {
		http.Error(w, "Update body should not with id", http.StatusBadRequest)
		return
	}

	if _, err := h.service.Update(r.Context(), id, params); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) Handler() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Get("/", h.Get)
		r.Get("/:id", h.GetById)
		r.Post("/", h.Create)
		r.Put("/:id", h.Update)
	})

	return r
}
