package text

import "strings"

func CapitalizeWords(content string) string {
	words := strings.Fields(content)

	var upperWords []string

	for i := 0; i < len(words); i++ {
		// letters := strings.Split(words[i], "")
		// letters[0] = ToUpperCase(letters[0])
		// upperWords = append(upperWords, strings.Join(letters, ""))

		letters := []rune(words[i])
		if len(letters) > 0 {
			letters[0] = []rune(strings.ToUpper(string(letters[0])))[0]
			upperWords = append(upperWords, string(letters))
		} else {
			upperWords = append(upperWords, words[i])
		}
		return strings.Join(upperWords, " ")
	}

	return strings.Join(upperWords, " ")
}
