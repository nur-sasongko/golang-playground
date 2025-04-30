// Goals: Implement a function to add two numbers and print the result

package main

// uint32 is an unsigned 32-bit integer, which is a non-negative integer

// solveMeFirst takes two uint32 numbers and returns their sum
func solveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

func main() {
	const a, b = 2, 3
	result := solveMeFirst(a, b)
	println(result)
}
