/*
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

package main

import "log"

func aVeryBigSum(ar []int64) int64 {
	var total int64 = 0

	for _, item := range ar {
		total += item
	}

	return total
}

func main() {
	var arr = []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	var result = aVeryBigSum(arr)

	log.Println("Result of a very big sum", result)
}
