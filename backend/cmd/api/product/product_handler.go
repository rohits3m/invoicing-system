package product

import (
	"encoding/json"
	"net/http"

	"github.com/rohits3m/billing-system/internal/models"
	"github.com/rohits3m/billing-system/internal/server"
)

type ProductHandler struct {
	*server.Server
	ProductModel *models.ProductModel
}

func NewProductHandler(server *server.Server) *ProductHandler {
	return &ProductHandler{
		Server:       server,
		ProductModel: &models.ProductModel{Db: server.Db},
	}
}

func (handler *ProductHandler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := handler.ProductModel.Get()
	if err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}

	handler.SuccessResponse(w, products, "")
}

func (handler *ProductHandler) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var data models.CreateProduct
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}
	defer r.Body.Close()

	id, err := handler.ProductModel.Create(data)
	if err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}

	handler.SuccessResponse(w, id, "product created")
}
