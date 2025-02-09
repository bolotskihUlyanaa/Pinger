package service

import (
	entity "github.com/bolotskihUlyanaa/pinger/server/internal/models"
	repository "github.com/bolotskihUlyanaa/pinger/server/internal/repository"
)

// Структура обращается к слою репозиторий для обработки данных о пинге
type PingerService struct {
	repos *repository.Repository
}

func NewPingerService(repository *repository.Repository) *PingerService {
	return &PingerService{repository}
}

// Функция для получения данных о пинге для всех адресов
func (p *PingerService) GetAll() ([]entity.Ping, error) {
	return p.repos.GetAll()
}

// Функция для обновления данных о пинге
func (p *PingerService) Update(item entity.Ping) error {
	return p.repos.Update(item)
}

// Функция для создания записи о пинге
func (p *PingerService) Create(item entity.Ping) error {
	return p.repos.Insert(item)
}
