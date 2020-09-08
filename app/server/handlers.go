package server

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	request := new(urlMapping)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		// s.respondError(w, r, http.StatusBadRequest, nil)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
