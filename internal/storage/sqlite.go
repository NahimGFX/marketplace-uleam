package storage

import (
	"marketplace-api/internal/models"

	"gorm.io/gorm"
)

type AlmacenSQLite struct {
	db *gorm.DB
}

func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{
		db: db,
	}
}

//
// MESSAGES
//

func (s *AlmacenSQLite) ListarMessages() []models.Message {
	var messages []models.Message
	s.db.Find(&messages)
	return messages
}

func (s *AlmacenSQLite) BuscarMessagePorID(id int) (models.Message, bool) {
	var message models.Message

	if err := s.db.First(&message, id).Error; err != nil {
		return models.Message{}, false
	}

	return message, true
}

func (s *AlmacenSQLite) CrearMessage(message models.Message) models.Message {
	s.db.Create(&message)
	return message
}

func (s *AlmacenSQLite) ActualizarMessage(id int, datos models.Message) (models.Message, bool) {
	var message models.Message

	if err := s.db.First(&message, id).Error; err != nil {
		return models.Message{}, false
	}

	datos.ID = id
	s.db.Save(&datos)

	return datos, true
}

func (s *AlmacenSQLite) BorrarMessage(id int) bool {
	result := s.db.Delete(&models.Message{}, id)
	return result.RowsAffected > 0
}

//
// MISIONES
//

func (s *AlmacenSQLite) ListarMissions() []models.Mission {
	var missions []models.Mission
	s.db.Find(&missions)
	return missions
}

func (s *AlmacenSQLite) BuscarMisionPorID(id int) (models.Mission, bool) {
	var mission models.Mission

	if err := s.db.First(&mission, id).Error; err != nil {
		return models.Mission{}, false
	}

	return mission, true
}

func (s *AlmacenSQLite) CrearMision(mission models.Mission) models.Mission {
	s.db.Create(&mission)
	return mission
}

func (s *AlmacenSQLite) ActualizarMision(id int, datos models.Mission) (models.Mission, bool) {
	var mission models.Mission

	if err := s.db.First(&mission, id).Error; err != nil {
		return models.Mission{}, false
	}

	datos.ID = id
	s.db.Save(&datos)

	return datos, true
}

func (s *AlmacenSQLite) BorrarMision(id int) bool {
	result := s.db.Delete(&models.Mission{}, id)
	return result.RowsAffected > 0
}

//
// USER MISSIONS
//

func (s *AlmacenSQLite) ListarUsermissions() []models.UserMission {
	var userMissions []models.UserMission
	s.db.Find(&userMissions)
	return userMissions
}

func (s *AlmacenSQLite) BuscarUserMissionPorID(id int) (models.UserMission, bool) {
	var userMission models.UserMission

	if err := s.db.First(&userMission, id).Error; err != nil {
		return models.UserMission{}, false
	}

	return userMission, true
}

func (s *AlmacenSQLite) CrearUserMission(userMission models.UserMission) models.UserMission {
	s.db.Create(&userMission)
	return userMission
}

func (s *AlmacenSQLite) ActualizarUserMission(id int, datos models.UserMission) (models.UserMission, bool) {
	var userMission models.UserMission

	if err := s.db.First(&userMission, id).Error; err != nil {
		return models.UserMission{}, false
	}

	datos.ID = id
	s.db.Save(&datos)

	return datos, true
}

func (s *AlmacenSQLite) BorrarUserMission(id int) bool {
	result := s.db.Delete(&models.UserMission{}, id)
	return result.RowsAffected > 0
}
