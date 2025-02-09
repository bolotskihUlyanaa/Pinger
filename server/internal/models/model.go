package models

import "time"

// Структура, описывающая ping
type Ping struct {
	Ip   string // IP-адрес
	Time int64  // Время пинга в нс
	Date string // Дата последней успешной попытки
}

// Для форматирования даты
const layout = "2 Jan 2006 15:04:05"

func NewPing(ip string, time_ping int64, date_ping time.Time) Ping {
	return Ping{ip, time_ping, date_ping.Format(layout)}
}
