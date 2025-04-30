package main

import "log"

func simpleArraySum(ar []int32) int32 {
	var sum int32 = 0

	for _, value := range ar {
		sum += value
	}

	return sum
}

func main() {
	ar := []int32{1, 2, 3, 4, 5, 6}
	var sum = simpleArraySum(ar)
	log.Println(sum)
}
