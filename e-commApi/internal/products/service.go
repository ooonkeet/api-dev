package products

import "net/http"

type Service interface {
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// call the service -> ListProduct
	// return JSON in an HTTP response
}