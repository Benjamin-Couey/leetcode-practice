package majority_element

import (
		"testing"
)

func TestMerge(t *testing.T) {
		testcases := []struct {
				nums1 []int
				solution int
				error bool
		}{
				{ []int{ 3, 2, 3 }, 3, false },
				{ []int{ 2, 2, 1, 1, 1, 2, 2 }, 2, false },
				{ []int{ 1 }, 1, false },
				{ []int{}, 0, true },
				{ []int{ 1, 2, 3 }, 0, true },
		}

		for _, testcase := range testcases {
				answer, error := MajorityElement( testcase.nums1 )

				if ( error != nil ) != testcase.error {
					var message string
					if testcase.error {
						message = "error message"
					} else {
						message = "no error message"
					}
					t.Errorf("MajorityElement: %v, want %v", error, message)
				}

				if answer != testcase.solution {
						t.Errorf("MajorityElement: %v, want %v", answer, testcase.solution)
				}
		}
}
