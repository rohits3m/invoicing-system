package invoice

import (
	"fmt"

	"github.com/rohits3m/billing-system/internal/server"
)

func RegisterInvoiceRoutes(server *server.Server, base string) {
	handler := NewInvoiceHandler(server)

	server.Mux.HandleFunc(fmt.Sprintf("GET %s", base), handler.HandleGetInvoices)
	server.Mux.HandleFunc(fmt.Sprintf("POST %s", base), handler.HandleCreateInvoice)
}
