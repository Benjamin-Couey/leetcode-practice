/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/longest-common-prefix/description/
*/
package longest_common_prefix

/*
LongestCommonPrefix returns the longest common prefix among strings in strs.
LongestCommonPrefix assumes that:
1 <= strs.length <= 200,
0 <= strs[i].length <= 200,
and strs[i] consists of only lowercase English letters.
*/
func LongestCommonPrefix(strs []string) string {
	lcp := ""

	// Convert strings to lists of runes so they can be indexed
	var rune_lists [][]rune
	for _, str := range strs {
		rune_lists = append(rune_lists, []rune(str))
	}

	// Label the outer loop so an inner loop can break out of it
index_loop:
	for index := 0; index < 200; index++ {

		current_rune := ' '

		if index < len(rune_lists[0]) {
			current_rune = rune_lists[0][index]
		} else {
			break
		}

		for _, rune_list := range rune_lists {
			if index >= len(rune_list) {
				break index_loop
			}
			if rune_list[index] != current_rune {
				break index_loop
			}
		}

		lcp += string(current_rune)
	}

	return lcp
}
