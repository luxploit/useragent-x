package uax

import (
	"slices"
	"strings"
	"unicode"

	"golang.org/x/text/language"
)

var forbiddenTokens = []string{
	"U", "compatible", "SLCC2", "Touch", "X11", "like", "I", "Nav", "wv",
}

func TokenizeUA(userAgent string) (props []string, botURL string) {
	tokens := []string{}
	url := strings.Builder{}
	buffer := strings.Builder{}

	recordingUrl := false

	tokenize := func(ch rune) {
		if buffer.Len() <= 0 {
			return
		}

		token := strings.TrimSpace(buffer.String())
		token = strings.Trim(token, "+")

		if slices.Contains(forbiddenTokens, token) {
			buffer.Reset()
			return
		}

		_, err := language.Parse(token)
		//@NOTE: valid language code, dont want that here
		if err == nil && (len(token) == 2 /*short code*/ || (len(token) == 5 /*extended*/) && (strings.Contains(token, "-") || strings.Contains(token, "_"))) {
			buffer.Reset()
			return
		}

		if recordingUrl && ch == ')' {
			recordingUrl = false
			buffer.Reset()
			return
		}

		if token == "http" || token == "https" {
			if !recordingUrl { // odd bug
				url.WriteString(token)
			}
			url.WriteRune(ch) // should be :
			recordingUrl = true
			return
		}

		println(token)
		tokens = append(tokens, token)
		buffer.Reset()
	}

	println(userAgent)
	for _, ch := range userAgent {
		switch ch {
		case '(', ')', '/', ':', ';', ',', ' ', '[', ']':
			tokenize(ch)
		default:
			if !unicode.IsControl(ch) {
				if recordingUrl {
					url.WriteRune(ch)
				} else {
					buffer.WriteRune(ch)
				}
			}
		}
	}
	tokenize('a') // hack to get it to also tokenize the last item in buffer

	return tokens, url.String()
}

func ParseUA(userAgent string) UserAgent {
	_, botUrl := TokenizeUA(userAgent)
	ua := UserAgent{
		RawString:  userAgent,
		CrawlerURL: botUrl,
	}

	// println("---")
	// idx := 0
	// for idx < len(tokens) {
	// 	println(tokens[idx])

	// 	if

	// 	idx++
	// }

	return ua
}
