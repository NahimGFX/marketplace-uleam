package service

import "errors"

// Errores de dominio. El handler los traduce a codigos HTTP:
//
//	ErrNombreVacio, ErrPrecioNegativo -> 400 Bad Request
//	ErrNoEncontrado                   -> 404 Not Found
//	ErrEmailEnUso                     -> 409 Conflict
//	ErrCredencialesInvalidas          -> 401 Unauthorized
//
// El repositorio sigue en comma-ok; es el service quien reintroduce el error
// con significado de negocio.
var (
	ErrNombreVacio           = errors.New("el campo nombre es obligatorio")
	ErrPrecioNegativo        = errors.New("el precio no puede ser negativo")
	ErrNoEncontrado          = errors.New("recurso no encontrado")
	ErrEmailEnUso            = errors.New("el email ya esta registrado")
	ErrCredencialesInvalidas = errors.New("email o contrasena incorrectos")
	ErrContentVacio          = errors.New("el campo contenido es obligatorio")
	ErrUserIDRequerido       = errors.New("el campo userID es obligatorio")
	ErrMissionIDRequerido    = errors.New("el campo missionID es obligatorio")
	ErrCategoriaInvalida     = errors.New("categoria_id debe ser mayor a cero")
	ErrProductoInvalido      = errors.New("producto_id debe ser mayor a cero")
	ErrCompradorInvalido     = errors.New("comprador_id debe ser mayor a cero")
	ErrEstadoVacio           = errors.New("el estado no puede estar vacío")
	ErrOrdenNoEncontrada     = errors.New("orden no encontrada")
)
