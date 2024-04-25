package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"simple-api/service/user"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Start() error {
	// Start the server
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Register routes
	userHandler := user.NewHandler()

	userHandler.RegisterRoutes(subrouter)

	fmt.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
