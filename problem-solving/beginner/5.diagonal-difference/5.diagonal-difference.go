/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

package main

import "log"

func diagonalDifference(arr [][]int32) int32 {
	var primaryDiagonal int32 = 0
	var secondDiagonal int32 = 0
	n := len(arr)

	// Write your code here
	for i := 0; i < n; i++ {
		primaryDiagonal += arr[i][i]
		secondDiagonal += arr[i][n-1-i]
	}

	diff := primaryDiagonal - secondDiagonal
	if diff < 0 {
		return -diff
	}

	return diff
}

func main() {
	arr := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	result := diagonalDifference(arr)

	log.Println("The result of diagonal difference", result)
}
