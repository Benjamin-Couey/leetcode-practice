/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/set-matrix-zeroes/description/
*/
package simplify_path

// Names given to specific characters used in the path.
const FORWARD_SLASH rune = '/'
const DOT rune = '.'
// The number of dots which indicate the previous directory.
const PREVIOUS_DIR int = 2

/*
SimplifyPath returns the simplified canonical path based on path. The simplified
path will follow these rules:
it will start with a single '/',
directories in the path are seperated by a single '/',
it will not end with a '/' unless it is the root directory,
and it excludes any '.' or '..' used to denote the current or parent directory.
SimplifyPath assumes that:
any sequence of periods longer than two is a valid name for a directory,
1 <= path.length <= 3000,
path consists of English letters, digits, period '.', slash '/' or '_',
and path is a valid absolute Unix path.
*/
func SimplifyPath(path string) string {

	return_path := []rune{}
	path_runes := []rune(path)
	dot_count := 0
	for _, rune := range path_runes {
		if rune == FORWARD_SLASH {
			// Handle dot(s).
			if dot_count == PREVIOUS_DIR {
				// Remove previous dir from stack.
				slashes_encountered := 0
				index := len(return_path) - 1
				for slashes_encountered < 2 && index >= 0 {
					if return_path[index] == FORWARD_SLASH {
						slashes_encountered++
					}
					index--
				}
				if index >= 0 {
					// Avoid removing second slash and rune before it.
					index += 2
					return_path = return_path[:index]
				}
			} else if dot_count > PREVIOUS_DIR {
				for index := 0; index < dot_count; index++ {
					return_path = append(return_path, DOT)
				}
			}
			// Reset count of dots.
			dot_count = 0
			// Avoid repeated slashes.
			if len(return_path) < 1 || return_path[len(return_path)-1] != FORWARD_SLASH {
				return_path = append(return_path, rune)
			}
		} else if rune == DOT {
			dot_count++
		} else {
			return_path = append(return_path, rune)
		}
	}

	// Remove trailing slash if not root.
	last_index := len(return_path) - 1
	if len(return_path) > 1 && return_path[last_index] == FORWARD_SLASH {
		return_path = return_path[:last_index]
	}

	return string(return_path)
}
