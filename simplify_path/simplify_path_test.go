package simplify_path

import (
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	testcases := []struct {
		path     string
		solution string
	}{
		{"/home/", "/home"},
		{"/home//foo/", "/home/foo"},
		{"/home/user/Documents/../Pictures", "/home/user/Pictures"},
		{"/.../a/../b/c/../d/./", "/.../b/d"},
		{"/../", "/"},
	}

	for _, testcase := range testcases {
		result := SimplifyPath(testcase.path)

		if result != testcase.solution {
			t.Errorf(
				"SimplifyPath: %v returned %v, want %v",
				testcase.path,
				result,
				testcase.solution,
			)
		}
	}
}
