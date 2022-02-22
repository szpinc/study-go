package main

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	findNumberIn2DArray(matrix, 5)
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 行数
	rows := len(matrix)

	if rows == 0 {
		return false
	}
	// 列数
	cols := len(matrix[0])

	i, j := rows-1, 0

	for i > 0 && j < cols {
		value := matrix[i][j]
		if value > target {
			i--
		} else if value < target {
			j++
		} else {
			return true
		}
	}
	return false
}
