/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	nums := make([]int, len(arr))
	for i, v := range arr {
		nums[i] = int(v)
	}

	sort.Ints(nums) // Sort in ascending order

	// Calculate sum of first 4 and last 4
	var minSum, maxSum int64 = 0, 0
	for i := 0; i < len(nums)-1; i++ {
		minSum += int64(nums[i])
		maxSum += int64(nums[i+1])
	}

	fmt.Printf("%d %d\n", minSum, maxSum)
}

func main() {
	arr := []int32{7, 69, 2, 221, 8974}
	miniMaxSum(arr)
}
