package handlers

import (
	"encoding/json"
	"marketplace-api/internal/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

// /Message handlers
func (s *Server) ListarMessages(w http.ResponseWriter, _ *http.Request) {
	messages := s.Storage.ListarMessages()
	RespondJSON(w, http.StatusOK, messages)
}

// ObtenerMessage atiende GET /api/v1/mensajes/{id}.
func (s *Server) ObtenerMessage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	message, encontrado := s.Storage.BuscarMessagePorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "mensaje no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, message)
}

// CrearMessage atiende POST /api/v1/mensajes.
func (s *Server) CrearMessage(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Message

	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	if nuevo.SenderID <= 0 {
		RespondError(w, http.StatusBadRequest, "sender_id es obligatorio")
		return
	}

	if nuevo.ReceiverID <= 0 {
		RespondError(w, http.StatusBadRequest, "receiver_id es obligatorio")
		return
	}

	if strings.TrimSpace(nuevo.Content) == "" {
		RespondError(w, http.StatusBadRequest, "el contenido del mensaje es obligatorio")
		return
	}

	creado := s.Storage.CrearMessage(nuevo)
	RespondJSON(w, http.StatusCreated, creado)
}

// ActualizarMessage atiende PUT /api/v1/mensajes/{id}.
func (s *Server) ActualizarMessage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos models.Message
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if datos.SenderID <= 0 {
		RespondError(w, http.StatusBadRequest, "sender_id es obligatorio")
		return
	}

	if datos.ReceiverID <= 0 {
		RespondError(w, http.StatusBadRequest, "receiver_id es obligatorio")
		return
	}

	if strings.TrimSpace(datos.Content) == "" {
		RespondError(w, http.StatusBadRequest, "el contenido es obligatorio")
		return
	}

	actualizado, encontrado := s.Storage.ActualizarMessage(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "mensaje no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

// BorrarMessage atiende DELETE /api/v1/mensajes/{id}.
func (s *Server) BorrarMessage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if !s.Storage.BorrarMessage(id) {
		RespondError(w, http.StatusNotFound, "mensaje no encontrado")
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}

// /Misiones handlers
func (s *Server) ListarMissions(w http.ResponseWriter, _ *http.Request) {
	missions := s.Storage.ListarMissions()
	RespondJSON(w, http.StatusOK, missions)
}
