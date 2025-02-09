package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Структура описывает базу данных
type PostgresConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgres(cfg PostgresConfig) (*sql.DB, error) {
	// Открытие соединения с базой данных
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, fmt.Errorf(
			"ошибка при открытии соединения с базой данных: %w", err)
	}
	// Проверка подключения
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ошибка соединения с базой данных: %w", err)
	}
	return db, nil
}
