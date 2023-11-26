package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *server) sendJSON(w http.ResponseWriter, payload any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		s.sendInternalError(w, err)
	}
}

func (s *server) sendError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func (s *server) sendInternalError(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	s.sendError(w, http.StatusInternalServerError)
}
