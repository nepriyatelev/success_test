package utils

import (
	"github.com/gookit/slog"
	"io"
	"net/http"
)

func ReadBody(resp *http.Response) (string, error) {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	slog.Debug("resp.Body = ", string(b))
	return string(b), nil
}
