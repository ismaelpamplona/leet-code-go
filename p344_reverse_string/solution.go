package p344_reverse_string

func reverseString(s []byte) {
	l, r := 0, len(s)-1 // left and right pointers

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
