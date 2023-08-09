package requests

import (
	"github.com/gookit/slog"
	"github.com/nepriyatelev/success_test.git/pkg"
	"net/http"
	"net/url"
)

func GetRequest(url string, client *pkg.Client) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("Ошибка создания GET-запроса у клиента ", client.ID, ": ", err)
		return nil, err
	}

	slog.Info("Клиент ", client.ID, " создал GET-запрос", req)
	return client.Client.Do(req)
}

func PostRequest(urlStr string, data map[string]string, client *pkg.Client) (*http.Response, error) {
	formData := url.Values{}

	for k, v := range data {
		formData.Set(k, v)
	}

	slog.Info("Клиент ", client.ID, " создал POST-запрос на ", urlStr, " с данными: ", formData)
	return client.Client.PostForm(urlStr, formData)
}
