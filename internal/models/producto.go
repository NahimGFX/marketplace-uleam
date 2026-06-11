package models

type Categoria struct {
	ID   int    `json:"id"`
	Name string `json:"nombre"`
}

type Producto struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	CategoriaID int     `json:"categoria_id"`
}

type Orden struct {
	ID          int    `json:"id"`
	ProductoID  int    `json:"producto_id"`
	IDComprador int    `json:"comprador_id"`
	Estado      string `json:"estado"`
}
