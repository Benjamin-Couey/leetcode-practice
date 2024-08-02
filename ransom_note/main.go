/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/ransom-note/description/
*/
package ransom_note

/*
CanConstruct reports whether ransomNote can be constructed using letters from
magazine. Each letter from magazine can only be used once.
CanConstruct assumes that:
1 <= ransomNote.length, magazine.length <= 10^5,
and ransomNote and magazine consist of lowercase English letters.
*/
func CanConstruct(ransomNote string, magazine string) bool {

	rune_counts := make(map[rune]int)

	for _, rune := range magazine {
		_, exists := rune_counts[rune]
		if exists {
			rune_counts[rune] += 1
		} else {
			rune_counts[rune] = 1
		}
	}

	for _, rune := range ransomNote {
		_, exists := rune_counts[rune]
		if exists {
			rune_counts[rune] -= 1
			if rune_counts[rune] < 0 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
