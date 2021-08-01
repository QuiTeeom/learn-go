package main

import "fmt"

func main() {
	rotate([][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	})
}
func rotate(matrix [][]int) {
	n := len(matrix)
	var x, y int

	var t int

	for x < n {
		for y < n {
			t = matrix[x][y]
			matrix[x][y] = matrix[n-x-1][y]
			matrix[n-x-1][y] = matrix[n-x-1][n-y-1]
			matrix[n-x-1][n-y-1] = matrix[x][n-y-1]
			matrix[x][n-y-1] = t
			y++
			fmt.Println(matrix)
		}
		x++
	}

}
