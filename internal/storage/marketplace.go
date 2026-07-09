package storage

import (
	"marketplace-api/internal/models"
)

// =========================================================
// MODULO 2
// =========================================================
// CATEGORIAS

func (m *Memoria) SeedCategorias() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.categorias = []models.Categoria{
		{ID: 1, Name: "Libros"},
		{ID: 2, Name: "Equipos tecnologicos"},
		{ID: 3, Name: "Material de laboratorio"},
		{ID: 4, Name: "Uniformes"},
		{ID: 5, Name: "Accesorios universitarios"},
		{ID: 6, Name: "Otros"},
	}
	m.nextCategoriaID = 6
}
func (m *Memoria) ListarCategorias() []models.Categoria {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Categoria, len(m.categorias))
	copy(copia, m.categorias)
	return copia
}

// GET:DEVUELVE UNA CATEGORIA POR ID, SI NO SE ENCUENTRA DEVUELVE UN BOOLEANO FALSE
func (m *Memoria) BuscarCategoriaPorID(id int) (models.Categoria, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, c := range m.categorias {
		if c.ID == id {
			return c, true
		}
	}
	return models.Categoria{}, false
}

func (m *Memoria) CrearCategoria(c models.Categoria) models.Categoria {
	m.mu.Lock()
	defer m.mu.Unlock()

	c.ID = m.nextCategoriaID
	m.nextCategoriaID++
	m.categorias = append(m.categorias, c)
	return c
}
func (m *Memoria) ActualizarCategoria(id int, datos models.Categoria) (models.Categoria, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, cat := range m.categorias {
		if cat.ID == id {
			datos.ID = id
			m.categorias[i] = datos
			return datos, true
		}
	}
	return models.Categoria{}, false
}
func (m *Memoria) BorrarCategoria(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.categorias {
		if c.ID == id {
			m.categorias = append(m.categorias[:i], m.categorias[i+1:]...)
			return true
		}
	}
	return false
}

// PRODUCTOS

func (m *Memoria) SeedProductos() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.productos = []models.Producto{
		{ID: 1, Nombre: "Libro de Programación", Descripcion: "Un libro completo sobre programación en Go.", Precio: 25.99, CategoriaID: 1},
		{ID: 2, Nombre: "Libro de Física General ", Descripcion: "Libro en buen estado para carreras de ingeniería.", Precio: 20.00, CategoriaID: 2},
		{ID: 3, Nombre: "Laptop Lenovo ThinkPad P15", Descripcion: "Laptop ideal para tareas y proyectos académicos.", Precio: 680.00, CategoriaID: 3},
		{ID: 4, Nombre: "Calculadora Científica Casio", Descripcion: "Calculadora para matemáticas, física y estadística.", Precio: 20.00, CategoriaID: 4},
		{ID: 5, Nombre: "Bata de laboratorio", Descripcion: "Bata protectora para prácticas de laboratorio.", Precio: 15.00, CategoriaID: 5},
		{ID: 6, Nombre: "Kit de disección", Descripcion: "Kit completo para prácticas de biología y química.", Precio: 22.00, CategoriaID: 6},
		{ID: 7, Nombre: "Mochila universitaria", Descripcion: "Mochila resistente y espaciosa para llevar libros y laptop.", Precio: 35.00, CategoriaID: 7},
		{ID: 8, Nombre: "Estuche para laptop", Descripcion: "Estuche protector para laptop.", Precio: 12.00, CategoriaID: 8},
		{ID: 9, Nombre: "Organizador de escritorio", Descripcion: "Organizador para útiles y materiales de estudio.", Precio: 10.00, CategoriaID: 9},
	}
	m.nextProductoID = 9
}

func (m *Memoria) ListarProductos() []models.Producto {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Producto, len(m.productos))
	copy(copia, m.productos)
	return copia
}

func (m *Memoria) BuscarProductoPorID(id int) (models.Producto, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, p := range m.productos {
		if p.ID == id {
			return p, true
		}
	}
	return models.Producto{}, false
}

func (m *Memoria) CrearProducto(p models.Producto) models.Producto {
	m.mu.Lock()
	defer m.mu.Unlock()

	p.ID = m.nextProductoID
	m.nextProductoID++
	m.productos = append(m.productos, p)
	return p
}

func (m *Memoria) ActualizarProducto(id int, datos models.Producto) (models.Producto, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.productos {
		if p.ID == id {
			datos.ID = id
			m.productos[i] = datos
			return datos, true
		}
	}
	return models.Producto{}, false
}

func (m *Memoria) BorrarProducto(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.productos {
		if p.ID == id {
			m.productos = append(m.productos[:i], m.productos[i+1:]...)
			return true
		}
	}
	return false
}

//ORDENES

func (m *Memoria) SeedOrdenes() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.orden = []models.Orden{
		{ID: 1, ProductoID: 1, IDComprador: 2, Estado: "Pendiente"},
		{ID: 2, ProductoID: 3, IDComprador: 4, Estado: "Enviado"},
		{ID: 3, ProductoID: 5, IDComprador: 1, Estado: "Entregado"},
		{ID: 4, ProductoID: 9, IDComprador: 2, Estado: "Cancelado"},
	}
	m.nextOrdenID = 4
}

func (m *Memoria) ListarOrdenes() []models.Orden {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Orden, len(m.orden))
	copy(copia, m.orden)
	return copia
}

func (m *Memoria) BuscarOrdenPorID(id int) (models.Orden, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, o := range m.orden {
		if o.ID == id {
			return o, true
		}
	}
	return models.Orden{}, false
}

func (m *Memoria) CrearOrden(o models.Orden) models.Orden {
	m.mu.Lock()
	defer m.mu.Unlock()

	o.ID = m.nextOrdenID
	m.nextOrdenID++
	m.orden = append(m.orden, o)
	return o
}

func (m *Memoria) ActualizarOrden(id int, datos models.Orden) (models.Orden, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, o := range m.orden {
		if o.ID == id {
			datos.ID = id
			m.orden[i] = datos
			return datos, true
		}
	}
	return models.Orden{}, false
}

func (m *Memoria) BorrarOrden(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, o := range m.orden {
		if o.ID == id {
			m.orden = append(m.orden[:i], m.orden[i+1:]...)
			return true
		}
	}
	return false
}
