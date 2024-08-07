package main

func main() {
	board := [][]string{
		{"1", "2", ".", ".", "3", ".", ".", ".", "."},
		{"4", ".", ".", "5", ".", ".", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", ".", "3"},
		{"5", ".", ".", ".", "6", ".", ".", ".", "4"},
		{".", ".", ".", "8", ".", "3", ".", ".", "5"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", ".", ".", ".", ".", ".", "2", ".", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "8"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
}
func isValidSudoku(board [][]byte) bool {
	check := make(map[string]bool)

	for rowIndex, row := range board {
		for columnIndex, value := range row {
			if value == '.' {
				continue
			}

			valueInCol := string(value) + "@col" + string(columnIndex)
			valueInRow := string(value) + "@row" + string(rowIndex)
			valueInBox := string(value) + "@box" + string(rowIndex/3) + string(columnIndex/3)

			if check[valueInCol] || check[valueInRow] || check[valueInBox] {
				return false
			}

			check[valueInCol] = true
			check[valueInRow] = true
			check[valueInBox] = true
		}
	}

	return true
}
