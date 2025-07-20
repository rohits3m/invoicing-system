package invoice

import (
	"encoding/json"
	"net/http"

	"github.com/rohits3m/billing-system/internal/models"
	"github.com/rohits3m/billing-system/internal/server"
)

type InvoiceHandler struct {
	*server.Server
	InvoiceModel     *models.InvoiceModel
	InvoiceItemModel *models.InvoiceItemModel
}

func NewInvoiceHandler(server *server.Server) *InvoiceHandler {
	return &InvoiceHandler{
		Server:           server,
		InvoiceModel:     &models.InvoiceModel{Db: server.Db},
		InvoiceItemModel: &models.InvoiceItemModel{Db: server.Db},
	}
}

func (handler *InvoiceHandler) HandleGetInvoices(w http.ResponseWriter, r *http.Request) {
	invoices, err := handler.InvoiceModel.Get()
	if err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}

	handler.SuccessResponse(w, invoices, "")
}

func (handler *InvoiceHandler) HandleCreateInvoice(w http.ResponseWriter, r *http.Request) {
	var data models.CreateInvoice
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}
	defer r.Body.Close()

	id, err := handler.InvoiceModel.Create(data)
	if err != nil {
		handler.FailureResponse(w, err.Error())
		return
	}

	handler.SuccessResponse(w, id, "invoice created")
}
