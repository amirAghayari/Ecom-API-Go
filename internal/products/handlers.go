package products

import (
	"log"
	"net/http"

	"github.com/amiraghayari/Ecom-API-Go/internal/json"
)

// struct
type handler struct {
	service Service
}

// constructor
func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

// handler method
func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Println("error while listing products:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := struct {
		Products []string `json:"products"`
	}{}

	json.Write(w, http.StatusOK, products)
}
