package service

import (
	"marketplace-api/internal/models"
	"marketplace-api/internal/storage"
)

// Categoria//

type CategoriaService struct {
	repo storage.OrdenRepository
}

func NewCategoriaService(repo storage.OrdenRepository) *CategoriaService {
	return &CategoriaService{repo: repo}
}

func (s *CategoriaService) Listar() []models.Categoria {
	return s.repo.ListarCategorias()
}

func (s *CategoriaService) Obtener(id int) (models.Categoria, bool) {
	c, ok := s.repo.BuscarCategoriaPorID(id)
	if !ok {
		return models.Categoria{}, false
	}
	return c, true
}

func (s *CategoriaService) Crear(c models.Categoria) (models.Categoria, error) {
	if err := validarCategoria(c); err != nil {
		return models.Categoria{}, err
	}

	return s.repo.CrearCategoria(c), nil
}

func (s *CategoriaService) Actualizar(id int, c models.Categoria) (models.Categoria, error) {
	if err := validarCategoria(c); err != nil {
		return models.Categoria{}, err
	}
	c, ok := s.repo.ActualizarCategoria(id, c)
	if !ok {
		return models.Categoria{}, ErrNoEncontrado

	}
	return c, nil
}

func (s *CategoriaService) Borrar(id int) error {
	if !s.repo.BorrarCategoria(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarCategoria(c models.Categoria) error {
	if c.Name == "" {
		return ErrNombreVacio
	}

	return nil
}

// Producto//
type ProductoService struct {
	repo storage.OrdenRepository
}

func NewPorductoService(repo storage.OrdenRepository) *ProductoService {
	return &ProductoService{repo: repo}
}

func (s *ProductoService) Listar() []models.Producto {
	return s.repo.ListarProductos()
}

func (s *ProductoService) Obtener(id int) (models.Producto, bool) {
	p, ok := s.repo.BuscarProductoPorID(id)
	if !ok {
		return models.Producto{}, false
	}
	return p, true
}

func (s *ProductoService) Crear(p models.Producto) (models.Producto, error) {
	if err := validarProducto(p); err != nil {
		return models.Producto{}, err
	}

	return s.repo.CrearProducto(p), nil
}

func (s *ProductoService) Actualizar(id int, p models.Producto) (models.Producto, error) {
	if err := validarProducto(p); err != nil {
		return models.Producto{}, err
	}
	p, ok := s.repo.ActualizarProducto(id, p)
	if !ok {
		return models.Producto{}, ErrNoEncontrado

	}
	return p, nil
}

func (s *ProductoService) Borrar(id int) error {
	if !s.repo.BorrarProducto(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarProducto(p models.Producto) error {
	if p.Nombre == "" {
		return ErrNombreVacio
	}
	if p.Precio < 0 {
		return ErrPrecioNegativo
	}
	if p.CategoriaID <= 0 {
		return ErrCategoriaInvalida
	}

	return nil
}

// Orden//

type OrdenService struct {
	repo storage.OrdenRepository
}

func NewOrdenService(repo storage.OrdenRepository) *OrdenService {
	return &OrdenService{repo: repo}
}

func (s *OrdenService) Listar() []models.Orden {
	return s.repo.ListarOrdenes()
}

func (s *OrdenService) Obtener(id int) (models.Orden, bool) {
	o, ok := s.repo.BuscarOrdenPorID(id)
	if !ok {
		return models.Orden{}, false
	}
	return o, true
}

func (s *OrdenService) Crear(o models.Orden) (models.Orden, error) {
	if err := validarOrden(o); err != nil {
		return models.Orden{}, err
	}

	return s.repo.CrearOrden(o), nil
}

func (s *OrdenService) Actualizar(id int, o models.Orden) (models.Orden, error) {
	if err := validarOrden(o); err != nil {
		return models.Orden{}, err
	}
	o, ok := s.repo.ActualizarOrden(id, o)
	if !ok {
		return models.Orden{}, ErrNoEncontrado

	}
	return o, nil
}

func (s *OrdenService) Borrar(id int) error {
	if !s.repo.BorrarOrden(id) {
		return ErrOrdenNoEncontrada
	}

	return nil
}

func validarOrden(o models.Orden) error {
	if o.ProductoID <= 0 {
		return ErrProductoInvalido
	}
	if o.IDComprador <= 0 {
		return ErrCompradorInvalido
	}
	if o.Estado == "" {
		return ErrEstadoVacio
	}
	return nil
}
