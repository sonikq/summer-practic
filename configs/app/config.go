package app

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"log"
	"os"
)

type Config struct {
	HTTP HTTPConf
	DB   PostgresConfig

	CtxTimeout int

	LogLevel string
}

type HTTPConf struct {
	Host string
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func Load(envFile string) (cfg Config) {

	if err := godotenv.Load(envFile); err != nil {
		log.Fatal(err)
		return
	}

	cfg.HTTP.Host = cast.ToString(os.Getenv("HTTP_HOST"))
	cfg.HTTP.Port = cast.ToString(os.Getenv("HTTP_PORT"))

	cfg.DB.Host = cast.ToString(os.Getenv("POSTGRES_HOST"))
	cfg.DB.Port = cast.ToString(os.Getenv("POSTGRES_PORT"))
	cfg.DB.Username = cast.ToString(os.Getenv("POSTGRES_USERNAME"))
	cfg.DB.Password = cast.ToString(os.Getenv("POSTGRES_PASSWORD"))
	cfg.DB.DBName = cast.ToString(os.Getenv("POSTGRES_DB"))
	cfg.DB.SSLMode = cast.ToString(os.Getenv("POSTGRES_SSLMODE"))

	cfg.CtxTimeout = cast.ToInt(os.Getenv("CTX_TIMEOUT"))

	cfg.LogLevel = cast.ToString(os.Getenv("LOG_LEVEL"))

	return cfg
}
