package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"marketplace-api/internal/models"

	"github.com/go-chi/chi/v5"
)

// ListarOrdenes atiende GET /api/v1/ordenes.
func (s *Server) ListarOrdenes(w http.ResponseWriter, _ *http.Request) {
	ordenes := s.Storage.ListarOrdenes()
	RespondJSON(w, http.StatusOK, ordenes)
}

// ObtenerOrden atiende GET /api/v1/ordenes/{id}.
func (s *Server) ObtenerOrden(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	orden, encontrado := s.Storage.BuscarOrdenPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "orden no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, orden)
}

// CrearOrden atiende POST /api/v1/ordenes.
func (s *Server) CrearOrden(w http.ResponseWriter, r *http.Request) {
	var nueva models.Orden
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if nueva.ProductoID <= 0 {
		RespondError(w, http.StatusBadRequest, "productoID es obligatorio")
		return
	}
	if nueva.IDComprador <= 0 {
		RespondError(w, http.StatusBadRequest, "idComprador es obligatorio")
		return
	}
	if strings.TrimSpace(nueva.Estado) == "" {
		RespondError(w, http.StatusBadRequest, "el campo estado es obligatorio")
		return
	}
	creada := s.Storage.CrearOrden(nueva)
	RespondJSON(w, http.StatusCreated, creada)
}

// ActualizarOrden atiende PUT /api/v1/ordenes/{id}.
func (s *Server) ActualizarOrden(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos models.Orden
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if datos.ProductoID <= 0 {
		RespondError(w, http.StatusBadRequest, "productoID es obligatorio")
		return
	}
	if datos.IDComprador <= 0 {
		RespondError(w, http.StatusBadRequest, "idComprador es obligatorio")
		return
	}
	if datos.Estado == "" {
		RespondError(w, http.StatusBadRequest, "el campo estado es obligatorio")
		return
	}

	actualizada, encontrada := s.Storage.ActualizarOrden(id, datos)
	if !encontrada {
		RespondError(w, http.StatusNotFound, "orden no encontrada")
		return
	}

	RespondJSON(w, http.StatusOK, actualizada)

}

// BorrarOrden atiende DELETE /api/v1/ordenes/{id}.
func (s *Server) BorrarOrden(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if !s.Storage.BorrarOrden(id) {
		RespondError(w, http.StatusNotFound, "orden no encontrada")
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
