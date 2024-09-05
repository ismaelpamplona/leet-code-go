package p392_is_subsequence

func isSubsequence(s string, t string) bool {
	si, ti := 0, 0

	for si < len(s) && ti < len(t) {
		if s[si] == t[ti] {
			si++
			ti++
		}
		ti++
	}

	return si == len(s)
}
