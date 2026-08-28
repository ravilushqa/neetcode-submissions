func numIslands(grid [][]byte) int {
    if grid == nil {
        return 0
    }

    directions := [][]int{{1,0}, {-1,0}, {0, 1}, {0, -1}}
    rows, cols := len(grid), len(grid[0])
    visited := make(map[[2]int]bool, rows*cols)
    islands := 0

    var bfs func(r, c int)

    bfs = func(r,c int) {
        queue := [][2]int{{r,c}}

        for len(queue) > 0 {
            cell := queue[0]
            queue = queue[1:]

            row, col := cell[0], cell[1]
            for _, dir := range directions {
                nr, nc := row + dir[0], col + dir[1]

                if nr >= 0 && nc >= 0 && nr < rows && nc < cols && grid[nr][nc] == '1' {
                    queue = append(queue, [2]int{nr, nc})
                    grid[nr][nc] = '0'
                }
            }
        }

    }


    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == '1' && !visited[[2]int{r,c}] {
                bfs(r,c)
                islands++
            }
        }
    }

    return islands
}