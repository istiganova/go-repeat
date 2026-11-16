package text

import "strings"

func CountWords(content string) int {
	text := strings.Fields(content)

	return len(text)
}
