package config

import (
	"log"
	"os"
)

type (
	Config struct {
		ServerAddress string // адрес и порт, который будет слушать HTTP сервер при запуске. Пример: 0.0.0.0:8080.

		PostgresConn string // URL-строка для подключения к PostgreSQL в формате postgres://{username}:{password}@{host}:{5432}/{dbname}.

		PostgresJdbcUrl string // JDBC-строка для подключения к PostgreSQL в формате jdbc:postgresql://{host}:{port}/{dbname}.

		PostgresUsername string // имя пользователя для подключения к PostgreSQL.

		PostgresPassword string // пароль для подключения к PostgreSQL.

		PostgresHost string // хост для подключения к PostgreSQL (например, localhost).

		PostgresPort string // порт для подключения к PostgreSQL (например, 5432).

		PostgresDatabase string // имя базы данных PostgreSQL, которую будет использовать приложение.

	}
)

func NewConfig() (*Config, error) {
	var config Config

	config.ServerAddress = os.Getenv("SERVER_ADDRESS")
	if config.ServerAddress == "" {
		log.Fatal("не установлен SERVER_ADDRESS")
	}

	config.PostgresConn = os.Getenv("POSTGRES_CONN")
	if config.PostgresConn == "" {
		log.Fatal("не установлен POSTGRES_CONN")
	}

	config.PostgresUsername = os.Getenv("POSTGRES_USERNAME")
	if config.PostgresUsername == "" {
		log.Fatal("не установлен POSTGRES_USERNAME")
	}

	config.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	if config.PostgresPassword == "" {
		log.Fatal("не установлен POSTGRES_PASSWORD")
	}

	config.PostgresHost = os.Getenv("POSTGRES_HOST")
	if config.PostgresHost == "" {
		log.Fatal("не установлен POSTGRES_HOST")
	}

	config.PostgresPort = os.Getenv("POSTGRES_PORT")
	if config.PostgresPort == "" {
		log.Fatal("не установлен POSTGRES_PORT")
	}

	config.PostgresDatabase = os.Getenv("POSTGRES_DATABASE")
	if config.PostgresDatabase == "" {
		log.Fatal("не установлен POSTGRES_DATABASE")
	}

	return &config, nil
}
