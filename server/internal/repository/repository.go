package repository

import (
	entity "github.com/bolotskihUlyanaa/pinger/server/internal/models"
)

// Контракт для слоя репозиторий
type Pinger interface {
	GetAll() ([]entity.Ping, error) // Чтение всех записей
	Update(item entity.Ping) error  // Обновление записи в базе данных
	Insert(item entity.Ping) error  // Вставка записи в базу данных
}

// Структура для работы с базой данных
type Repository struct {
	Pinger
}

func NewRepository(pinger Pinger) *Repository {
	return &Repository{pinger}
}
