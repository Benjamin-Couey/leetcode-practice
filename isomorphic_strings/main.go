/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/isomorphic-strings/description/
*/
package isomorphic_strings

/*
IsIsomorphic reports whether s and t are isomorphic. s and t are isomorphic if
the characters in s can be replaced to get t.
IsIsomorphic assumes that:
no two characters may map to the same character, but a character may map to itself,
1 <= s.length <= 5 * 10^4,
t.length == s.length,
and s and t consist of any valid ascii character.
*/
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
