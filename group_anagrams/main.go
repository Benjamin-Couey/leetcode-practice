package group_anagrams

import (
	"leetcode/valid_anagram"
)

/* Assumes that:
The answer may be returned in any order.
1 <= strs.length <= 10^4
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters. */
func GroupAnagrams(strs []string) [][]string {
	grouped_anagrams := make( [][]string, 0 )

	for _, str := range strs {
		if len(grouped_anagrams) < 1 {
			grouped_anagrams = append( grouped_anagrams, []string{ str } )
		} else {
			index := 0
			found_anagram := false
			for !found_anagram && index < len( grouped_anagrams ) {
				anagrams := grouped_anagrams[ index ]
				if valid_anagram.IsAnagram( str, anagrams[0] ) {
					anagrams = append( anagrams, str )
					grouped_anagrams[ index ] = anagrams
					found_anagram = true
				}
				index++
			}
			if !found_anagram {
				grouped_anagrams = append( grouped_anagrams, []string{ str } )
			}
		}
	}
	return grouped_anagrams
}
