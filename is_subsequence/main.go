package is_subsequence

/* Assumes that:
0 <= s.length <= 100
0 <= t.length <= 104
s and t consist only of lowercase English letters. */
func IsSubsequence(s string, t string) bool {

		rune_s := []rune( s )
		s_index := 0

		for _, rune := range( t ) {
				if rune_s[ s_index ] == rune {
						s_index += 1
				}
				if s_index >= len( s ) {
						return true
				}
		}
		return false
}
