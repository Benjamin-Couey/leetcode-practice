/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/is-subsequence/description/
*/
package is_subsequence

/*
IsSubsequence reports whether s is a subsequence of t.
IsSubsequence assumes that:
0 <= s.length <= 100,
0 <= t.length <= 104,
and s and t consist only of lowercase English letters.
*/
func IsSubsequence(s string, t string) bool {

	rune_s := []rune(s)
	s_index := 0

	for _, rune := range t {
		if rune_s[s_index] == rune {
			s_index += 1
		}
		if s_index >= len(s) {
			return true
		}
	}
	return false
}
