/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/reverse-words-in-a-string/description/
*/
package reverse_words_in_a_string

/*
ReverseWords reverses the order of the words in s and returns the resulting
string.
ReverseWords assumes that:
1 <= s.length <= 104,
s contains English letters (upper-case and lower-case), digits, and spaces ' ',
a word is defined as a sequence of non-space characters,
the words in s will be separated by at least one space,
and there is at least one word in s.
*/
func ReverseWords(s string) string {
	rune_s := []rune(s)
	return_runes := make([]rune, 0)

	found_word := false
	for index := len(rune_s) - 1; index >= 0; index-- {
		if !found_word && rune_s[index] != ' ' {
			found_word = true
		} else if found_word && rune_s[index] == ' ' {
			read_index := index + 1
			for read_index < len(rune_s) && rune_s[read_index] != ' ' {
				return_runes = append(return_runes, rune_s[read_index])
				read_index++
			}
			return_runes = append(return_runes, ' ')
			found_word = false
		}
	}

	if found_word {
		read_index := 0
		for read_index < len(rune_s) && rune_s[read_index] != ' ' {
			return_runes = append(return_runes, rune_s[read_index])
			read_index++
		}
		return_runes = append(return_runes, ' ')
	}

	// Remove trailing space
	return_runes = return_runes[:(len(return_runes) - 1)]
	return string(return_runes)
}
