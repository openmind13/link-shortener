package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openmind13/link-shortener/app/store"
)

// Config - server config
type Config struct {
	BindAddr       string `toml:"bind_addr"`
	ShortURLLength int    `toml:"shorturl_length"`
	// database parameters
	MongodbConnection string `toml:"mongodb_conn"`
	DBName            string `toml:"dbname"`
	CollectionName    string `toml:"collection_name"`
}

// Server struct
type Server struct {
	router *mux.Router
	store  *store.Store
	config *Config
}

// New - crete and init server
func New(config *Config) (*Server, error) {
	dbconfig := store.Config{
		MongodbConnection: config.MongodbConnection,
		DBName:            config.DBName,
		CollectionName:    config.CollectionName,
	}

	store, err := store.New(&dbconfig)
	if err != nil {
		fmt.Println("error in creating new store")
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
