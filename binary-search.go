package main

// Perform binary search.
// Return the target's location in the slice and the number of tests.
// If the item is not found, return -1 and the number tests.
func binarySearch(slice []int, target int) (index, numTests int) {
	// Set the initial low and high bounds.
	low := 0
	high := len(slice) - 1

	// Loop until the bounds cross.
	for low <= high {
		// Increment the number of tests.
		numTests++

		// Calculate the middle index.
		mid := (low + high) / 2

		// If the target is less than the middle item, set the new high.
		if target < slice[mid] {
			high = mid - 1

			// If the target is greater than the middle item, set the new low.
		} else if target > slice[mid] {
			low = mid + 1

			// Otherwise we found it. Return the index and number of tests.
		} else {
			return mid, numTests
		}
	}

	// Return -1 and the number of tests.
	return -1, numTests
}
