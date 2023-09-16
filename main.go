package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Make and sort an slice.
	var numItems, maxIt int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&maxIt)
	slice := makeRandomSlice(numItems, maxIt)
	printSlice(slice, 40)

	quicksort(slice)
	fmt.Println("sorted:\n", slice)

	for {
		// Get the target as a string.
		var targetString string
		fmt.Printf("Target: ")
		fmt.Scanln(&targetString)

		// If the target string is blank, break out of the loop.
		if len(targetString) == 0 {
			break
		}

		// Convert to int and find it.
		target, _ := strconv.Atoi(targetString)

		var index, numTests int
		if len(slice) < 4 {
			index, numTests = linearSearch(slice, target)
		} else {
			index, numTests = binarySearch(slice, target)
		}

		if index < 0 || index >= len(slice) {
			fmt.Printf("Target %d not found, %d tests\n", target, numTests)
		} else {
			fmt.Printf("slice[%d] = %d, %d tests\n", index, slice[index], numTests)
		}
	}
}
