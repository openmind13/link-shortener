package server

import "net/http"

func (s *Server) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type request struct {
	}
}
