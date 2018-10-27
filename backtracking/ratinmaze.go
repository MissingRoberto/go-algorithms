package backtracking

type Grid [][]int

func (g Grid) IsValid(i, j int) bool {
	return i < len(g) && i >= 0 && j >= 0 && j < len(g[0]) && g[i][j] == 1
}

func resolveMaze(originX, originY, destX, destY int, grid, solution Grid) bool {
	if !grid.IsValid(originX, originY) {
		return false
	}

	solution[originX][originY] = 1
	if originX == destX && originY == destY {
		return true
	}

	if resolveMaze(originX+1, originY, destX, destY, grid, solution) {
		return true
	}

	if resolveMaze(originX, originY+1, destX, destY, grid, solution) {
		return true
	}
	solution[originX][originY] = 0
	return false
}

func ResolveMaze(originX, originY, destX, destY int, grid [][]int) ([][]int, bool) {
	result := make([][]int, 0)
	for i := 0; i < len(grid); i++ {
		result = append(result, make([]int, len(grid[0])))
	}

	isResolved := resolveMaze(originX, originY, destX, destY, Grid(grid), Grid(result))
	return result, isResolved
}
