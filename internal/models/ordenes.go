package models

type Categoria struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"nombre"`
}

type Producto struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	CategoriaID uint    `json:"categoria_id"`

	Categoria Categoria `json:"-" gorm:"foreignKey:CategoriaID"`
}

type Orden struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	ProductoID  int    `json:"producto_id"`
	IDComprador int    `json:"comprador_id"`
	Estado      string `json:"estado"`

	Producto Producto ` json:"-" gorm:"foreignKey:ProductoID"`
	User     User     ` json:"-" gorm:"foreignKey:IDComprador"`
}
