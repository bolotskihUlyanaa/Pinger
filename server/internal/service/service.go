package service

import (
	entity "github.com/bolotskihUlyanaa/pinger/server/internal/models"
)

// Контракт для слоя сервис
type Pinger interface {
	GetAll() ([]entity.Ping, error) // Чтение всех записей
	Update(item entity.Ping) error  // Обновление данных о записи
	Create(item entity.Ping) error  // Создание записи
}

// Структура, представляющая сервис для работы с данными о пинге
type Service struct {
	Pinger
}

func NewService(pinger Pinger) *Service {
	return &Service{pinger}
}
