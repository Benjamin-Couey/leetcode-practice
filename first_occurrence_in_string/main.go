/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
*/
package first_occurrence_in_string

/*
StrStr returns the index of the first occurrence of the substring needle in
the string haystack. Returns -1 if needle is not present in haystack.
StrStr assumes that:
1 <= haystack.length, needle.length <= 10^4,
and haystack and needle consist of only lowercase English characters.
*/
func StrStr(haystack string, needle string) int {
	rune_haystack := []rune(haystack)
	rune_needle := []rune(needle)
	first_occurrence := -1
	needle_index := 0
	for index := 0; index < len(haystack); index++ {
		if rune_haystack[index] == rune_needle[needle_index] {
			if first_occurrence == -1 {
				first_occurrence = index
			}
			needle_index++
			if needle_index >= len(needle) {
				return first_occurrence
			}
		} else {
			first_occurrence = -1
			needle_index = 0
		}
	}

	return first_occurrence
}
