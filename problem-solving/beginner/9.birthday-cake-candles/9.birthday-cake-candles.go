package main

import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	var maxHeight int32 = 0
	var count int32 = 0

	for _, height := range candles {
		if height > maxHeight {
			maxHeight = height
			count = 1
		} else if height == maxHeight {
			count++
		}
	}

	return count
}

func main() {
	candles := []int32{3, 2, 1, 3}
	result := birthdayCakeCandles(candles)
	fmt.Println(result)
}
