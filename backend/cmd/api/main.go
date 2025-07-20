package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rohits3m/billing-system/cmd/api/invoice"
	"github.com/rohits3m/billing-system/cmd/api/product"
	"github.com/rohits3m/billing-system/cmd/api/user"
	"github.com/rohits3m/billing-system/internal/server"
)

func main() {
	// Loading the environment variables
	godotenv.Load()
	config := server.ServerConfig{
		Port:  os.Getenv("PORT"),
		Env:   os.Getenv("ENV"),
		DbStr: os.Getenv("DB_URL"),
	}

	// Creating the server
	server := server.NewServer(config)

	// Registering routes
	user.RegisterUserRoutes(server, "/api/v1/user")
	product.RegisterProductRoutes(server, "/api/v1/product")
	invoice.RegisterInvoiceRoutes(server, "/api/v1/invoice")

	// Starting the server
	server.Run()
}
