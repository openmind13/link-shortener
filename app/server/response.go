package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) respondJSON(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *Server) respondError(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respondJSON(w, r, code, map[string]string{"error": err.Error()})
}
