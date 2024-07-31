package word_pattern

import (
	"strings"
)

/* Assumes that:
1 <= pattern.length <= 300
pattern contains only lower-case English letters.
1 <= s.length <= 3000
s contains only lowercase English letters and spaces ' '.
s does not contain any leading or trailing spaces.
All the words in s are separated by a single space. */
func WordPattern(pattern string, s string) bool {

	words := strings.Split(s, " ")
	rune_to_word_map := make(map[rune]string)
	rune_pattern := []rune(pattern)

	if len(rune_pattern) != len(words) {
		return false
	}

	for index := 0; index < len(rune_pattern); index++ {
		rune_word, exists := rune_to_word_map[rune_pattern[index]]

		if exists {
			if rune_word != words[index] {
				return false
			}
		} else {
			rune_to_word_map[rune_pattern[index]] = words[index]
		}
	}

	return true
}
