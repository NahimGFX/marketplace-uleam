package service

import (
	"strings"

	"marketplace-api/internal/models"
	"marketplace-api/internal/storage"
)

// MessageService contiene la logica de negocio de mensajes.
// Depende solo de storage.ComunidadRepository.
// =========================================================
// Messages
// =========================================================

type MessageService struct {
	repo storage.ComunidadRepository
}

func NuevoMessageService(repo storage.ComunidadRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) Listar() []models.Message {
	return s.repo.ListarMessages()
}

func (s *MessageService) Obtener(id int) (models.Message, error) {
	m, ok := s.repo.BuscarMessagePorID(id)
	if !ok {
		return models.Message{}, ErrNoEncontrado
	}
	return m, nil
}

func (s *MessageService) Crear(m models.Message) (models.Message, error) {
	if err := validarMessage(m); err != nil {
		return models.Message{}, err
	}
	return s.repo.CrearMessage(m), nil
}

func (s *MessageService) Actualizar(id int, datos models.Message) (models.Message, error) {
	if err := validarMessage(datos); err != nil {
		return models.Message{}, err
	}
	actualizado, ok := s.repo.ActualizarMessage(id, datos)
	if !ok {
		return models.Message{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *MessageService) Borrar(id int) error {
	if !s.repo.BorrarMessage(id) {
		return ErrNoEncontrado
	}
	return nil
}

// validarMessage centraliza las reglas de negocio que antes vivian en el handler.
func validarMessage(p models.Message) error {
	if strings.TrimSpace(p.Content) == "" {
		return ErrContentVacio
	}
	return nil
}

// =========================================================
// Missions
// =========================================================
type MissionService struct {
	repo storage.ComunidadRepository
}

func NuevoMissionService(repo storage.ComunidadRepository) *MissionService {
	return &MissionService{repo: repo}
}

func (s *MissionService) Listar() []models.Mission {
	return s.repo.ListarMissions()
}

func (s *MissionService) Obtener(id int) (models.Mission, error) {
	mi, ok := s.repo.BuscarMisionPorID(id)
	if !ok {
		return models.Mission{}, ErrNoEncontrado
	}
	return mi, nil
}

func (s *MissionService) Crear(mi models.Mission) (models.Mission, error) {
	if err := validarMision(mi); err != nil {
		return models.Mission{}, err
	}
	return s.repo.CrearMision(mi), nil
}

func (s *MissionService) Actualizar(id int, datos models.Mission) (models.Mission, error) {
	if err := validarMision(datos); err != nil {
		return models.Mission{}, err
	}
	actualizado, ok := s.repo.ActualizarMision(id, datos)
	if !ok {
		return models.Mission{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *MissionService) Borrar(id int) error {
	if !s.repo.BorrarMision(id) {
		return ErrNoEncontrado
	}
	return nil
}

// validarMision centraliza las reglas de negocio que antes vivian en el handler.
func validarMision(p models.Mission) error {
	if strings.TrimSpace(p.Description) == "" {
		return ErrContentVacio
	}

	return nil
}

// =========================================================
// userMissions
// =========================================================

type UserMissionService struct {
	repo storage.ComunidadRepository
}

func NuevoUserMissionService(repo storage.ComunidadRepository) *UserMissionService {
	return &UserMissionService{repo: repo}
}

func (s *UserMissionService) Listar() []models.UserMission {
	return s.repo.ListarUserMissions()
}

func (s *UserMissionService) Obtener(id int) (models.UserMission, error) {
	um, ok := s.repo.BuscarUserMissionPorID(id)
	if !ok {
		return models.UserMission{}, ErrNoEncontrado
	}
	return um, nil
}

func (s *UserMissionService) Crear(um models.UserMission) (models.UserMission, error) {
	if err := validarUserMission(um); err != nil {
		return models.UserMission{}, err
	}
	return s.repo.CrearUserMission(um), nil
}

func (s *UserMissionService) Actualizar(id int, datos models.UserMission) (models.UserMission, error) {
	if err := validarUserMission(datos); err != nil {
		return models.UserMission{}, err
	}
	actualizado, ok := s.repo.ActualizarUserMission(id, datos)
	if !ok {
		return models.UserMission{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *UserMissionService) Borrar(id int) error {
	if !s.repo.BorrarUserMission(id) {
		return ErrNoEncontrado
	}
	return nil
}

// validarUserMission centraliza las reglas de negocio que antes vivian en el handler.
func validarUserMission(um models.UserMission) error {
	if um.UserID == 0 {
		return ErrUserIDRequerido
	}

	if um.MissionID == 0 {
		return ErrMissionIDRequerido
	}
	return nil
}
