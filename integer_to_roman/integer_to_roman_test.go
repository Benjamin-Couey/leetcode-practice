package integer_to_roman

import (
	"testing"
)

func TestIntToRoman(t *testing.T) {
	testcases := []struct {
		num      int
		solution string
	}{
		{3749, "MMMDCCXLIX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, testcase := range testcases {
		value := IntToRoman(testcase.num)

		if value != testcase.solution {
			t.Errorf("TestIntToRoman: %v, want %v", value, testcase.solution)
		}

	}
}
