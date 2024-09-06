package p557_reverse_words_in_a_string_iii

import (
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	ans := make([]string, len(words))

	for i, word := range words {
		chars := []rune(word)
		l, r := 0, len(chars)-1
		for l < r {
			chars[l], chars[r] = chars[r], chars[l]
			l++
			r--
		}
		ans[i] = string(chars)
	}

	return strings.Join(ans, " ")
}
