package product

import (
	"fmt"

	"github.com/rohits3m/billing-system/internal/server"
)

func RegisterProductRoutes(server *server.Server, base string) {
	handler := NewProductHandler(server)

	server.Mux.HandleFunc(fmt.Sprintf("GET %s", base), handler.HandleGetProducts)
	server.Mux.HandleFunc(fmt.Sprintf("POST %s", base), handler.HandleCreateProduct)
}
