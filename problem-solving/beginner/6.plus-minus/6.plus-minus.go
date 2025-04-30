/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

package main

import "fmt"

func plusMinus(arr []int32) {
	var minusCount, plusCount, zeroCount int32 = 0, 0, 0
	var ratios = [3]float64{0, 0, 0}

	for _, item := range arr {
		if item == 0 {
			zeroCount++
		} else if item > 0 {
			plusCount++
		} else if item < 0 {
			minusCount++
		}
	}

	n := float64(len(arr)) // convert to float64 to get decimal result
	ratios[0] = float64(plusCount) / n
	ratios[1] = float64(minusCount) / n
	ratios[2] = float64(zeroCount) / n

	for _, ratio := range ratios {
		fmt.Printf("%.6f\n", ratio)
	}
}

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}
