package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openmind13/link-shortener/app/model"
	"github.com/openmind13/link-shortener/app/utils"
)

// errors
var (
	errURLNotFound  = errors.New("Something wrong here. URL not found")
	errIncorrectURL = errors.New("Incorrect URL")
)

// GET /
func (s *Server) infoHandler(w http.ResponseWriter, r *http.Request) {
	// Show info about server
	var infoString = "Welcome to the Go link shortner API"
	fmt.Fprintf(w, infoString)
}

// POST /create
func (s *Server) handleCreateRandomURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//request := model.RequestAddRandom{}
	request := &model.RequestAddRandom{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		s.respondError(w, r, http.StatusBadRequest, err)
		return
	}

	fmt.Println(request)

	if err := utils.ValidateURL(request.LongURL); err != nil {
		s.respondError(w, r, http.StatusUnprocessableEntity, errIncorrectURL)
		return
	}

	shortURL := utils.GenerateRandomShortURL(s.config.ShortURLLength)

	// save into db and if all is Ok return to client
	if err := s.store.AddURL(request.LongURL, shortURL); err != nil {
		s.respondError(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	s.respondJSON(w, r, http.StatusCreated, model.ResponseAddRandom{
		ShortURL: "http://" + s.config.BindAddr + "/" + shortURL,
	})
}

// POST /createcustom
func (s *Server) handleCreateCustomURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(model.RequestAddCustom)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		s.respondError(w, r, http.StatusBadRequest, err)
		return
	}

	if err := utils.ValidateURL(request.LongURL); err != nil {
		s.respondError(w, r, http.StatusUnprocessableEntity, errIncorrectURL)
		return
	}

	if err := s.store.AddURL(request.LongURL, request.ShortURL); err != nil {
		s.respondError(w, r, http.StatusUnprocessableEntity, err)
	}

	s.respondJSON(w, r, http.StatusCreated, model.ResponseAddCustom{
		ShortURL: "http://" + s.config.BindAddr + "/" + request.ShortURL,
	})
}

// GET /{shorturl}
func (s *Server) handleShortURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shorturl"]

	// get longurl from db
	longURL, err := s.store.GetLongURL(shortURL)
	if err != nil {
		s.respondError(w, r, http.StatusUnprocessableEntity, errURLNotFound)
		return
	}

	// if all is OK - redirect to url
	http.Redirect(w, r, longURL, http.StatusPermanentRedirect)
}
