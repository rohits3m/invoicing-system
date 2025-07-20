package user

import (
	"fmt"

	"github.com/rohits3m/billing-system/internal/server"
)

func RegisterUserRoutes(server *server.Server, base string) {
	handler := NewUserHandler(server)

	server.Mux.HandleFunc(fmt.Sprintf("GET %s/{id}", base), handler.HandleGetUserById)
	server.Mux.HandleFunc(fmt.Sprintf("POST %s", base), handler.HandleCreateUser)
}
