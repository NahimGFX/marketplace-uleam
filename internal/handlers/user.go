package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"marketplace-api/internal/models"
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

func (s *Server) ObtenerUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	user, encontrado := s.Storage.BuscarUserPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "user no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, user)
}

func (s *Server) CrearUser(w http.ResponseWriter, r *http.Request) {
	var nuevo models.User
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nuevo.Name) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if strings.TrimSpace(nuevo.Password) == "" {
		RespondError(w, http.StatusBadRequest, "el campo contrseña es obligatorio")
		return
	}
	if strings.TrimSpace(nuevo.Email) == "" {
		RespondError(w, http.StatusBadRequest, "el campo email es obligatorio")
		return
	}
	if nuevo.Level < 0 {
		RespondError(w, http.StatusBadRequest, "el nivel no puede ser negativo")
		return
	}
	if nuevo.Reputation < 0 {
		RespondError(w, http.StatusBadRequest, "la reputacion no puede ser negativo")
		return
	}

	creado := s.Storage.CrearUser(nuevo)
	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos models.User
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	var nuevo models.User
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nuevo.Name) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if strings.TrimSpace(nuevo.Password) == "" {
		RespondError(w, http.StatusBadRequest, "el campo contrseña es obligatorio")
		return
	}
	if strings.TrimSpace(nuevo.Email) == "" {
		RespondError(w, http.StatusBadRequest, "el campo email es obligatorio")
		return
	}
	if nuevo.Level < 0 {
		RespondError(w, http.StatusBadRequest, "el nivel no puede ser negativo")
		return
	}
	if nuevo.Reputation < 0 {
		RespondError(w, http.StatusBadRequest, "la reputacion no puede ser negativo")
		return
	}

	actualizado, encontrado := s.Storage.ActualizarUser(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "User no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if !s.Storage.BorrarUser(id) {
		RespondError(w, http.StatusNotFound, "producto no encontrado")
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
