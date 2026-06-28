package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"marketplace-api/internal/models"

	"github.com/go-chi/chi/v5"
)

// =====================================
// CATEGORIAS
// =====================================

// ListarCategorias atiende GET /api/v1/categorias.
func (s *Server) ListarCategorias(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Categorias.Listar())
}

// ObtenerCategoria atiende GET /api/v1/categorias/{id}.
func (s *Server) ObtenerCategoria(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	categoria, encontrado := s.Categorias.Obtener(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "categoría no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, categoria)
}

// CrearCategoria atiende POST /api/v1/categorias.
func (s *Server) CrearCategoria(w http.ResponseWriter, r *http.Request) {
	var nueva models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nueva.Name) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	creada, err := s.Categorias.Crear(nueva)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creada)
}

// ActualizarCategoria atiende PUT /api/v1/categorias/{id}.
func (s *Server) ActualizarCategoria(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var datos models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(datos.Name) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	actualizada, err := s.Categorias.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

// BorrarCategoria atiende DELETE /api/v1/categorias/{id}.
func (s *Server) BorrarCategoria(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if err := s.Categorias.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// =====================================
// PRODUCTOS
// =====================================

// ListarProductos atiende GET /api/v1/productos.
func (s *Server) ListarProductos(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Productos.Listar())
}

// ObtenerProducto atiende GET /api/v1/productos/{id}.
func (s *Server) ObtenerProducto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	producto, encontrado := s.Productos.Obtener(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "producto no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, producto)
}

// CrearProducto atiende POST /api/v1/productos.
func (s *Server) CrearProducto(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Producto
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nuevo.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if nuevo.Precio < 0 {
		RespondError(w, http.StatusBadRequest, "el precio no puede ser negativo")
		return
	}
	if nuevo.CategoriaID <= 0 {
		RespondError(w, http.StatusBadRequest, "categoriaID es obligatorio")
		return
	}
	creado, err := s.Productos.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

// ActualizarProducto atiende PUT /api/v1/productos/{id}.
func (s *Server) ActualizarProducto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var datos models.Producto
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(datos.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if strings.TrimSpace(datos.Descripcion) == "" {
		RespondError(w, http.StatusBadRequest, "el campo descripcion es obligatorio")
		return
	}
	if datos.Precio < 0 {
		RespondError(w, http.StatusBadRequest, "el precio no puede ser negativo")
		return
	}
	if datos.CategoriaID <= 0 {
		RespondError(w, http.StatusBadRequest, "categoriaID es obligatorio")
		return
	}
	actualizado, err := s.Productos.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

// BorrarProducto atiende DELETE /api/v1/productos/{id}.
func (s *Server) BorrarProducto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if err := s.Productos.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// =====================================
// ORDENES
// =====================================

// ListarOrdenes atiende GET /api/v1/ordenes.
func (s *Server) ListarOrdenes(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Ordenes.Listar())
}

// ObtenerOrden atiende GET /api/v1/ordenes/{id}.
func (s *Server) ObtenerOrden(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	orden, encontrado := s.Ordenes.Obtener(id)
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
	creada, err := s.Ordenes.Crear(nueva)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
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
	actualizada, err := s.Ordenes.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
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
	if err := s.Ordenes.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}
