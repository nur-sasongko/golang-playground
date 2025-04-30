/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

package main

import "log"

func compareTriplets(a []int32, b []int32) []int32 {
	var aPoints, bPoints int32 = 0, 0

	for i := range a {
		if a[i] > b[i] {
			aPoints++
		} else if b[i] > a[i] {
			bPoints++
		}
	}

	return []int32{aPoints, bPoints}
}

func main() {
	var alicePoints = []int32{17, 28, 30}
	var bobPoints = []int32{99, 16, 8}

	var result = compareTriplets(alicePoints, bobPoints)
	log.Println(`The result: `, result)
}
