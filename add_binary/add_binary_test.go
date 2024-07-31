package add_binary

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	testcases := []struct {
		a, b, solution string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"0", "0", "0"},
	}

	for _, testcase := range testcases {
		result := AddBinary(testcase.a, testcase.b)

		if result != testcase.solution {
			t.Errorf(
				"AddBinary: %v and %v returned %v, want %v",
				testcase.a,
				testcase.b,
				result,
				testcase.solution,
			)
		}
	}
}
