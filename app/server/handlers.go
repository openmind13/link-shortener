package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openmind13/link-shortener/app/model"
)

// Content-Type
const (
	JSON = "application/json"
	HTML = "text/html"
)

// GET
// /
func (s *Server) infoHandler(w http.ResponseWriter, r *http.Request) {
	// switch request type and display some information about server

	var infoString = "Welcome to the Go link shortner API"
	// switch r.Header.Get("Content-Type") {
	// case JSON:
	// 	s.respondJSON(w, r, http.StatusOK, infoString)
	// case HTML:
	// 	fmt.Fprintf(w, infoString)
	// default:
	// 	fmt.Println("unrecognized")
	// }

	// fmt.Println(r.Header)

	fmt.Fprintf(w, infoString)
}

// POST
// /create
func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := new(model.AddCustomRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.respondError(w, r, http.StatusBadRequest, err)
		return
	}

	fmt.Println(req)

	return

	// req := new(model.AddRequest)
	// if err := json.NewDecoder(r.Body).Decode(req); err != nil {
	// 	// s.respondError(w, r, http.StatusBadRequest, nil)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// if err := s.store.AddURL(req.LongURL, req.ShortURL); err != nil {
	// 	s.respondError(w, r, http.StatusUnprocessableEntity, err)
	// }

	// type responseData struct {
	// 	ShortURL string `json:"shorturl"`
	// }

	// response := responseData{
	// 	ShortURL: "http://" + s.config.BindAddr + "/" + req.ShortURL,
	// }

	// s.respondJSON(w, r, http.StatusOK, response)
}

// POST
// /createcustom
func (s *Server) handleCreateCustom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var shortURL string
	s.respondJSON(w, r, http.StatusCreated, shortURL)
}

// GET
// /{shorturl}
func (s *Server) handleShortURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shorturl"]
	fmt.Println(shortURL)

	// get longurl from db

	// if error - return json with "something wrong here"

	var longURL = "https://vk.com/im"
	http.Redirect(w, r, longURL, http.StatusPermanentRedirect)
}
