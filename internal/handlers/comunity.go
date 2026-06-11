package handlers

import (
	"net/http"
)

func (s *Server) ListarMessages(w http.ResponseWriter, _ *http.Request) {
	messages := s.Storage.ListarMessages()
	RespondJSON(w, http.StatusOK, messages)
}
