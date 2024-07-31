package set_matrix_zeroes

/* Assumes that:
m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-2^31 <= matrix[i][j] <= 2^31 - 1*/
func SetZeroes(matrix [][]int) {
	/* Flag rows and columns that will be set to 0 by setting first element of
	row or column to 0. */
	for row_index, _ := range matrix {
		for column_index, val := range matrix[row_index] {
			if val == 0 {
				matrix[0][column_index] = 0
				matrix[row_index][0] = 0
			}
		}
	}

	// Set non-first elements to 0 based on flags.
	for row_index := 1; row_index < len(matrix); row_index++ {
		for column_index := 1; column_index < len(matrix[0]); column_index++ {
			if matrix[0][column_index] == 0 || matrix[row_index][0] == 0 {
				matrix[row_index][column_index] = 0
			}
		}
	}

	// Handle first elements in rows and columns.
	if matrix[0][0] == 0 {
		for row_index := 0; row_index < len(matrix); row_index++ {
			matrix[row_index][0] = 0
		}

		for column_index := 0; column_index < len(matrix[0]); column_index++ {
			matrix[0][column_index] = 0
		}
	}
}
