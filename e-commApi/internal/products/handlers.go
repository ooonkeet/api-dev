package products

import (
	// "encoding/json"
	// "encoding/json"
	"net/http"

	"github.com/ooonkeet/api-dev/e-commApi/internal/json"
)

type handler struct{
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// call the service -> ListProduct
	// return JSON in an HTTP response
	products := struct{
		Products []string `json:"products"`
	}{}
	json.Write(w,http.StatusOK,products)
}