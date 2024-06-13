package main

import (
	"fmt"
)

func main() {
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	matrix := make([][]int, n)
	maxIndex := 0
	maxScore := 0

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			_, _ = fmt.Scan(&matrix[i][j])
		}
	}
	for i := 0; i < n; i++ {
		score := 0
		for j := 0; j < m; j++ {
			score += matrix[i][j]
		}
		if maxScore < score {
			maxScore = score
			maxIndex = i
		}
	}
	fmt.Println(maxScore)
	fmt.Println(maxIndex)
}
