package valid_anagram

/* Assumes that:
1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters. */
func IsAnagram(s string, t string) bool {
	rune_s := []rune(s)
	rune_t := []rune(t)

	if len(s) != len(t) {
		return false
	}

	rune_counts := make( map[rune]int )

	for index := 0; index < len(rune_s); index++ {
		s_rune := rune_s[index]
		_, exists := rune_counts[ s_rune ]
		if !exists {
			rune_counts[ s_rune ] = 0
		}
		rune_counts[ s_rune ]++

		t_rune := rune_t[index]
		_, exists = rune_counts[ t_rune ]
		if !exists {
			rune_counts[ t_rune ] = 0
		}
		rune_counts[ t_rune ]--
	}

	for _, value := range rune_counts {
		if value != 0 {
			return false
		}
	}

	return true
}
