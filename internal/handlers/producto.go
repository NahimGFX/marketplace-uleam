package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// ListarProductos atiende GET /api/v1/productos.
func (s *Server) ListarProductos(w http.ResponseWriter, _ *http.Request) {
	productos := s.Storage.ListarProductos()
	RespondJSON(w, http.StatusOK, productos)
}

// ObtenerProducto atiende GET /api/v1/productos/{id}.
func (s *Server) ObtenerProducto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	producto, encontrado := s.Storage.BuscarProductoPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "producto no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, producto)
}
