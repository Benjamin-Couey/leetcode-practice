package ransom_note

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
