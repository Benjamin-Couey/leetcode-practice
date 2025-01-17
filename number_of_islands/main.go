/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/number-of-islands/description/
*/
package number_of_islands

/*
NumIslands returns the number of islands in grid. An island is a number of '1'
bytes (land) connected horizontally and verticlaly and surrounded by '0' bytes
(water).
NumIslands assumes that:
all four edges of grid are surrounded by water,
m == grid.length,
n == grid[i].length,
1 <= m, n <= 300,
and grid[i][j] is '0' or '1'. */
func NumIslands(grid [][]byte) int {

	islands := 0
	encountered_cells := make(map[[2]int]bool)

	for row_index, row := range grid {
		for col_index, cell := range row {
			_, exists := encountered_cells[[2]int{row_index, col_index}]

			if !exists && cell == byte('1') {
				// Recursively encounter all cells in the new island.
				encounterIsland(grid, row_index, col_index, encountered_cells)
				islands++
			}
		}
	}

	return islands
}

func encounterIsland(grid [][]byte, row, col int, encountered_cells map[[2]int]bool) {
	position := [2]int{row, col}
	_, exists := encountered_cells[position]
	if !exists && grid[row][col] == byte('1') {
		encountered_cells[position] = true
		// Up
		if row-1 >= 0 {
			encounterIsland(grid, row-1, col, encountered_cells)
		}
		// Down
		if row+1 < len(grid) {
			encounterIsland(grid, row+1, col, encountered_cells)
		}
		// Left
		if col-1 >= 0 {
			encounterIsland(grid, row, col-1, encountered_cells)
		}
		// Right
		if col+1 < len(grid[0]) {
			encounterIsland(grid, row, col+1, encountered_cells)
		}
	}
}
