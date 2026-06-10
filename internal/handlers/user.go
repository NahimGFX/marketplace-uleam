package handlers

import (
	"marketplace-api/internal/storage"
	"net/http"
)

type Server struct {
	Storage storage.Almacen
}

func NewServer(s storage.Almacen) *Server {
	return &Server{Storage: s}
}

func (s *Server) ListarUsers(w http.ResponseWriter, _ *http.Request) {
	users := s.Storage.ListarUsers()
	RespondJSON(w, http.StatusOK, users)
}
