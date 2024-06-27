package surrounded_regions

/* Assumes that:
A cell is connected to adjacent cells horizontally or vertically.
Every connected 'O' cell forms a region.
A region is surrounded with 'X' cells if the region is connected with 'X' cells
and none of the region cells are on the edge of the board.
m == board.length.
n == board[i].length.
1 <= m, n <= 200.
board[i][j] is 'X' or 'O'. */
func Solve(board [][]byte)  {

	if len(board) < 3 || len(board[0]) < 3 {
		return
	}

	cell_is_free := make( map[[2]int]bool )

	last_row := len(board) - 1
	last_col := len(board[0]) - 1

	// Find all cells in regions which touch the border and thus should not be
	// captured.
	for row_index := 0; row_index < len(board); row_index += last_row  {
		for col_index := 0; col_index < len(board[0]); col_index++ {
			cell := board[ row_index ][ col_index ]
			_, exists := cell_is_free[ [2]int{ row_index, col_index } ]
			if !exists && cell == byte('O') {
				solveRegion( board, row_index, col_index, cell_is_free )
			}
		}
	}

	for col_index := 0; col_index < len(board[0]); col_index += last_col {
		for row_index := 0; row_index < len(board); row_index++ {
			cell := board[ row_index ][ col_index ]
			_, exists := cell_is_free[ [2]int{ row_index, col_index } ]
			if !exists && cell == byte('O') {
				solveRegion( board, row_index, col_index, cell_is_free )
			}
		}
	}

	// Capture all cells not on the boarder and not in a region that touches the
	// boarder.
	for row_index := 1; row_index < last_row; row_index++ {
		for col_index := 1; col_index < last_col; col_index++ {
			cell := board[ row_index ][ col_index ]
			_, exists := cell_is_free[ [2]int{ row_index, col_index } ]
			if !exists && cell == byte('O') {
				board[ row_index ][ col_index ] = byte('X')
			}
		}
	}

}

func solveRegion( board [][]byte, row, col int, cell_is_free map[[2]int]bool ) {
	position := [2]int{ row, col }
	_, exists := cell_is_free[ position ]
	if !exists && board[row][col] == byte('O') {
		cell_is_free[ position ] = true
		// Up
		if row - 1 >= 0 {
			solveRegion( board, row - 1, col, cell_is_free )
		}
		// Down
		if row + 1 < len( board ) {
			solveRegion( board, row + 1, col, cell_is_free )
		}
		// Left
		if col - 1 >= 0 {
			solveRegion( board, row, col - 1, cell_is_free )
		}
		// Right
		if col + 1 < len( board[0] ) {
			solveRegion( board, row, col + 1, cell_is_free )
		}
	}
}
