package simplify_path

const FORWARD_SLASH rune = '/'
const DOT rune = '.'
const PREVIOUS_DIR int = 2

/* Assumes that:
Any sequence of periods longer than two is a valid name for a directory.
1 <= path.length <= 3000.
path consists of English letters, digits, period '.', slash '/' or '_'.
path is a valid absolute Unix path.*/
func SimplifyPath(path string) string {

	return_path := []rune{}
	path_runes := []rune( path )
	dot_count := 0
	for _, rune := range path_runes {
		if rune == FORWARD_SLASH {
			// Handle dot(s)
			if dot_count == PREVIOUS_DIR {
				// Remove previous dir from stack
				slashes_encountered := 0
				index := len(return_path) - 1
				for slashes_encountered < 2 && index >= 0 {
					if return_path[ index ] == FORWARD_SLASH {
						slashes_encountered++
					}
					index--
				}
				if index >= 0 {
					// Avoid removing second slash and rune before it
					index += 2
					return_path = return_path[:index]
				}
			} else if dot_count > PREVIOUS_DIR {
				for index := 0; index < dot_count; index++ {
					return_path = append( return_path, DOT )
				}
			}
			// Reset count of dots
			dot_count = 0
			// Avoid repeated slashes
			if len( return_path ) < 1 || return_path[ len(return_path)-1 ] != FORWARD_SLASH {
				return_path = append( return_path, rune )
			}
		} else if rune == DOT {
			dot_count++
		} else {
			return_path = append( return_path, rune )
		}
	}

	// Remove trailing slash if not root
	last_index := len(return_path) - 1
	if len( return_path ) > 1 && return_path[ last_index ] == FORWARD_SLASH {
		return_path = return_path[:last_index]
	}

	return string( return_path )
}
