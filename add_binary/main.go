package add_binary

/* Assumes that:
1 <= a.length, b.length <= 10^4
a and b consist only of '0' or '1' characters.
Each string does not contain leading zeros except for the zero itself. */
func AddBinary(a string, b string) string {
	rune_a := []rune(a)
	rune_b := []rune(b)
	shorter := rune_a
	longer := rune_b
	if len(rune_b) < len(rune_a) {
		shorter = rune_b
		longer = rune_a
	}

	shorter_index := len(shorter) - 1
	longer_index := len(longer) - 1
	carry := false
	for shorter_index >= 0 {
		ones := 0
		if carry {
			ones++
		}
		if shorter[shorter_index] == '1' {
			ones++
		}
		if longer[longer_index] == '1' {
			ones++
		}
		switch ones {
		case 1:
			longer[longer_index] = '1'
			carry = false
		case 2:
			longer[longer_index] = '0'
			carry = true
		case 3:
			longer[longer_index] = '1'
			carry = true
		default:
			longer[longer_index] = '0'
			carry = false
		}

		shorter_index--
		longer_index--
	}

	for longer_index >= 0 && carry {
		if longer[longer_index] == '0' {
			longer[longer_index] = '1'
			carry = false
		} else {
			longer[longer_index] = '0'
			carry = true
		}
		longer_index--
	}

	if carry {
		longer = append([]rune{'1'}, longer...)
	}

	return string(longer)
}
