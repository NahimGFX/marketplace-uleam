package service

import "errors"

var (
	ErrNoEncontrado      = errors.New("no encontrado")
	ErrNombreVacio       = errors.New("el nombre no puede estar vacío")
	ErrPrecioNegativo    = errors.New("el precio no puede ser negativo")
	ErrCategoriaInvalida = errors.New("categoria_id debe ser mayor a cero")
	ErrProductoInvalido  = errors.New("producto_id debe ser mayor a cero")
	ErrCompradorInvalido = errors.New("comprador_id debe ser mayor a cero")
	ErrEstadoVacio       = errors.New("el estado no puede estar vacío")
	ErrOrdenNoEncontrada = errors.New("orden no encontrada")
)
