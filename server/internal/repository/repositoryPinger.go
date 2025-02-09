package repository

import (
	"database/sql"
	"time"

	entity "github.com/bolotskihUlyanaa/pinger/server/internal/models"
)

// Реализация интерфейса Pinger - слой репозиторий для работы с базой данных
type PingerRepository struct {
	db *sql.DB
}

func NewPingerRepository(db *sql.DB) *PingerRepository {
	return &PingerRepository{db: db}
}

// Вставка в базу данных
func (p *PingerRepository) Insert(item entity.Ping) error {
	_, err := p.db.Exec("INSERT INTO ping_data (ip, time_ping, date_ping) "+
		"VALUES ($1, $2, $3) ON CONFLICT DO NOTHING;", item.Ip, item.Time, item.Date)
	return err
}

// Обновления записи в базе данных
func (p *PingerRepository) Update(item entity.Ping) error {
	_, err := p.db.Exec("UPDATE ping_data SET "+
		"time_ping = $1, date_ping=$2 WHERE ip = $3;", item.Time, item.Date, item.Ip)
	return err
}

// Получение данных всей таблицы
func (p *PingerRepository) GetAll() ([]entity.Ping, error) {
	rows, err := p.db.Query("SELECT * FROM ping_data;")
	if err != nil {
		return nil, err
	}
	var list []entity.Ping // Слайс пингов
	var ip string
	var time_ping int64
	var date_ping time.Time
	// Обработка каждой строки ответа
	// на основе каждой строки создается структура пинг и добавляется в слайс пингов
	for rows.Next() {
		if err = rows.Scan(&ip, &time_ping, &date_ping); err != nil {
			return nil, err
		}
		list = append(list, entity.NewPing(ip, time_ping, date_ping))
	}
	return list, nil
}
