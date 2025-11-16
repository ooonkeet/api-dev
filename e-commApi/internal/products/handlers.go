package products

import (
	// "encoding/json"
	// "encoding/json"
	"log"
	"net/http"

	"github.com/ooonkeet/api-dev/e-commApi/internal/json"
	// "github.com/ooonkeet/api-dev/e-commApi/internal/products"
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
	err:=h.service.ListProducts(r.Context())
	if err!=nil{
		log.Println(err)
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	products := struct{
		Products []string `json:"products"`
	}{}
	json.Write(w,http.StatusOK,products)
}