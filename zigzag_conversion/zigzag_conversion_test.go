package zigzag_conversion

import (
	"testing"
)

func TestConvert(t *testing.T) {
	testcases := []struct {
		s string
		numRows int
		solution string
	}{
		{ "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR" },
		{ "PAYPALISHIRING", 4, "PINALSIGYAHRPI" },
		{ "A", 1, "A" },
	}

	for _, testcase := range testcases {
		result := Convert( testcase.s, testcase.numRows )

		if result != testcase.solution {
			t.Errorf(
				"Convert: %v and %v returned %v, want %v",
				testcase.s,
				testcase.numRows,
				result,
				testcase.solution,
			)
		}
	}
}
