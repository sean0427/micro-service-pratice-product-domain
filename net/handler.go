package net

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean0427/micro-service-pratice-product-domain/model"
)

type service interface {
	Get(context.Context, *model.GetProductsParams) ([]*model.Product, error)
	GetByID(context.Context, string) (*model.Product, error)
	Create(ctx context.Context, product *model.Product) (string, error)
	Update(ctx context.Context, id string, params *model.Product) (*model.Product, error)
	Delete(ctx context.Context, id string) error
}

type handler struct {
	service service
}

func New(service service) *handler {
	return &handler{
		service: service}
}

func getProductsParams(c *gin.Context) (*model.GetProductsParams, error) {
	params := model.GetProductsParams{}

	name := c.Params.ByName("name")
	if name != "" {
		params.Name = &name
	}

	manu := c.Params.ByName("manufacturer")
	if manu != "" {
		params.ManufacturerID = &manu
	}

	return &params, nil
}

func (h *handler) Get(c *gin.Context) {

	params, err := getProductsParams(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	products, err := h.service.Get(c, params)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSONP(http.StatusOK, products)
}

func (h *handler) GetById(c *gin.Context) {
	id, found := c.Params.Get("id")
	if !found {
		// should never happen
		c.AbortWithStatus(http.StatusBadRequest)
	}
	product, err := h.service.GetByID(c, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSONP(http.StatusOK, product)
}

func (h *handler) Create(c *gin.Context) {
	var params *model.Product

	if err := c.BindJSON(&params); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	id, err := h.service.Create(c, params)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *handler) Update(c *gin.Context) {
	var params *model.Product

	id, found := c.Params.Get("id")
	if found {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if params.ID.IsZero() {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Update body should not with id"))
	}

	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.Update(c, id, params); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Status(http.StatusCreated)
}

func (h *handler) Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := h.service.Delete(c, id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Status(http.StatusNoContent)
}

func (h *handler) Register(r *gin.Engine) {
	p := r.Group("/products")
	{
		p.GET("/", h.Get)
		p.GET("/:id")
		r.GET("/:id", h.GetById)
		r.POST("/", h.Create)
		r.PUT("/:id", h.Update)
		r.DELETE("/:id", h.Delete)
	}
}
