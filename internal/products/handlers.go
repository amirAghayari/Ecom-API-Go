package products

import (
	"log"
	"net/http"
	"strconv"

	"github.com/amiraghayari/Ecom-API-Go/internal/json"
	"github.com/go-chi/chi/v5"
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
	products, err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Println("error while listing products:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}

func (h *handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id <= 0 {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProductByID(r.Context(), id)
	if err != nil {
		log.Println("error while getting product by id:", err)
		http.Error(w, "failed to get product", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}
