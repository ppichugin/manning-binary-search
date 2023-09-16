package main

import (
	"fmt"
	"math/rand"
	"time"
)

// makeRandomSlice makes a slice containing pseudorandom numbers in [0, max).
func makeRandomSlice(numItems, max int) []int {
	// Initialize a pseudorandom number generator.
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

	slice := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		slice[i] = random.Intn(max)
	}
	return slice
}

// printSlice prints at most numItems items.
func printSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}
