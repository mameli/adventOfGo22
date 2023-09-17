package calories

import (
	"fmt"
	"go_advent/utility"
	"sort"
	"strconv"
)

func Count_cal() {
	lines := utility.Read_text_file("./inputs/1_calories.txt")

	// Create a vector of vectors to store the numbers
	var vectorOfVectors [][]int
	var currentVector []int

	// loop through the lines
	for _, line := range lines {
		// if the line is empty, append the current vector to the vector of vectors
		if line == "" {
			vectorOfVectors = append(vectorOfVectors, currentVector)
			currentVector = []int{}
			continue
		}

		// convert the string to int and append it to the current vector
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}
		currentVector = append(currentVector, num)
	}

	// Check for any remaining numbers after the last blank line
	if len(currentVector) > 0 {
		vectorOfVectors = append(vectorOfVectors, currentVector)
	}

	fmt.Println(vectorOfVectors)

	// loop through the vector of vectors and sum the numbers
	var maxSum int
	for _, vector := range vectorOfVectors {
		sum := 0
		for _, num := range vector {
			sum += num
		}
		if sum > maxSum {
			maxSum = sum
		}

	}

	// print the max sum
	fmt.Println("Max sum:", maxSum)

	// array of length len(vectorOfVectors) to store sum of each vector
	sumOfVectors := make([]int, len(vectorOfVectors))

	for i, vector := range vectorOfVectors {
		sum := 0
		for _, num := range vector {
			sum += num
		}
		sumOfVectors[i] = sum
	}

	// sort the array of sums with the sort package
	sort.Ints(sumOfVectors)

	// print the last 3 elements of the array
	fmt.Println("Last 3 elements:", sumOfVectors[len(sumOfVectors)-3:])

	// sum the last 3 elements
	sum := 0
	for _, num := range sumOfVectors[len(sumOfVectors)-3:] {
		sum += num
	}

	// print the sum
	fmt.Println("Sum of last 3 elements:", sum)
}
