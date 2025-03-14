package common

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

func Get(url string) (string, string, error) {
	// Отправляем GET-запрос
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close() // Закрываем тело ответа после завершения работы

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return "", resp.Status, fmt.Errorf("server returned non-200 status: %s", resp.Status)
	}

	// Читаем ответ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	return string(body), resp.Status, nil
}

func Post(url string, data []byte) (string, string, error) {
	// Отправляем POST-запрос
	resp, err := http.Post(url, "text/plain", bytes.NewBuffer(data))
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close() // Закрываем тело ответа после завершения работы

	// Читаем ответ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	return string(body), resp.Status, nil
}

func GetIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		parts := strings.Split(ip, ",")
		return strings.TrimSpace(parts[0])
	}

	ip = r.RemoteAddr
	host, _, err := net.SplitHostPort(ip)
	if err != nil {
		return ip
	}

	// Проверка на IPv6
	if strings.Contains(host, ":") {
		return host
	}
	return host
}
