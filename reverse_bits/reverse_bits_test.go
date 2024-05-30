package reverse_bits

import (
	"testing"
)

func TestReverseBits(t *testing.T) {
	testcases := []struct {
		num, solution uint32
	}{
		/* Input 00000010100101000001111010011100 represents 43261596 so return
		964176192 which is 00111001011110000010100101000000. */
		{ 43261596, 964176192 },
		/* The input 11111111111111111111111111111101 represents 4294967293, so
		return 3221225471 which is 10111111111111111111111111111111. */
		{ 3221225471, 4294967293 },
	}

	for _, testcase := range testcases {
		result := ReverseBits( testcase.num )

		if result != testcase.solution {
			t.Errorf(
				"ReverseBits: %v returned %v, want %v",
				testcase.num,
				result,
				testcase.solution,
			)
		}
	}
}
