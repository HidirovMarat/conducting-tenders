package config

import (
	"log"
	"os"
)

/*
Подключение к БД

	host=rc1b-5xmqy6bq501kls4m.mdb.yandexcloud.net
	port=6432
	dbname=cnrprod1725728083-team-77059
	user=cnrprod1725728083-team-77059
	password=cnrprod1725728083-team-77059
	target_session_attrs=read-write

postgres://cnrprod1725728083-team-77059:cnrprod1725728083-team-77059@cnrprod1725728083-team-77059:5432/cnrprod1725728083-team-77059
*/

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
		//config.ServerAddress = "rc1b-5xmqy6bq501kls4m.mdb.yandexcloud.net:6432"
		log.Fatal("не установлен SERVER_ADDRESS")
	}

	config.PostgresConn = os.Getenv("POSTGRES_CONN")
	if config.PostgresConn == "" {
		//config.PostgresConn = "postgres://cnrprod1725728083-team-77059:cnrprod1725728083-team-77059@cnrprod1725728083-team-77059:5432/cnrprod1725728083-team-77059"
		log.Fatal("не установлен POSTGRES_CONN")
	}

	config.PostgresJdbcUrl = os.Getenv("POSTGRES_JDBC_URL")
	if config.PostgresJdbcUrl == "" {
		//config.PostgresJdbcUrl = ""
		log.Fatal("не установлен POSTGRES_JDBC_URL")
	}

	config.PostgresUsername = os.Getenv("POSTGRES_USERNAME")
	if config.PostgresUsername == "" {
		//config.PostgresUsername = "cnrprod1725728083-team-77059"
		log.Fatal("не установлен POSTGRES_USERNAME")
	}

	config.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	if config.PostgresPassword == "" {
		//config.PostgresPassword = "cnrprod1725728083-team-77059"
		log.Fatal("не установлен POSTGRES_PASSWORD")
	}

	config.PostgresHost = os.Getenv("POSTGRES_HOST")
	if config.PostgresHost == "" {
		//log.Fatal("не установлен POSTGRES_HOST")
	}

	config.PostgresPort = os.Getenv("POSTGRES_PORT")
	if config.PostgresPort == "" {
		//config.PostgresPort = "5432"
		log.Fatal("не установлен POSTGRES_PORT")
	}

	config.PostgresDatabase = os.Getenv("POSTGRES_DATABASE")
	if config.PostgresDatabase == "" {
		//config.PostgresDatabase = "cnrprod1725728083-team-77059"
		log.Fatal("не установлен POSTGRES_DATABASE")
	}

	return &config, nil
}
