package main

import "fmt"

/* Problem 896: A range of integers is a range of numbers that is divisible by a given number.

- A contiguous range of positive integers is called a divisible range if all the integers in the range can be arranged in a row such that the n-th
term is a multiple of of n.

	- For example, the range [6..9] is a divisible range because we can arrange the numbers as [7, 6, 9, 8]

	- In fact, it is the 4th divisible range of length of 4, the first three being [1..4], [2..5], and [3..6].

Expected Output: Find the 36th divisible range of length 36, and return the smallest number in the range
Extra Credit: Keep track of the range as you progress through the loop, return the range when it is finished along side the smallest number
*/

/* getDesiredDivisibleRange prompts the user to enter the n-th divisible range to find and returns the entered value.
//
// This function does the following:
//   1. Prints a message to the console asking the user to enter the n-th divisible range to find.
//   2. Waits for the user to enter a value and stores it in the rangeInput variable.
//   3. Returns the value of rangeInput.
//
// Returns an integer representing the entered value.
*/
func getDesiredDivisibleRange() int {
	var rangeInput int
	fmt.Println("Enter N (the n-th divisible range to find):")
	fmt.Scan(&rangeInput)
	return rangeInput 
}

/* getDesiredLength prompts the user to enter the length of the desired divisible range and returns the entered value.
//
// This function does the following:
//   1. Prints a message to the console asking the user to enter the length of the divisible range to find.
//   2. Waits for the user to enter a value and stores it in the length variable.
//   3. Returns the value of length.
//
// Returns an integer representing the entered value.
*/
func getDesiredLength() int {
	var length int
	// Prompt for Length Input
	fmt.Println("Enter L (length of the divisible range):")
	// Await for the response from user
	fmt.Scan(&length)
	return length
}

func isInDivisibleRangeWereLookingFor(smallestNumber int, desiredRangeLength int) bool {
	return true
}

// findSmallestDivisibleRange finds the smallest number in a given range that forms a divisible range of a specified length.
//
// Parameters:
// - desiredRangeInput: an integer representing the range of numbers to search within.
// - desiredRangeLength: an integer representing the length of the desired divisible range.
//
// Returns:
// - an integer representing the smallest number in the range that forms a divisible range of the specified length.
func findSmallestDivisibleRange(desiredRange int, desiredLength int) int {
	// Defines the smallest number in a divisible range, this is what we'll return. 
	smallestNumber := 1

	// Define a slice that will represent the current range (extra credit)

	// Define a counter that keeps track of how many multiple numbers have been found in the loop.
	// When the count reaches the desired rangeLength, return the current smallestNumber
	count := 0

	for {
		// Determine if the next smallestNumber is part of of the range we're looking for.
		if isInDivisibleRangeWereLookingFor(smallestNumber, desiredLength) {
			// If so, add to the current count
			count++

			// Add the current smallestNumber to the slice (extra credit)

			// If the count has reached the desiredRangeInput, return the current smallestNumber (TODO: Use if, ok to shorten the declaration of currentRange, and return its state too)
			if count == desiredRange {
				return smallestNumber;
				// return smallestNumber, currentRange
			}
		}

		// If current smallest number isn't in the divisible range we're looking for, keep going
		smallestNumber++
	}
}

func main() {
	// Get Range & Length from User
	desiredRange := getDesiredDivisibleRange()
	desiredLength := getDesiredLength()

	// Find the divisible range and return its smallest number
	result := findSmallestDivisibleRange(desiredRange, desiredLength)

	// Print the smallest number in the range as the desired answer.
	fmt.Printf("The smallest number in the %d-th divisible range of length %d is %d\n", desiredRange, desiredLength, result)
}

