package api

import (
	"database/sql"
	"log"
	"net/http"
	"standardApi/services/user"

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

// Router and server start
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	log.Println("Listen on: ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
