package server

import "net/http"

func (s *Server) panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				s.respondError(w, r, http.StatusInternalServerError, nil)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
