# Marketplace ULEAM

API REST para un marketplace academico de la Universidad Laica Eloy Alfaro de Manabi. Permite registrar usuarios, publicar productos academicos, gestionar ordenes, reseñas, insignias, mensajes y misiones de participacion.

## Integrantes y modulos

| Modulo | Responsable | Entidades principales | Capas |
| --- | --- | --- | --- |
| Perfil y reputacion | Jostin Alvarado | users, reviews, badges | handler -> service -> repository |
| Marketplace | Daivelyn Pincay | categorias, productos, ordenes | handler -> service -> repository |
| Comunidad | Patricio Simba | messages, missions, user_missions | handler -> service -> repository |

## Stack

- Go 1.26
- Chi Router
- GORM
- PostgreSQL en Docker
- SQLite para ejecucion local rapida y pruebas
- JWT con roles `admin` y `estudiante`
- GitHub Actions para build, vet y test
- Dockerfile multi-stage y docker-compose

## Como correr con Docker

1. Copiar variables de ejemplo si hace falta:

```bash
cp .env.example .env
```

2. Levantar API + PostgreSQL:

```bash
docker-compose up --build
```

La API queda disponible en `http://localhost:8080/api/v1`. Al iniciar, GORM ejecuta `AutoMigrate` y el repositorio siembra datos iniciales si la base esta vacia.

Credenciales de demo:

| Rol | Email | Password |
| --- | --- | --- |
| admin | juan@uleam.edu.ec | 123456 |
| estudiante | maria@uleam.edu.ec | 123456 |

## Autenticacion

Primero inicia sesion y copia el token JWT:

```http
POST /api/v1/auth/login
Content-Type: application/json

{"email":"juan@uleam.edu.ec","password":"123456"}
```

Usa el token en rutas protegidas:

```http
Authorization: Bearer <token>
```

Las rutas de lectura y flujo normal aceptan usuarios autenticados. Las rutas administrativas de categorias, badges y missions requieren rol `admin`.

## Endpoints por modulo

### Auth

| Metodo | Ruta | Descripcion |
| --- | --- | --- |
| POST | `/api/v1/auth/register` | Registra usuario estudiante |
| POST | `/api/v1/auth/login` | Devuelve JWT |

### Perfil y reputacion

| Metodo | Ruta | Descripcion | Rol |
| --- | --- | --- | --- |
| GET | `/api/v1/users` | Lista usuarios | autenticado |
| GET | `/api/v1/users/{id}` | Obtiene usuario | autenticado |
| POST | `/api/v1/users` | Crea usuario desde modulo perfil | autenticado |
| PUT | `/api/v1/users/{id}` | Actualiza usuario | autenticado |
| DELETE | `/api/v1/users/{id}` | Elimina usuario | autenticado |
| GET | `/api/v1/reviews` | Lista reseñas | autenticado |
| POST | `/api/v1/reviews` | Crea reseña | autenticado |
| GET | `/api/v1/badges` | Lista insignias | autenticado |
| POST/PUT/DELETE | `/api/v1/badges` | Gestiona insignias | admin |

### Marketplace

| Metodo | Ruta | Descripcion | Rol |
| --- | --- | --- | --- |
| GET | `/api/v1/categorias` | Lista categorias | autenticado |
| POST/PUT/DELETE | `/api/v1/categorias` | Gestiona categorias | admin |
| GET | `/api/v1/productos` | Lista productos | autenticado |
| POST | `/api/v1/productos` | Publica producto | autenticado |
| GET | `/api/v1/ordenes` | Lista ordenes | autenticado |
| POST | `/api/v1/ordenes` | Crea orden asociando producto y comprador | autenticado |

### Comunidad

| Metodo | Ruta | Descripcion | Rol |
| --- | --- | --- | --- |
| GET | `/api/v1/messages` | Lista mensajes | autenticado |
| POST | `/api/v1/messages` | Crea mensaje entre usuarios | autenticado |
| GET | `/api/v1/missions` | Lista misiones | autenticado |
| POST/PUT/DELETE | `/api/v1/missions` | Gestiona misiones | admin |
| GET | `/api/v1/usermissions` | Lista misiones asignadas | autenticado |
| POST | `/api/v1/usermissions` | Asigna/completa mision de usuario | autenticado |

## Arquitectura

El flujo principal es:

```text
HTTP request -> routes -> handlers -> services -> repositories -> GORM -> PostgreSQL/SQLite
```

`cmd/marketplace-api/main.go` abre la base, ejecuta migraciones, crea repositorios, inyecta servicios y registra rutas. Cada modulo mantiene handlers delgados, servicios con reglas de negocio y repositorios para persistencia.

Ver tambien: `docs/arquitectura.md`.

## Pruebas y CI

Ejecutar localmente:

```bash
go test ./...
go test ./... -cover
```

El pipeline `.github/workflows/ci.yml` ejecuta:

```text
go build ./...
go vet ./...
go test -v ./...
```

## Entregables Hito 3

- README profesional: este archivo.
- Coleccion Postman: `postman/marketplace-uleam.postman_collection.json`.
- Diagrama de arquitectura: `docs/arquitectura.md`.
- Documento de cierre: `docs/cierre.md`.
