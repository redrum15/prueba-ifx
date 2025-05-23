package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl                  string
	JWTSecret              string
	UserPassw              string
	JWTExpirationInSeconds int64
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		DBUrl:     getEnv("DB_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", ""),
		UserPassw: getEnv("USER_PASSW", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
