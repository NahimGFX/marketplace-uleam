package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// ListarCategorias atiende GET /api/v1/categorias.
func (s *Server) ListarCategorias(w http.ResponseWriter, _ *http.Request) {
	categorias := s.Storage.ListarCategorias()
	RespondJSON(w, http.StatusOK, categorias)
}

// ObtenerCategoria atiende GET /api/v1/categorias/{id}.
func (s *Server) ObtenerCategoria(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	categoria, encontrado := s.Storage.BuscarCategoriaPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "categoría no encontrada")
		return
	}

	RespondJSON(w, http.StatusOK, categoria)
}
