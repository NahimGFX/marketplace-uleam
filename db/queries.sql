--Modulo 2
-- ==========================================
-- USERS
-- ==========================================

-- name: ListarUsers :many
SELECT id, nombre, email, password, nivel, reputacion
FROM usuarios;

-- name: BuscarUserPorID :one
SELECT id, nombre, email, password, nivel, reputacion
FROM usuarios
WHERE id = ?;

-- name: CrearUser :one
INSERT INTO usuarios (nombre, email, password, nivel, reputacion)
VALUES (?, ?, ?, ?, ?)
RETURNING id, nombre, email, password, nivel, reputacion;

-- name: ActualizarUser :one
UPDATE usuarios
SET nombre = ?,
    email = ?,
    password = ?,
    nivel = ?,
    reputacion = ?
WHERE id = ?
RETURNING id, nombre, email, password, nivel, reputacion;

-- name: BorrarUser :execrows
DELETE FROM usuarios
WHERE id = ?;


-- ==========================================
-- REVIEWS
-- ==========================================

-- name: ListarReviews :many
SELECT id, reviewer_id, reviewed_id, rating, comment
FROM reviews;

-- name: BuscarReviewPorID :one
SELECT id, reviewer_id, reviewed_id, rating, comment
FROM reviews
WHERE id = ?;

-- name: CrearReview :one
INSERT INTO reviews (reviewer_id, reviewed_id, rating, comment)
VALUES (?, ?, ?, ?)
RETURNING id, reviewer_id, reviewed_id, rating, comment;

-- name: ActualizarReview :one
UPDATE reviews
SET reviewer_id = ?,
    reviewed_id = ?,
    rating = ?,
    comment = ?
WHERE id = ?
RETURNING id, reviewer_id, reviewed_id, rating, comment;

-- name: BorrarReview :execrows
DELETE FROM reviews
WHERE id = ?;


-- ==========================================
-- BADGES
-- ==========================================

-- name: ListarBadges :many
SELECT id, nombre, descripcion, required_rep
FROM badges;

-- name: BuscarBadgePorID :one
SELECT id, nombre, descripcion, required_rep
FROM badges
WHERE id = ?;

-- name: CrearBadge :one
INSERT INTO badges (nombre, descripcion, required_rep)
VALUES (?, ?, ?)
RETURNING id, nombre, descripcion, required_rep;

-- name: ActualizarBadge :one
UPDATE badges
SET nombre = ?,
    descripcion = ?,
    required_rep = ?
WHERE id = ?
RETURNING id, nombre, descripcion, required_rep;

-- name: BorrarBadge :execrows
DELETE FROM badges
WHERE id = ?;

--Modulo 2
-- ==========================================
-- CATEGORIAS
-- ==========================================
-- name: ListarCategorias :many
SELECT id, nombre
FROM categorias;

-- name: BuscarCategoriaPorID :one
SELECT id, nombre
FROM categorias
WHERE id = ?;

-- name: CrearCategoria :one
INSERT INTO categorias (nombre)
VALUES (?)
RETURNING id, nombre;

-- name: ActualizarCategoria :one
UPDATE categorias
SET nombre = ?
WHERE id = ?
RETURNING id, nombre;

-- name: BorrarCategoria :execrows
DELETE FROM categorias
WHERE id = ?;
-- ==========================================
-- PRODUCTOS
-- ==========================================

-- name: ListarProductos :many
SELECT id, nombre, descripcion, precio, categoria_id
FROM productos;

-- name: BuscarProductoPorID :one
SELECT id, nombre, descripcion, precio, categoria_id
FROM productos
WHERE id = ?;

-- name: CrearProducto :one
INSERT INTO productos (nombre, descripcion, precio, categoria_id)
VALUES (?, ?, ?, ?)
RETURNING id, nombre, descripcion, precio, categoria_id;

-- name: ActualizarProducto :one
UPDATE productos
SET nombre = ?,
    descripcion = ?,
    precio = ?,
    categoria_id = ?
WHERE id = ?
RETURNING id, nombre, descripcion, precio, categoria_id;

-- name: BorrarProducto :execrows
DELETE FROM productos
WHERE id = ?;

-- ==========================================
-- ORDENES
-- ==========================================

-- name: ListarOrdenes :many
SELECT id, producto_id, comprador_id, estado
FROM ordenes;

-- name: BuscarOrdenPorID :one
SELECT id, producto_id, comprador_id, estado
FROM ordenes
WHERE id = ?;

-- name: CrearOrden :one
INSERT INTO ordenes (producto_id, comprador_id, estado)
VALUES (?, ?, ?)
RETURNING id, producto_id, comprador_id, estado;

-- name: ActualizarOrden :one
UPDATE ordenes
SET producto_id = ?,
    comprador_id = ?,
    estado = ?
WHERE id = ?
RETURNING id, producto_id, comprador_id, estado;

-- name: BorrarOrden :execrows
DELETE FROM ordenes
WHERE id = ?;


--Modulo 3


-- ==========================================
-- MISSIONS
-- ==========================================

-- name: ListarMissions :many
SELECT id, title, description, required_level, reward_points
FROM misiones;

-- name: BuscarMisionPorID :one
SELECT id, title, description, required_level, reward_points
FROM misiones
WHERE id = ?;

-- name: CrearMision :one
INSERT INTO misiones (title, description, required_level, reward_points)
VALUES (?, ?, ?, ?)
RETURNING id, title, description, required_level, reward_points;

-- name: ActualizarMision :one
UPDATE misiones
SET title = ?,
    description = ?,
    required_level = ?,
    reward_points = ?
WHERE id = ?
RETURNING id, title, description, required_level, reward_points;

-- name: BorrarMision :execrows
DELETE FROM misiones
WHERE id = ?;


-- ==========================================
-- USER MISSIONS
-- ==========================================

-- name: ListarUserMissions :many
SELECT id, user_id, mission_id, completed
FROM user_missions;

-- name: BuscarUserMissionPorID :one
SELECT id, user_id, mission_id, completed
FROM user_missions
WHERE id = ?;

-- name: CrearUserMission :one
INSERT INTO user_missions (user_id, mission_id, completed)
VALUES (?, ?, ?)
RETURNING id, user_id, mission_id, completed;

-- name: ActualizarUserMission :one
UPDATE user_missions
SET user_id = ?,
    mission_id = ?,
    completed = ?
WHERE id = ?
RETURNING id, user_id, mission_id, completed;

-- name: BorrarUserMission :execrows
DELETE FROM user_missions
WHERE id = ?;