/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/group-anagrams/description/
*/
package group_anagrams

import (
	"leetcode/valid_anagram"
)

/*
GroupAnagrams groups strings in strs that are anagrams of one another and
returns a slice of these groups.
GroupAnagrams assumes that:
anagrams use all letters exactly once,
the answer may be returned in any order,
1 <= strs.length <= 10^4,
0 <= strs[i].length <= 100,
and strs[i] consists of lowercase English letters.
*/
func GroupAnagrams(strs []string) [][]string {
	grouped_anagrams := make([][]string, 0)

	for _, str := range strs {
		if len(grouped_anagrams) < 1 {
			grouped_anagrams = append(grouped_anagrams, []string{str})
		} else {
			index := 0
			found_anagram := false
			for !found_anagram && index < len(grouped_anagrams) {
				anagrams := grouped_anagrams[index]
				if valid_anagram.IsAnagram(str, anagrams[0]) {
					anagrams = append(anagrams, str)
					grouped_anagrams[index] = anagrams
					found_anagram = true
				}
				index++
			}
			if !found_anagram {
				grouped_anagrams = append(grouped_anagrams, []string{str})
			}
		}
	}
	return grouped_anagrams
}
