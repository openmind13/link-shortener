package server

import (
	"github.com/gorilla/mux"
	"github.com/openmind13/link-shortener/app/store"
)

// Server struct
type Server struct {
	router *mux.Router
	store  *store.Store
}

// New - crete and init server
func New() (*Server, error) {

	return nil, nil
}

// Start - start server
func (s *Server) Start() error {
	return nil
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/users/add", s.handle).Methods("POST")
}
