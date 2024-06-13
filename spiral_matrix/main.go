package spiral_matrix


const RIGHT int = 0
const DOWN int = 1
const LEFT int = 2
const UP int = 3

/* Assumes that:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100 */
func SpiralOrder(matrix [][]int) []int {
	spiral_order := make( []int, 0 )

	left_bound := -1
	right_bound := len( matrix[0] )
	up_bound := -1
	down_bound := len( matrix )

	row_index := 0
	column_index := 0
	direction := RIGHT
	done := false

	for !done {
		switch direction {
			case RIGHT:
				for column_index < right_bound {
					spiral_order = append( spiral_order, matrix[row_index][column_index] )
					column_index++
				}

				if down_bound <= row_index + 1 {
					done = true
				} else {
					row_index++
					column_index--
					up_bound++
					direction = DOWN
				}

			case DOWN:
				for row_index < down_bound {
					spiral_order = append( spiral_order, matrix[row_index][column_index] )
					row_index++
				}

				if left_bound >= column_index - 1 {
					done = true
				} else {
					column_index--
					row_index--
					right_bound--
					direction = LEFT
				}

			case LEFT:
				for column_index > left_bound {
					spiral_order = append( spiral_order, matrix[row_index][column_index] )
					column_index--
				}

				if up_bound >= row_index - 1 {
					done = true
				} else {
					row_index--
					column_index++
					down_bound--
					direction = UP
				}

			case UP:
				for row_index > up_bound {
					spiral_order = append( spiral_order, matrix[row_index][column_index] )
					row_index--
				}

				if right_bound <= column_index + 1 {
					done = true
				} else {
					column_index++
					row_index++
					left_bound++
					direction = RIGHT
				}
		}
	}
	return spiral_order
}
