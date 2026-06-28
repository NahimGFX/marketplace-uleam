package service

import (
	"strings"

	"marketplace-api/internal/models"
	"marketplace-api/internal/storage"
)

// =========================================================
// Users
// =========================================================

type UserService struct {
	repo storage.PerfilRepository
}

func NuevoUserService(repo storage.PerfilRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Listar() []models.User {
	return s.repo.ListarUsers()
}

func (s *UserService) Obtener(id int) (models.User, error) {
	u, ok := s.repo.BuscarUserPorID(id)
	if !ok {
		return models.User{}, ErrNoEncontrado
	}
	return u, nil
}

func (s *UserService) Crear(u models.User) (models.User, error) {
	if err := validarUser(u); err != nil {
		return models.User{}, err
	}
	return s.repo.CrearUser(u), nil
}

func (s *UserService) Actualizar(id int, datos models.User) (models.User, error) {
	if err := validarUser(datos); err != nil {
		return models.User{}, err
	}

	actualizado, ok := s.repo.ActualizarUser(id, datos)
	if !ok {
		return models.User{}, ErrNoEncontrado
	}

	return actualizado, nil
}

func (s *UserService) Borrar(id int) error {
	if !s.repo.BorrarUser(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarUser(u models.User) error {
	if strings.TrimSpace(u.Name) == "" {
		return ErrNombreVacio
	}

	if strings.TrimSpace(u.Email) == "" {
		return ErrEmailVacio
	}

	if strings.TrimSpace(u.Password) == "" {
		return ErrPasswordVacia
	}

	return nil
}

// =========================================================
// Reviews
// =========================================================

type ReviewService struct {
	repo storage.PerfilRepository
}

func NuevoReviewService(repo storage.PerfilRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) Listar() []models.Review {
	return s.repo.ListarReviews()
}

func (s *ReviewService) Obtener(id int) (models.Review, error) {
	r, ok := s.repo.BuscarReviewPorID(id)
	if !ok {
		return models.Review{}, ErrNoEncontrado
	}
	return r, nil
}

func (s *ReviewService) Crear(r models.Review) (models.Review, error) {
	if err := validarReview(r); err != nil {
		return models.Review{}, err
	}
	return s.repo.CrearReview(r), nil
}

func (s *ReviewService) Actualizar(id int, datos models.Review) (models.Review, error) {
	if err := validarReview(datos); err != nil {
		return models.Review{}, err
	}

	actualizado, ok := s.repo.ActualizarReview(id, datos)
	if !ok {
		return models.Review{}, ErrNoEncontrado
	}

	return actualizado, nil
}

func (s *ReviewService) Borrar(id int) error {
	if !s.repo.BorrarReview(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarReview(r models.Review) error {
	if r.ReviewerID == 0 {
		return ErrReviewerIDRequerido
	}

	if r.ReviewedID == 0 {
		return ErrReviewedIDRequerido
	}

	if strings.TrimSpace(r.Comment) == "" {
		return ErrContentVacio
	}

	return nil
}

// =========================================================
// Badges
// =========================================================

type BadgeService struct {
	repo storage.PerfilRepository
}

func NuevoBadgeService(repo storage.PerfilRepository) *BadgeService {
	return &BadgeService{repo: repo}
}

func (s *BadgeService) Listar() []models.Badge {
	return s.repo.ListarBadges()
}

func (s *BadgeService) Obtener(id int) (models.Badge, error) {
	b, ok := s.repo.BuscarBadgePorID(id)
	if !ok {
		return models.Badge{}, ErrNoEncontrado
	}
	return b, nil
}

func (s *BadgeService) Crear(b models.Badge) (models.Badge, error) {
	if err := validarBadge(b); err != nil {
		return models.Badge{}, err
	}
	return s.repo.CrearBadge(b), nil
}

func (s *BadgeService) Actualizar(id int, datos models.Badge) (models.Badge, error) {
	if err := validarBadge(datos); err != nil {
		return models.Badge{}, err
	}

	actualizado, ok := s.repo.ActualizarBadge(id, datos)
	if !ok {
		return models.Badge{}, ErrNoEncontrado
	}

	return actualizado, nil
}

func (s *BadgeService) Borrar(id int) error {
	if !s.repo.BorrarBadge(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarBadge(b models.Badge) error {
	if strings.TrimSpace(b.Name) == "" {
		return ErrNombreVacio
	}

	if strings.TrimSpace(b.Description) == "" {
		return ErrContentVacio
	}

	return nil
}
