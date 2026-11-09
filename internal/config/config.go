package config

import "os"

type Config struct {
	Addr string
}

// Load читает настройки из переменных окружения.
func Load() Config {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	return Config{Addr: addr}
}
