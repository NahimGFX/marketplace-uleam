package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	Puerto           string
	DBDriver         string
	DBDSN            string
	RutaDB           string
	Storage          string
	JWTSecreto       string
	JWTDuracion      time.Duration
	HTTPReadTimeout  time.Duration
	HTTPWriteTimeout time.Duration
}

func Cargar() (Config, error) {
	_ = cargarDotEnv(".env")

	jwtDuracion, err := duracion("JWT_DURACION", 24*time.Hour)
	if err != nil {
		return Config{}, err
	}

	readTimeout, err := duracion("HTTP_READ_TIMEOUT", 10*time.Second)
	if err != nil {
		return Config{}, err
	}

	writeTimeout, err := duracion("HTTP_WRITE_TIMEOUT", 10*time.Second)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		Puerto:           valor("PUERTO", ":8080"),
		DBDriver:         strings.ToLower(valor("DB_DRIVER", "sqlite")),
		DBDSN:            valor("DB_DSN", ""),
		RutaDB:           valor("RUTA_DB", "marketplace.db"),
		Storage:          strings.ToLower(valor("STORAGE", "gorm")),
		JWTSecreto:       valor("JWT_SECRETO", "cambia-esto-por-un-secreto-largo-y-unico"),
		JWTDuracion:      jwtDuracion,
		HTTPReadTimeout:  readTimeout,
		HTTPWriteTimeout: writeTimeout,
	}

	if cfg.Puerto != "" && !strings.HasPrefix(cfg.Puerto, ":") {
		cfg.Puerto = ":" + cfg.Puerto
	}

	if cfg.DBDriver != "sqlite" && cfg.DBDriver != "postgres" {
		return Config{}, fmt.Errorf("DB_DRIVER invalido: %s", cfg.DBDriver)
	}
	if cfg.DBDriver == "postgres" && strings.TrimSpace(cfg.DBDSN) == "" {
		return Config{}, fmt.Errorf("DB_DSN es obligatorio cuando DB_DRIVER=postgres")
	}
	if cfg.Storage != "gorm" && cfg.Storage != "sqlc" {
		return Config{}, fmt.Errorf("STORAGE invalido: %s", cfg.Storage)
	}

	return cfg, nil
}

func valor(nombre, defecto string) string {
	if v := strings.TrimSpace(os.Getenv(nombre)); v != "" {
		return v
	}
	return defecto
}

func duracion(nombre string, defecto time.Duration) (time.Duration, error) {
	v := strings.TrimSpace(os.Getenv(nombre))
	if v == "" {
		return defecto, nil
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return 0, fmt.Errorf("%s invalido: %w", nombre, err)
	}
	return d, nil
}

func cargarDotEnv(ruta string) error {
	archivo, err := os.Open(ruta)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := strings.TrimSpace(scanner.Text())
		if linea == "" || strings.HasPrefix(linea, "#") {
			continue
		}

		nombre, valor, ok := strings.Cut(linea, "=")
		if !ok {
			continue
		}

		nombre = strings.TrimSpace(nombre)
		valor = strings.Trim(strings.TrimSpace(valor), `"'`)
		if nombre == "" {
			continue
		}
		if _, existe := os.LookupEnv(nombre); !existe {
			_ = os.Setenv(nombre, valor)
		}
	}

	return scanner.Err()
}
