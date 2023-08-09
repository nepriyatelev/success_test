package pkg

import (
	"github.com/gookit/slog"
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	ID     int
	Client *http.Client
}

func NewClientWithJar() (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		slog.Error("Ошибка создания cookiejar: ", err)
		return nil, err
	}
	return &Client{
		ID:     0,
		Client: &http.Client{Jar: jar},
	}, nil
}
