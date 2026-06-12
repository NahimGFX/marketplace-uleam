package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"marketplace-api/internal/models"
	"marketplace-api/internal/storage"
)

type Server struct {
	Storage storage.Almacen
}

func NewServer(s storage.Almacen) *Server {
	return &Server{Storage: s}
}

// Users

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

	if strings.TrimSpace(datos.Name) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if strings.TrimSpace(datos.Password) == "" {
		RespondError(w, http.StatusBadRequest, "el campo contrseña es obligatorio")
		return
	}
	if strings.TrimSpace(datos.Email) == "" {
		RespondError(w, http.StatusBadRequest, "el campo email es obligatorio")
		return
	}
	if datos.Level < 0 {
		RespondError(w, http.StatusBadRequest, "el nivel no puede ser negativo")
		return
	}
	if datos.Reputation < 0 {
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

// Reviews
func (s *Server) ListarReviews(w http.ResponseWriter, r *http.Request) {
	reviews := s.Storage.ListarReviews()
	RespondJSON(w, http.StatusOK, reviews)
}

func (s *Server) ObteneReview(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	review, encontrado := s.Storage.BuscarReviewPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "review no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, review)
}

func (s *Server) CrearReview(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Review
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if nuevo.ReviewerID < 0 {
		RespondError(w, http.StatusBadRequest, "el campo reviewer_id es obligatorio")
		return
	}
	if nuevo.ReviewedID < 0 {
		RespondError(w, http.StatusBadRequest, "el campo reviewed_id es obligatorio")
		return
	}
	if nuevo.Rating < 0 || nuevo.Rating > 5 {
		RespondError(w, http.StatusBadRequest, "el rating debe ser un número entre 0 y 5")
		return
	}
	if strings.TrimSpace(nuevo.Comment) == "" {
		RespondError(w, http.StatusBadRequest, "el comentario es obligatorio")
		return
	}
	if nuevo.ReviewerID == nuevo.ReviewedID {
		RespondError(w, http.StatusBadRequest, "un usuario no puede calificarse a si mismo")
		return
	}

	creado := s.Storage.CrearReview(nuevo)
	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarReview(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos models.Review
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if datos.ReviewerID < 0 {
		RespondError(w, http.StatusBadRequest, "el campo reviewer_id es obligatorio")
		return
	}
	if datos.ReviewedID < 0 {
		RespondError(w, http.StatusBadRequest, "el campo reviewed_id es obligatorio")
		return
	}
	if datos.Rating < 0 || datos.Rating > 5 {
		RespondError(w, http.StatusBadRequest, "el rating debe ser un número entre 0 y 5")
		return
	}
	if strings.TrimSpace(datos.Comment) == "" {
		RespondError(w, http.StatusBadRequest, "el comentario es obligatorio")
		return
	}
	if datos.ReviewerID == datos.ReviewedID {
		RespondError(w, http.StatusBadRequest, "un usuario no puede calificarse a si mismo")
		return
	}

	actualizado, encontrado := s.Storage.ActualizarReview(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "Review no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

// Badges
func (s *Server) ListarBadges(w http.ResponseWriter, r *http.Request) {
	badges := s.Storage.ListarBadges()
	RespondJSON(w, http.StatusOK, badges)
}

func (s *Server) ObteneBadge(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	badge, encontrado := s.Storage.BuscarBadgePorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "badge no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, badge)
}

func (s *Server) CrearBadge(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Badge
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nuevo.Name) == "" {
		RespondError(w, http.StatusBadRequest, "el nombre es obligatorio")
		return
	}
	if strings.TrimSpace(nuevo.Description) == "" {
		RespondError(w, http.StatusBadRequest, "la descripcion es obligatorio")
		return
	}
	if nuevo.RequiredRep < 0 {
		RespondError(w, http.StatusBadRequest, "la reputacion requerida no puede ser negativa")
		return
	}

	creado := s.Storage.CrearBadge(nuevo)
	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarBadge(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos models.Badge
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(datos.Name) == "" {
		RespondError(w, http.StatusBadRequest, "el nombre es obligatorio")
		return
	}
	if strings.TrimSpace(datos.Description) == "" {
		RespondError(w, http.StatusBadRequest, "la descripcion es obligatorio")
		return
	}
	if datos.RequiredRep < 0 {
		RespondError(w, http.StatusBadRequest, "la reputacion requerida no puede ser negativa")
		return
	}

	actualizado, encontrado := s.Storage.ActualizarBadge(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "Badge no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}
