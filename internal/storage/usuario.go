package storage

import (
	"gorm.io/gorm"

	"marketplace-api/internal/models"
)

// UsuarioGORM implementa UserRepository sobre la tabla users de GORM.
type UsuarioGORM struct {
	db *gorm.DB
}

// NuevoUsuarioGORM envuelve una conexion *gorm.DB ya abierta.
func NuevoUsuarioGORM(db *gorm.DB) *UsuarioGORM {
	return &UsuarioGORM{db: db}
}

// CrearUsuario inserta un registro en users. Devuelve error si el insert falla
// por ejemplo, email duplicado por el indice unico.
func (r *UsuarioGORM) CrearUsuario(u models.User) (models.User, error) {
	if err := r.db.Create(&u).Error; err != nil {
		return models.User{}, err
	}
	return u, nil
}

// BuscarUsuarioPorEmail busca en users por email.
func (r *UsuarioGORM) BuscarUsuarioPorEmail(email string) (models.User, bool) {
	var u models.User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return models.User{}, false
	}
	return u, true
}

// ActualizarPasswordUsuario reemplaza el password guardado en users.
func (r *UsuarioGORM) ActualizarPasswordUsuario(id int, password string) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("password", password).Error
}

// Chequeo en tiempo de compilacion: UsuarioGORM debe cumplir UserRepository.
var _ UserRepository = (*UsuarioGORM)(nil)
