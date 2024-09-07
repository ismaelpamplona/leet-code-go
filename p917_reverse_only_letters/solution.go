package p917_reverse_only_letters

import (
	"unicode"
)

func reverseOnlyLetters1(s string) string {
	chars := []rune(s)
	left, right := 0, len(chars)-1

	for left < right {
		if isLetter(chars[left]) && isLetter(chars[right]) {
			chars[left], chars[right] = chars[right], chars[left]
			left++
			right--
		} else if !isLetter(chars[left]) {
			left++
		} else if !isLetter(chars[right]) {
			right--
		}
	}

	return string(chars)

}

func isLetter(c rune) bool {
	if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
		return true
	}
	return false
}

func reverseOnlyLetters2(s string) string {
	chars := []rune(s)
	left, right := 0, len(chars)-1

	for left < right {
		if unicode.IsLetter(chars[left]) && unicode.IsLetter(chars[right]) {
			chars[left], chars[right] = chars[right], chars[left]
			left++
			right--
		} else if !unicode.IsLetter(chars[left]) {
			left++
		} else if !unicode.IsLetter(chars[right]) {
			right--
		}
	}

	return string(chars)
}
