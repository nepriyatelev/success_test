package scrubber

import (
	"github.com/gookit/slog"
	"golang.org/x/net/html"
	"strings"
)

const (
	Title  = "title"
	Select = "select"
	Input  = "input"
	Name   = "name"
	Type   = "type"
	Text   = "text"
	Test   = "test"
	Radio  = "radio"
	Value  = "value"
	Option = "option"
)

func TitleScrubber(s string) string {
	tkn := html.NewTokenizer(strings.NewReader(s))
	for {
		tknType := tkn.Next()
		switch tknType {
		case html.ErrorToken:
			slog.Error("Error token")
			return ""
		case html.StartTagToken:
			token := tkn.Token()
			if token.Data == Title {
				tknType = tkn.Next()
				if tknType == html.TextToken {
					return tkn.Token().Data
				}
			}
		}
	}
}

func TestFieldScrubber(r string) map[string]string {
	tkn := html.NewTokenizer(strings.NewReader(r))

	values := make(map[string]string)

	for {
		tknType := tkn.Next()
		switch tknType {
		case html.ErrorToken:
			return values
		case html.StartTagToken:
			token := tkn.Token()
			if token.Data == Select {
				var name string
				if len(token.Attr) > 0 {
					for _, attr := range token.Attr {
						if attr.Key == Name {
							name = attr.Val
							break
						}
					}
				}
				value := findLongestOptionValue(tkn)
				values[name] = value
			} else if token.Data == Input {
				for _, attr := range token.Attr {
					if attr.Key == Type && attr.Val == Text {
						name := getNameText(token)
						values[name] = Test
					} else if attr.Key == Type && attr.Val == Radio {
						name, value := getNameAndValue(token)
						if v, ok := values[name]; ok {
							if len(v) < len(value) {
								values[name] = value
							}
						} else {
							values[name] = value
						}
					}
				}
			}
		}
	}
}

func getNameText(tkn html.Token) string {
	var value string
	for _, a := range tkn.Attr {
		if a.Key == Name {
			value = a.Val
		}
	}
	return value
}

func getNameAndValue(tkn html.Token) (string, string) {
	var name string
	var value string
	for _, a := range tkn.Attr {
		if a.Key == Name {
			name = a.Val
		} else if a.Key == Value && a.Val != "" {
			value = a.Val
		}
	}
	return name, value
}

func findLongestOptionValue(tkn *html.Tokenizer) string {
	var longestValue string

	for {
		tknType := tkn.Next()

		if tknType == html.ErrorToken {
			slog.Error("Error token")
			break
		}

		token := tkn.Token()
		if tknType == html.StartTagToken && token.Data == Option {
			for _, attr := range token.Attr {
				if attr.Key == Value {
					value := attr.Val
					if len(value) > len(longestValue) {
						longestValue = value
					}
				}
			}
		} else if tknType == html.EndTagToken && token.Data == Select {
			break
		}
	}
	return longestValue
}
