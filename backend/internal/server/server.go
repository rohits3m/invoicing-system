package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rohits3m/billing-system/internal/middlewares"
)

type ServerResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type ServerConfig struct {
	Port  string
	Env   string
	DbStr string
}

type Server struct {
	Mux    *http.ServeMux
	Db     *pgxpool.Pool
	Logger *slog.Logger
	Config *ServerConfig
}

func NewServer(config ServerConfig) *Server {
	// Loading the environment variables
	godotenv.Load()

	// Default logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	// Connecting to the database
	pool, err := pgxpool.New(context.Background(), config.DbStr)
	if err != nil {
		logger.Error(err.Error())
	}

	return &Server{
		Mux:    http.NewServeMux(),
		Db:     pool,
		Logger: logger,
		Config: &config,
	}
}

func (server *Server) Run() {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:      middlewares.Cors(server.Mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(server.Logger.Handler(), slog.LevelError),
	}

	server.Logger.Info("Server listening", "port", server.Config.Port, "env", server.Config.Env)
	if err := srv.ListenAndServe(); err != nil {
		server.Logger.Error(err.Error())
	}
}

func (server *Server) SuccessResponse(w http.ResponseWriter, data any, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ServerResponse{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func (server *Server) FailureResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ServerResponse{
		Success: false,
		Message: message,
	})
}
