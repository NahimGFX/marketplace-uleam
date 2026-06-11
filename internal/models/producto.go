package models

type Categoria struct {
	ID   int
	Name string
}

type Producto struct {
	ID          int
	Nombre      string
	Descripcion string
	Precio      float64
	CategoriaID int
}

type Orden struct {
	ID          int
	ProductoID  int
	IDComprador int
	Estado      string
}
