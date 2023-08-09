package app

import (
	"github.com/gookit/slog"
	"github.com/nepriyatelev/success_test.git/internal/requests"
	"github.com/nepriyatelev/success_test.git/internal/responses"
	"github.com/nepriyatelev/success_test.git/pkg"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	StartURL        = "http://147.78.65.149/start"
	QuestionURL     = "http://147.78.65.149/question/"
	RequestInterval = 3
)

func RunApp(n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go testSuccess(&wg, i+1)
	}
	wg.Wait()
}

func testSuccess(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	client, err := pkg.NewClientWithJar()
	if err != nil {
		slog.Error("Ошибка создания клиента: ", id, err)
		return
	}
	client.ID = id
	slog.Info("Клиент ", client.ID, " создан")

	resp, err := requests.GetRequest(StartURL, client)
	if err != nil {
		slog.Error("Ошибка GET-запроса у клиента ", client.ID, ": ", err)
		return
	}
	slog.Info("Клиент ", client.ID, " получил ответ на GET-запрос: ", resp)

	values, questions, err := responses.ParseGetResponse(resp)
	if err != nil || resp.StatusCode != http.StatusOK {
		slog.Error("Ошибка парсинга GET-ответа у клиента ", id, ": ", err, "статус код: ", resp.Status)
		return
	}
	slog.Info("Клиент ", client.ID, " получил данные для POST-запроса: ", values, questions)

	for i := 1; i <= questions; i++ {
		time.Sleep(RequestInterval * time.Second)
		qNumber := strconv.Itoa(i)
		url := QuestionURL + qNumber
		resp, err = requests.PostRequest(url, values, client)
		if err != nil || resp.StatusCode != http.StatusOK {
			slog.Error("Ошибка POST-запроса у клиента ", client.ID, ": ", err, " статус код: ", resp.Status)
			return
		}
		slog.Info("Клиент ", client.ID, " получил ответ на POST-запрос: ", resp)
		values = responses.ParsePostResponse(resp)
		slog.Info("Клиент ", client.ID, " получил данные для следующего POST-запроса: ", values)
	}

	if v, ok := values["status"]; ok {
		slog.Info("Клиент ", client.ID, " прошел тест со статусом: ", v)
	} else {
		slog.Info("Клиент ", client.ID, " не прошел тест")
	}
}
