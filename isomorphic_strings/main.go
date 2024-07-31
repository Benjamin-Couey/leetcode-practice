package isomorphic_strings

/* Assumes that
1 <= s.length <= 5 * 10^4
t.length == s.length
s and t consist of any valid ascii character. */
func IsIsomorphic(s string, t string) bool {

	rune_s := []rune(s)
	rune_t := []rune(t)
	s_to_t_map := make(map[rune]rune)

	for index := 0; index < len(s); index++ {
		t_rune, exists := s_to_t_map[rune_s[index]]

		if exists {
			if t_rune != rune_t[index] {
				return false
			}
		} else {
			s_to_t_map[rune_s[index]] = rune_t[index]
		}
	}

	return true
}
