package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	probing "github.com/prometheus-community/pro-bing"
	yaml "gopkg.in/yaml.v3"
)

// Функция инициализирует структуру Pinger из конфигурационного файла
func NewPinger() *Pinger {
	var pinger Pinger
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(fmt.Errorf("ошибка при чтении конфигурационного файла: %w", err))
	}
	err = yaml.Unmarshal(yamlFile, &pinger)
	if err != nil {
		panic(fmt.Errorf(
			"ошибка при декодировании конфигурационного файла: %w", err))
	}
	return &pinger
}

// Функция пингует service и отправляет данные на сервер
// между пингами есть заданный промежуток времени
func (p *Pinger) PingSend(service string) {
	method := http.MethodPost
	for {
		data, err := p.ping(service)
		if err != nil {
			log.Println("Ошибка во время пинга:", err)
		} else {
			err = p.send(*data, method)
			if err != nil {
				log.Println("Ошибка при отправке данных на сервер:", err)
			}
			if method != http.MethodPut {
				method = http.MethodPut
			}
		}
		time.Sleep(time.Duration(p.Delay) * time.Second)
	}
}

// функиця пингует service и возвращает информацию о пинге
func (p *Pinger) ping(service string) (*Ping, error) {
	pinger, err := probing.NewPinger(service)
	if err != nil {
		return nil, err
	}
	pinger.Count = p.Packages                               // Количество пакетов, которые нужно отправить
	pinger.Timeout = time.Duration(p.Timeout) * time.Second // Таймаут
	err = pinger.Run()
	if err != nil {
		return nil, err
	}
	// Ошибка если количество отправленных пакетов не равно количеству полученных
	if pinger.PacketsSent != pinger.PacketsRecv {
		return nil, fmt.Errorf("отправлено: %d, получено: %d пакетов",
			pinger.PacketsSent, pinger.PacketsRecv)
	}
	stats := pinger.Statistics()
	return &Ping{stats.IPAddr.String(), stats.AvgRtt.Nanoseconds(), time.Now()}, nil
}

// Функция отправки данных о пинге ping на сервер в формате json
// post или put в зависиммости от переданного метода httpMethod
func (p *Pinger) send(ping Ping, httpMethod string) error {
	data, err := json.Marshal(ping)
	if err != nil {
		return err
	}
	client := http.Client{}
	req, err := http.NewRequest(httpMethod, p.Server, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}
