package storage

import (
	"context"
	"database/sql"

	"marketplace-api/internal/models"
	"marketplace-api/internal/storage/sqlcdb"
)

// =====================================
// STORAGE
// =====================================

type AlmacenSQLC struct {
	q *sqlcdb.Queries
}

func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{
		q: sqlcdb.New(db),
	}
}

// =====================================
// MAPPERS
// =====================================

func aUser(u sqlcdb.Usuario) models.User {
	return models.User{
		ID:         int(u.ID),
		Name:       u.Nombre,
		Email:      u.Email,
		Password:   u.Password,
		Level:      int(u.Nivel),
		Reputation: int(u.Reputacion),
	}
}

func aCategoria(c sqlcdb.Categoria) models.Categoria {
	return models.Categoria{
		ID:   int(c.ID),
		Name: c.Nombre,
	}
}

func aProducto(p sqlcdb.Producto) models.Producto {
	return models.Producto{
		ID:          int(p.ID),
		Nombre:      p.Nombre,
		Descripcion: p.Descripcion,
		Precio:      p.Precio,
		CategoriaID: uint(p.CategoriaID),
	}
}

func aReview(r sqlcdb.Review) models.Review {
	return models.Review{
		ID:         int(r.ID),
		ReviewerID: uint(r.ReviewerID),
		ReviewedID: uint(r.ReviewedID),
		Rating:     int(r.Rating),
		Comment:    r.Comment,
	}
}

func aBadge(b sqlcdb.Badge) models.Badge {
	return models.Badge{
		ID:          int(b.ID),
		Name:        b.Nombre,
		Description: b.Descripcion,
		RequiredRep: int(b.RequiredRep),
	}
}

func aOrden(o sqlcdb.Ordene) models.Orden {
	return models.Orden{
		ID:          int(o.ID),
		ProductoID:  int(o.ProductoID),
		IDComprador: int(o.CompradorID),
		Estado:      o.Estado,
	}
}

func aMission(m sqlcdb.Misione) models.Mission {
	return models.Mission{
		ID:            int(m.ID),
		Title:         m.Title,
		Description:   m.Description,
		RequiredLevel: int(m.RequiredLevel),
		RewardPoints:  int(m.RewardPoints),
	}
}

func aUserMission(m sqlcdb.UserMission) models.UserMission {
	return models.UserMission{
		ID:        int(m.ID),
		UserID:    int(m.UserID),
		MissionID: int(m.MissionID),
		Completed: m.Completed,
	}
}

func aMensaje(m sqlcdb.Mensaje) models.Message {
	return models.Message{
		ID:         int(m.ID),
		SenderID:   int(m.SenderID),
		ReceiverID: int(m.ReceiverID),
		Content:    m.Content,
	}
}

// =====================================
// USERS
// =====================================

func (a *AlmacenSQLC) ListarUsers() []models.User {
	rows, err := a.q.ListarUsers(context.Background())
	if err != nil {
		return nil
	}

	out := make([]models.User, 0, len(rows))
	for _, r := range rows {
		out = append(out, aUser(r))
	}
	return out
}

func (a *AlmacenSQLC) BuscarUserPorID(id int) (models.User, bool) {
	u, err := a.q.BuscarUserPorID(context.Background(), int64(id))
	if err != nil {
		return models.User{}, false
	}
	return aUser(u), true
}

func (a *AlmacenSQLC) CrearUser(u models.User) models.User {
	r, err := a.q.CrearUser(context.Background(), sqlcdb.CrearUserParams{
		Nombre:     u.Name,
		Email:      u.Email,
		Password:   u.Password,
		Nivel:      int64(u.Level),
		Reputacion: int64(u.Reputation),
	})
	if err != nil {
		return models.User{}
	}
	return aUser(r)
}

func (a *AlmacenSQLC) ActualizarUser(id int, u models.User) (models.User, bool) {
	r, err := a.q.ActualizarUser(context.Background(), sqlcdb.ActualizarUserParams{
		Nombre:     u.Name,
		Email:      u.Email,
		Password:   u.Password,
		Nivel:      int64(u.Level),
		Reputacion: int64(u.Reputation),
		ID:         int64(id),
	})
	if err != nil {
		return models.User{}, false
	}
	return aUser(r), true
}

func (a *AlmacenSQLC) BorrarUser(id int) bool {
	n, err := a.q.BorrarUser(context.Background(), int64(id))
	return err == nil && n > 0
}

// =====================================
// CATEGORIAS
// =====================================
func (a *AlmacenSQLC) ListarCategorias() []models.Categoria {
	rows, err := a.q.ListarCategorias(context.Background())
	if err != nil {
		return nil
	}

	out := make([]models.Categoria, 0, len(rows))
	for _, r := range rows {
		out = append(out, aCategoria(r))
	}
	return out
}

func (a *AlmacenSQLC) BuscarCategoriaPorID(id int) (models.Categoria, bool) {
	c, err := a.q.BuscarCategoriaPorID(context.Background(), int64(id))
	if err != nil {
		return models.Categoria{}, false
	}
	return aCategoria(c), true
}

func (a *AlmacenSQLC) CrearCategoria(c models.Categoria) models.Categoria {
	r, err := a.q.CrearCategoria(context.Background(), c.Name)
	if err != nil {
		return models.Categoria{}
	}
	return aCategoria(r)
}

func (a *AlmacenSQLC) ActualizarCategoria(id int, c models.Categoria) (models.Categoria, bool) {
	r, err := a.q.ActualizarCategoria(context.Background(), sqlcdb.ActualizarCategoriaParams{
		Nombre: c.Name,
		ID:     int64(id),
	})
	if err != nil {
		return models.Categoria{}, false
	}
	return aCategoria(r), true
}

func (a *AlmacenSQLC) BorrarCategoria(id int) bool {
	n, err := a.q.BorrarCategoria(context.Background(), int64(id))
	return err == nil && n > 0
}

// =====================================
// PRODUCTOS
// =====================================

func (a *AlmacenSQLC) ListarProductos() []models.Producto {
	rows, err := a.q.ListarProductos(context.Background())
	if err != nil {
		return nil
	}

	out := make([]models.Producto, 0, len(rows))
	for _, r := range rows {
		out = append(out, aProducto(r))
	}
	return out
}

func (a *AlmacenSQLC) BuscarProductoPorID(id int) (models.Producto, bool) {
	r, err := a.q.BuscarProductoPorID(context.Background(), int64(id))
	if err != nil {
		return models.Producto{}, false
	}
	return aProducto(r), true
}

func (a *AlmacenSQLC) CrearProducto(p models.Producto) models.Producto {
	r, err := a.q.CrearProducto(context.Background(), sqlcdb.CrearProductoParams{
		Nombre:      p.Nombre,
		Descripcion: p.Descripcion,
		Precio:      p.Precio,
		CategoriaID: int64(p.CategoriaID),
	})
	if err != nil {
		return models.Producto{}
	}
	return aProducto(r)
}

func (a *AlmacenSQLC) ActualizarProducto(id int, p models.Producto) (models.Producto, bool) {
	r, err := a.q.ActualizarProducto(context.Background(), sqlcdb.ActualizarProductoParams{
		Nombre:      p.Nombre,
		Descripcion: p.Descripcion,
		Precio:      p.Precio,
		CategoriaID: int64(p.CategoriaID),
		ID:          int64(id),
	})
	if err != nil {
		return models.Producto{}, false
	}
	return aProducto(r), true
}

func (a *AlmacenSQLC) BorrarProducto(id int) bool {
	n, err := a.q.BorrarProducto(context.Background(), int64(id))
	return err == nil && n > 0
}

// =====================================
// REVIEWS
// =====================================

func (a *AlmacenSQLC) ListarReviews() []models.Review {
	rows, err := a.q.ListarReviews(context.Background())
	if err != nil {
		return nil
	}

	out := make([]models.Review, 0, len(rows))
	for _, r := range rows {
		out = append(out, aReview(r))
	}
	return out
}

func (a *AlmacenSQLC) BuscarReviewPorID(id int) (models.Review, bool) {
	r, err := a.q.BuscarReviewPorID(context.Background(), int64(id))
	if err != nil {
		return models.Review{}, false
	}
	return aReview(r), true
}

func (a *AlmacenSQLC) CrearReview(r models.Review) models.Review {
	row, err := a.q.CrearReview(context.Background(), sqlcdb.CrearReviewParams{
		ReviewerID: int64(r.ReviewerID),
		ReviewedID: int64(r.ReviewedID),
		Rating:     int64(r.Rating),
		Comment:    r.Comment,
	})
	if err != nil {
		return models.Review{}
	}
	return aReview(row)
}

func (a *AlmacenSQLC) ActualizarReview(id int, r models.Review) (models.Review, bool) {
	row, err := a.q.ActualizarReview(context.Background(), sqlcdb.ActualizarReviewParams{
		ReviewerID: int64(r.ReviewerID),
		ReviewedID: int64(r.ReviewedID),
		Rating:     int64(r.Rating),
		Comment:    r.Comment,
		ID:         int64(id),
	})
	if err != nil {
		return models.Review{}, false
	}
	return aReview(row), true
}

func (a *AlmacenSQLC) BorrarReview(id int) bool {
	n, err := a.q.BorrarReview(context.Background(), int64(id))
	return err == nil && n > 0
}

// =====================================
// BADGES
// =====================================

func (a *AlmacenSQLC) ListarBadges() []models.Badge {
	rows, err := a.q.ListarBadges(context.Background())
	if err != nil {
		return nil
	}

	out := make([]models.Badge, 0, len(rows))
	for _, r := range rows {
		out = append(out, aBadge(r))
	}
	return out
}

func (a *AlmacenSQLC) BuscarBadgePorID(id int) (models.Badge, bool) {
	r, err := a.q.BuscarBadgePorID(context.Background(), int64(id))
	if err != nil {
		return models.Badge{}, false
	}
	return aBadge(r), true
}

func (a *AlmacenSQLC) CrearBadge(b models.Badge) models.Badge {
	r, err := a.q.CrearBadge(context.Background(), sqlcdb.CrearBadgeParams{
		Nombre:      b.Name,
		Descripcion: b.Description,
		RequiredRep: int64(b.RequiredRep),
	})
	if err != nil {
		return models.Badge{}
	}
	return aBadge(r)
}

func (a *AlmacenSQLC) ActualizarBadge(id int, b models.Badge) (models.Badge, bool) {
	r, err := a.q.ActualizarBadge(context.Background(), sqlcdb.ActualizarBadgeParams{
		Nombre:      b.Name,
		Descripcion: b.Description,
		RequiredRep: int64(b.RequiredRep),
		ID:          int64(id),
	})
	if err != nil {
		return models.Badge{}, false
	}
	return aBadge(r), true
}

func (a *AlmacenSQLC) BorrarBadge(id int) bool {
	n, err := a.q.BorrarBadge(context.Background(), int64(id))
	return err == nil && n > 0
}

// =====================================
// MISSIONS
// =====================================

func (a *AlmacenSQLC) ListarMissions() []models.Mission {
	rows, err := a.q.ListarMissions(context.Background())
	if err != nil {
		return nil
	}

	out := make([]models.Mission, 0, len(rows))
	for _, r := range rows {
		out = append(out, aMission(r))
	}
	return out
}

func (a *AlmacenSQLC) BuscarMisionPorID(id int) (models.Mission, bool) {
	r, err := a.q.BuscarMisionPorID(context.Background(), int64(id))
	if err != nil {
		return models.Mission{}, false
	}
	return aMission(r), true
}

func (a *AlmacenSQLC) CrearMision(m models.Mission) models.Mission {
	r, err := a.q.CrearMision(context.Background(), sqlcdb.CrearMisionParams{
		Title:         m.Title,
		Description:   m.Description,
		RequiredLevel: int64(m.RequiredLevel),
		RewardPoints:  int64(m.RewardPoints),
	})
	if err != nil {
		return models.Mission{}
	}
	return aMission(r)
}

func (a *AlmacenSQLC) ActualizarMision(id int, m models.Mission) (models.Mission, bool) {
	r, err := a.q.ActualizarMision(context.Background(), sqlcdb.ActualizarMisionParams{
		Title:         m.Title,
		Description:   m.Description,
		RequiredLevel: int64(m.RequiredLevel),
		RewardPoints:  int64(m.RewardPoints),
		ID:            int64(id),
	})
	if err != nil {
		return models.Mission{}, false
	}
	return aMission(r), true
}

func (a *AlmacenSQLC) BorrarMision(id int) bool {
	n, err := a.q.BorrarMision(context.Background(), int64(id))
	return err == nil && n > 0
}

// =====================================
// USER MISSIONS
// =====================================

func (a *AlmacenSQLC) ListarUsermissions() []models.UserMission {
	rows, err := a.q.ListarUserMissions(context.Background())
	if err != nil {
		return nil
	}

	out := make([]models.UserMission, 0, len(rows))
	for _, r := range rows {
		out = append(out, aUserMission(r))
	}
	return out
}

func (a *AlmacenSQLC) BuscarUserMissionPorID(id int) (models.UserMission, bool) {
	r, err := a.q.BuscarUserMissionPorID(context.Background(), int64(id))
	if err != nil {
		return models.UserMission{}, false
	}
	return aUserMission(r), true
}

func (a *AlmacenSQLC) CrearUserMission(m models.UserMission) models.UserMission {
	r, err := a.q.CrearUserMission(context.Background(), sqlcdb.CrearUserMissionParams{
		UserID:    int64(m.UserID),
		MissionID: int64(m.MissionID),
		Completed: m.Completed,
	})
	if err != nil {
		return models.UserMission{}
	}
	return aUserMission(r)
}

func (a *AlmacenSQLC) ActualizarUserMission(id int, m models.UserMission) (models.UserMission, bool) {
	r, err := a.q.ActualizarUserMission(context.Background(), sqlcdb.ActualizarUserMissionParams{
		UserID:    int64(m.UserID),
		MissionID: int64(m.MissionID),
		Completed: m.Completed,
		ID:        int64(id),
	})
	if err != nil {
		return models.UserMission{}, false
	}
	return aUserMission(r), true
}

func (a *AlmacenSQLC) BorrarUserMission(id int) bool {
	n, err := a.q.BorrarUserMission(context.Background(), int64(id))
	return err == nil && n > 0
}

// =====================================
// MENSAJES
// =====================================

func (a *AlmacenSQLC) ListarMessages() []models.Message {
	rows, err := a.q.ListarMensajes(context.Background())
	if err != nil {
		return nil
	}

	out := make([]models.Message, 0, len(rows))
	for _, r := range rows {
		out = append(out, aMensaje(r))
	}
	return out
}

func (a *AlmacenSQLC) BuscarMessagePorID(id int) (models.Message, bool) {
	r, err := a.q.BuscarMensajePorID(context.Background(), int64(id))
	if err != nil {
		return models.Message{}, false
	}
	return aMensaje(r), true
}

func (a *AlmacenSQLC) CrearMessage(m models.Message) models.Message {
	r, err := a.q.CrearMensaje(context.Background(), sqlcdb.CrearMensajeParams{
		SenderID:   int64(m.SenderID),
		ReceiverID: int64(m.ReceiverID),
		Content:    m.Content,
	})
	if err != nil {
		return models.Message{}
	}
	return aMensaje(r)
}

func (a *AlmacenSQLC) ActualizarMessage(id int, m models.Message) (models.Message, bool) {
	r, err := a.q.ActualizarMensaje(context.Background(), sqlcdb.ActualizarMensajeParams{
		SenderID:   int64(m.SenderID),
		ReceiverID: int64(m.ReceiverID),
		Content:    m.Content,
		ID:         int64(id),
	})
	if err != nil {
		return models.Message{}, false
	}
	return aMensaje(r), true
}

func (a *AlmacenSQLC) BorrarMessage(id int) bool {
	n, err := a.q.BorrarMensaje(context.Background(), int64(id))
	return err == nil && n > 0
}
