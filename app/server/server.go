package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openmind13/link-shortener/app/store"
)

// Server struct
type Server struct {
	router *mux.Router
	store  *store.Store
	config *Config
}

// New - crete and init server
func New(config *Config) (*Server, error) {
	store, err := store.New(config.MongoDBConn)
	if err != nil {
		return nil, err
	}

	s := &Server{
		router: mux.NewRouter(),
		store:  store,
		config: config,
	}

	s.configureRouter()

	// init database

	return s, nil
}

func (s *Server) configureRouter() {
	// s.router.HandleFunc("/", s.infoHandler).Methods("GET", "POST")

	s.router.HandleFunc("/create", s.handleCreate).Methods("POST")
	s.router.HandleFunc("/{shorturl}", s.handleShortURL).Methods("GET")
}

// Start - start server
func (s *Server) Start() error {
	if err := http.ListenAndServe(s.config.BindAddr, s.router); err != nil {
		return err
	}
	return nil
}
