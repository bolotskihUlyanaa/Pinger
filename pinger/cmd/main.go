package main

import (
	ping "github.com/bolotskihUlyanaa/pinger/pinger/internal"
)

func main() {
	// pinger содержит список сервисов, настройки пинга, адрес сервера
	pinger := ping.NewPinger()

	// Операции с каждым сервисом переводятся в отдельный поток
	for _, service := range pinger.Services {
		go pinger.PingSend(service)
	}
	select {}
}
