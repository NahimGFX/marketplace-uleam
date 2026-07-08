package service

import (
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"marketplace-api/internal/models"
	"marketplace-api/internal/storage"
)

const (
	jwtSecretoDefecto  = "marketplace-uleam-secreto-demo-cambiar-en-S12"
	jwtDuracionDefecto = 24 * time.Hour
)

// Claims es el contenido del JWT: el ID del usuario + los campos estandar (exp, iat).
type Claims struct {
	UsuarioID int `json:"uid"`
	jwt.RegisteredClaims
}

// AuthService concentra TODA la logica de autenticacion: hashing de contrasenas
// (bcrypt) y generacion/validacion de JWT. El handler y el middleware no saben
// de bcrypt ni de firmas: solo llaman a este servicio. Esa es la razon de ser
// de la capa de servicio.
type AuthService struct {
	repo        storage.UserRepository
	jwtSecreto  []byte
	jwtDuracion time.Duration
}

func NuevoAuthService(repo storage.UserRepository) *AuthService {
	return NuevoAuthServiceConJWT(repo, jwtSecretoDesdeEnv(), jwtDuracionDesdeEnv())
}

func NuevoAuthServiceConJWT(repo storage.UserRepository, secreto string, duracion time.Duration) *AuthService {
	if strings.TrimSpace(secreto) == "" {
		secreto = jwtSecretoDefecto
	}
	if duracion <= 0 {
		duracion = jwtDuracionDefecto
	}
	return &AuthService{
		repo:        repo,
		jwtSecreto:  []byte(secreto),
		jwtDuracion: duracion,
	}
}

// Registrar crea un usuario nuevo en users con la contrasena hasheada (bcrypt).
func (s *AuthService) Registrar(email, password string) (models.User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	password = strings.TrimSpace(password)
	if email == "" || password == "" {
		return models.User{}, ErrCredencialesInvalidas
	}
	if _, existe := s.repo.BuscarUsuarioPorEmail(email); existe {
		return models.User{}, ErrEmailEnUso
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	u, err := s.repo.CrearUsuario(models.User{
		Name:       nombreDesdeEmail(email),
		Email:      email,
		Password:   string(hash),
		Level:      1,
		Reputation: 0,
	})
	if err != nil {
		return models.User{}, err
	}

	u.Password = ""
	return u, nil
}

// Login verifica las credenciales y, si son validas, devuelve un JWT firmado.
func (s *AuthService) Login(email, password string) (string, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	password = strings.TrimSpace(password)

	u, existe := s.repo.BuscarUsuarioPorEmail(email)
	if !existe {
		return "", ErrCredencialesInvalidas
	}
	if !s.passwordValido(u, password) {
		return "", ErrCredencialesInvalidas
	}

	return s.generarToken(u)
}

func (s *AuthService) passwordValido(u models.User, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err == nil {
		return true
	}

	if u.Password != password {
		return false
	}

	// Compatibilidad con datos antiguos guardados en texto plano.
	if hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err == nil {
		_ = s.repo.ActualizarPasswordUsuario(u.ID, string(hash))
	}
	return true
}

// generarToken arma y firma el JWT con el ID del usuario y la expiracion.
func (s *AuthService) generarToken(u models.User) (string, error) {
	claims := Claims{
		UsuarioID: u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.jwtDuracion)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecreto)
}

// ValidarToken verifica firma y expiracion, y devuelve el ID del usuario.
// Lo usa el middleware de autenticacion: el JWT vive aqui, no en el middleware.
func (s *AuthService) ValidarToken(tokenStr string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrCredencialesInvalidas
		}
		return s.jwtSecreto, nil
	})
	if err != nil || !token.Valid {
		return 0, ErrCredencialesInvalidas
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, ErrCredencialesInvalidas
	}
	return claims.UsuarioID, nil
}

func nombreDesdeEmail(email string) string {
	nombre, _, ok := strings.Cut(email, "@")
	if !ok || strings.TrimSpace(nombre) == "" {
		return email
	}
	return nombre
}

func jwtSecretoDesdeEnv() string {
	if v := strings.TrimSpace(os.Getenv("JWT_SECRETO")); v != "" {
		return v
	}
	return jwtSecretoDefecto
}

func jwtDuracionDesdeEnv() time.Duration {
	if v := strings.TrimSpace(os.Getenv("JWT_DURACION")); v != "" {
		if d, err := time.ParseDuration(v); err == nil && d > 0 {
			return d
		}
	}
	return jwtDuracionDefecto
}
