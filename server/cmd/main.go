package main

import (
	"fmt"
	"net/http"
	"os"

	handler "github.com/bolotskihUlyanaa/pinger/server/internal/handler"
	repository "github.com/bolotskihUlyanaa/pinger/server/internal/repository"
	service "github.com/bolotskihUlyanaa/pinger/server/internal/service"
	yaml "gopkg.in/yaml.v3"
)

// Структура для чтения конфигурационного файла
type Config struct {
	Server         int                       // Порт сервера
	Client         string                    // Ip адрес клиента
	PostgresConfig repository.PostgresConfig // Конфигурация базы данных
}

// Функция для инициализации конфигурационного файла
func initConfig(cfg *Config) {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(fmt.Errorf("ошибка при чтении конфигурационного файла: %w", err))
	}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		panic(fmt.Errorf(
			"ошибка при декодировании конфигурационного файла: %w", err))
	}
}

func main() {
	// Инициализация сервера из конфигурационного файла
	var config Config
	initConfig(&config)

	// Подключение к базе данных
	db, err := repository.NewPostgres(config.PostgresConfig)
	if err != nil {
		panic(err)
	}

	// Внедрение зависимостей
	repository := repository.NewRepository(repository.NewPingerRepository(db))
	service := service.NewService(service.NewPingerService(repository))
	h := handler.NewHandler(service, config.Client)

	// Регистрация функции-обработчика для URL "/"
	http.HandleFunc("/", h.ServeHTTP)

	// Запустить HTTP-сервер
	go http.ListenAndServe(fmt.Sprintf(":%d", config.Server), nil)
	select {}
}
