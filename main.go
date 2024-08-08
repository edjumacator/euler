package main

import "fmt"

/* Problem 896: A range of integers is a range of numbers that is divisible by a given number.

- A contiguous range of positive integers is called a divisible range if all the integers in the range can be arranged in a row such that the n-th
term is a multiple of of n.

	- For example, the range [6..9] is a divisible range because we can arrange the numbers as [7, 6, 9, 8]
	is a divisible range because we can arrange the numbers as

	- In fact, it is the 4th divisible range of length of 4, the first three being [1..4], [2..5], and [3..6].
*/

func findSmallestDivisibleRange(k int, rangeLength int) int {
	smallestNumber := 0
	return smallestNumber
}

func getDesiredDivisibleRange() int {
	var rangeInput int

	// Prompt for Range Input
	fmt.Println("Enter N (the n-th divisible range to find):")
	// Await for the response from user 
	fmt.Scan(&rangeInput)

	return rangeInput 
}

func getDesiredLength() int {
	var length int
	// Prompt for Length Input
	fmt.Println("Enter L (length of the divisible range):")
	// Await for the response from user
	fmt.Scan(&length)
	return length
}

func main() {
	// Get Range & Length from User
	desiredRange := getDesiredDivisibleRange()
	desiredLength := getDesiredLength()

	result := findSmallestDivisibleRange(desiredRange, desiredLength)

	// Print the smallest number in the range as the desired answer.
	fmt.Printf("The smallest number in the %d-th divisible range of length %d is %d\n", desiredRange, desiredLength, result)
}

