package server

import (
	"encoding/json"
	"errors"
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

// POST /create
func (s *Server) handleCreateRandomURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//request := model.RequestAddRandom{}
	request := &model.RequestAddRandom{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		s.respondError(w, r, http.StatusBadRequest, err)
		return
	}
	if err := utils.ValidateURL(request.LongURL); err != nil {
		s.respondError(w, r, http.StatusUnprocessableEntity, err)
		return
	}
	shortURL := utils.GenerateRandomShortURL(s.config.ShortURLLength)
	// save into db and if all is Ok return to client
	data := model.Data{
		LongURL:  request.LongURL,
		ShortURL: shortURL,
	}
	if err := s.store.Add(data); err != nil {
		s.respondError(w, r, http.StatusUnprocessableEntity, err)
		return
	}
	s.respondJSON(w, r, http.StatusCreated, model.ResponseAddRandom{
		ShortURL: "http://localhost:8080/" + shortURL,
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
	data := model.Data{
		LongURL:  request.LongURL,
		ShortURL: request.ShortURL,
	}
	if err := s.store.Add(data); err != nil {
		s.respondError(w, r, http.StatusUnprocessableEntity, err)
	}
	s.respondJSON(w, r, http.StatusCreated, model.ResponseAddCustom{
		ShortURL: "http://localhost:8080/" + request.ShortURL,
	})
}

// GET /{shorturl}
func (s *Server) handleShortURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shorturl"]
	// if err := utils.ValidateURL(shortURL); err != nil {
	// 	s.respondError(w, r, http.StatusNotFound, err)
	// 	return
	// }
	// get longurl from db
	data := model.Data{
		ShortURL: shortURL,
	}
	answer, err := s.store.Get(data)
	if err != nil {
		s.respondError(w, r, http.StatusUnprocessableEntity, errURLNotFound)
		return
	}
	// if all is OK - redirect to url
	http.Redirect(w, r, answer.LongURL, http.StatusPermanentRedirect)
}
