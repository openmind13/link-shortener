package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Content-Type
const (
	JSON = "application/json"
	HTML = "text/html"
)

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

func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type urlMapping struct {
		LongURL  string `json:"longurl"`
		ShortURL string `json:"shorturl"`
	}

	req := new(urlMapping)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		// s.respondError(w, r, http.StatusBadRequest, nil)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := s.store.AddURL(req.LongURL, req.ShortURL); err != nil {
		s.respondError(w, r, http.StatusUnprocessableEntity, err)
	}

	type responseData struct {
		ShortURL string `json:"shorturl"`
	}

	response := responseData{
		ShortURL: "http://" + s.config.BindAddr + "/" + req.ShortURL,
	}

	s.respondJSON(w, r, http.StatusOK, response)
}

func (s *Server) handleShortURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shorturl"]
	fmt.Println(shortURL)
	// get longurl from db

	// if error - return json with "something wrong here"

	var longURL = "https://ubuntu.com/download/desktop/thank-you?version=20.04.1&architecture=amd64"
	http.Redirect(w, r, longURL, http.StatusPermanentRedirect)
}
