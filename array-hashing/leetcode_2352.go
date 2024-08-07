package main

func main() {

}

// Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
// Output: 1
// Explanation: There is 1 equal row and column pair:
// - (Row 2, Column 1): [2,7,7]Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
// Output: 1
// Explanation: There is 1 equal row and column pair:
// - (Row 2, Column 1): [2,7,7]
func equalPairs(grid [][]int) int {
	n := len(grid)
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			k := 0
			for ; k < n; k++ {
				if grid[i][k] != grid[k][j] {
					break
				}
			}
			if k == n{
				count++
			}
		}
	}
	return count

}
