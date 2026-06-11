package handlers

import (
	"net/http"
	"strconv"

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
