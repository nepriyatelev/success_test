package responses

import (
	"github.com/gookit/slog"
	"github.com/nepriyatelev/success_test.git/internal/scrubber"
	"github.com/nepriyatelev/success_test.git/utils"
	"net/http"
	"strconv"
	"strings"
)

const (
	Passed = "passed"
)

func ParseGetResponse(resp *http.Response) (map[string]string, int, error) {
	bodyStr, err := utils.ReadBody(resp)
	if err != nil {
		slog.Error("Ошибка чтения тела ответа: ", err)
		return nil, 0, err
	}
	title := scrubber.TitleScrubber(bodyStr)
	slog.Debug("title = ", title)
	titleSplit := strings.Split(title, " ")
	slog.Debug("titleSplit = ", titleSplit)
	values := scrubber.TestFieldScrubber(bodyStr)

	question, err := strconv.Atoi(titleSplit[len(titleSplit)-1])
	return values, question, err
}

func ParsePostResponse(resp *http.Response) map[string]string {
	bodyStr, err := utils.ReadBody(resp)
	if err != nil {
		slog.Error("Ошибка чтения тела ответа: ", err)
		return nil
	}
	values := scrubber.TestFieldScrubber(bodyStr)
	slog.Debug("values = ", values)
	title := scrubber.TitleScrubber(bodyStr)
	slog.Debug("title = ", title)

	if len(values) == 0 && title == Passed {
		values["status"] = "Test successfully passed"
	}

	return values
}
