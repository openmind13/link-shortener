package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openmind13/link-shortener/app/config"
	"github.com/openmind13/link-shortener/app/store"
	"github.com/openmind13/link-shortener/app/store/mongodb"
)

// Server struct
type Server struct {
	router *mux.Router
	store  store.Store
	config *config.Config
}

// New - crete and init server
func New(config *config.Config) (*Server, error) {
	dbconfig := mongodb.Config{
		MongodbConnection: config.MongodbConnection,
		DBName:            config.DBName,
		CollectionName:    config.CollectionName,
	}
	store, err := mongodb.NewMongodbStore(&dbconfig)
	if err != nil {
		return nil, err
	}
	s := &Server{
		router: mux.NewRouter(),
		store:  store,
		config: config,
	}
	s.configureRouter()
	return s, nil
}

// adding handlers functions
func (s *Server) configureRouter() {
	s.router.HandleFunc("/{shorturl}", s.handleShortURL).Methods("GET")
	s.router.HandleFunc("/create", s.handleCreateRandomURL).Methods("POST")
	s.router.HandleFunc("/createcustom", s.handleCreateCustomURL).Methods("POST")
}

func (s *Server) registerMiddleware() {
	s.router.Use(s.panicMiddleware)
}

// Start - start server
func (s *Server) Start() error {
	if err := http.ListenAndServe(s.config.BindAddr, s.router); err != nil {
		return err
	}
	return nil
}

// ServeHTTP - method for testing
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
