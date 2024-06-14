package game_of_life

/* Assumes that:
Any live cell with fewer than two live neighbors dies as if caused by under-population.
Any live cell with two or three live neighbors lives on to the next generation.
Any live cell with more than three live neighbors dies, as if by over-population.
Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
These rule are applied simultaneously to every cell in a given state.
m == board.length
n == board[i].length
1 <= m, n <= 25
board[i][j] is 0 or 1.
Cells outside of the bounds of the board are dead. */
const DEAD int = 0
const ALIVE int = 1
const DIE int = 2
const SURVIVE int = 3
const STAY_DEAD int = 4
const REPRODUCE int = 5

func GameOfLife(board [][]int)  {

	for row_index := 0; row_index < len(board); row_index++ {
		for column_index := 0; column_index < len(board[0]); column_index++ {
			live_neighbors := 0
			for adjacent_row := row_index - 1; adjacent_row <= row_index + 1; adjacent_row++ {
				for adjacent_column := column_index - 1; adjacent_column <= column_index + 1; adjacent_column++ {
					row_in_bounds := adjacent_row >= 0 && adjacent_row < len(board)
					col_in_bounds := adjacent_column >= 0 && adjacent_column < len(board[0])
					not_center := !( adjacent_row == row_index && adjacent_column == column_index )
					if row_in_bounds && col_in_bounds && not_center {
						neighbor := board[ adjacent_row ][ adjacent_column ]
						if neighbor == ALIVE || neighbor == DIE || neighbor == SURVIVE {
							live_neighbors++
						}
					}
				}
			}
			alive := board[ row_index ][ column_index ] == ALIVE
			if alive {
				if live_neighbors < 2 || live_neighbors > 3 {
					board[ row_index ][ column_index ] = DIE
				} else {
					board[ row_index ][ column_index ] = SURVIVE
				}
			} else {
				if live_neighbors == 3 {
					board[ row_index ][ column_index ] = REPRODUCE
				} else {
					board[ row_index ][ column_index ] = STAY_DEAD
				}
			}
		}
	}

	for row_index := 0; row_index < len(board); row_index++ {
		for column_index := 0; column_index < len(board[0]); column_index++ {
			cell := board[ row_index ][ column_index ]
			if cell == DIE || cell == STAY_DEAD {
				board[ row_index ][ column_index ] = DEAD
			} else if cell == SURVIVE || cell == REPRODUCE {
				board[ row_index ][ column_index ] = ALIVE
			}
		}
	}

}
