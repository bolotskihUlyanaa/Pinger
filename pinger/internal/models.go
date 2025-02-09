package internal

import "time"

// Структура, описывающая ping
type Ping struct {
	Ip   string    // IP-адрес
	Time int64     // Время пинга в нс
	Date time.Time // Дата последней успешной попытки
}

// Структура, описывающая данные, полученные из конфигурационного файла
type Pinger struct {
	Server   string   // Адрес сервера
	Services []string // Список docker-контейнеров
	Delay    int      // Задержка между пингами (в секундах)
	Packages int      // Количество пакетов, отправляемых во время пинга
	Timeout  int      // Таймаут пинга
}
