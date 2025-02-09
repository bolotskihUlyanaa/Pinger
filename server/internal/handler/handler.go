package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	entity "github.com/bolotskihUlyanaa/pinger/server/internal/models"
	service "github.com/bolotskihUlyanaa/pinger/server/internal/service"
)

// Структура реализует слой контроллера
// Реализует интерфейс http.Handler
// Структура обращается к слою сервис для обработки данных о пинге
type Handler struct {
	*service.Service
	clientURL string // Для разрешения доступа клиента к ответу
}

func NewHandler(service *service.Service, clientURL string) *Handler {
	return &Handler{service, clientURL}
}

// Функция обрабатывает запрос и отправляет ответ
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Установка в заголовок информацию о том, что будет передаваться json
	w.Header().Set("Content-Type", "application/json")
	// Обработка запроса в зависимости от типа
	switch r.Method {
	case "GET":
		err := h.Get(w)
		if err != nil {
			log.Println("ERROR GET:", err)
		}
	case "PUT":
		err := h.PutPost(w, r.Body, h.Update)
		if err != nil {
			log.Println("ERROR PUT:", err)
		}
	case "POST":
		err := h.PutPost(w, r.Body, h.Create)
		if err != nil {
			log.Println("ERROR POST:", err)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// Функция для ответа на запрос GET
func (h *Handler) Get(w http.ResponseWriter) error {
	rows, err := h.GetAll() // Вызов метода сервиса
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	buf, err := json.Marshal(rows) // Кодирование данных в json
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	// Разрешить доступ к ответу определенному домену (домен клиента)
	w.Header().Set("Access-Control-Allow-Origin", h.clientURL)
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
	return nil
}

// Функция для ответа на запросы POST и PUT
// Отличаются обработчиком слоя сервис
// Функция обработчика передается в функцию в качестве параметра
func (h *Handler) PutPost(w http.ResponseWriter, body io.ReadCloser,
	action func(item entity.Ping) error) error {
	buf, err := io.ReadAll(body) // Чтение тела запроса
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	var data entity.Ping
	// Декодирование полученного json в структуру Ping
	err = json.Unmarshal(buf, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	err = action(data) // Функция сервиса
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}
