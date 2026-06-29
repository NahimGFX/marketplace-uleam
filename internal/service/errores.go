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
	ErrEmailVacio            = errors.New("el email es obligatorio")
	ErrPasswordVacia         = errors.New("la contraseña es obligatoria")
	ErrReviewerIDRequerido   = errors.New("el reviewer_id es obligatorio")
	ErrReviewedIDRequerido   = errors.New("el reviewed_id es obligatorio")
	ErrCategoriaInvalida     = errors.New("la categoria es invalida")
	ErrOrdenNoEncontrada     = errors.New("la orden no fue encontrada")
	ErrProductoInvalido      = errors.New("el producto es invalido")
	ErrEstadoVacio           = errors.New("el estado es obligatorio")
	ErrCompradorInvalido     = errors.New("el comprador es invalido")
)
