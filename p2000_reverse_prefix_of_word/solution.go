package p2000_reverse_prefix_of_word

func reversePrefix(word string, ch byte) string {
	bytes := []byte(word)

	left, right := 0, 0

	for i, c := range bytes {
		if c == ch {
			right = i
			break
		}
	}

	if right == -1 {
		return word
	}

	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}

	return string(bytes)

}
