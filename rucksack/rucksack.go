package main

import (
	"go_advent/utility"
)

// create set of chars from string
func create_char_set(str string) map[rune]bool {
	char_set := make(map[rune]bool)
	for _, char := range str {
		char_set[char] = true
	}
	return char_set
}

// intersect two sets
func intersect_sets(set1 map[rune]bool, set2 map[rune]bool) map[rune]bool {
	intersect_set := make(map[rune]bool)
	for key := range set1 {
		if set2[key] {
			intersect_set[key] = true
		}
	}
	return intersect_set
}

func main() {
	lines := utility.Read_text_file("./inputs/3_rucksack.txt")

	priority_sum := 0
	for _, line := range lines[:] {
		var first_half = line[:len(line)/2]
		var second_half = line[len(line)/2:]

		first_char_set := create_char_set(first_half)
		second_char_set := create_char_set(second_half)
		inter := intersect_sets(first_char_set, second_char_set)

		for key := range inter {
			priority_sum += get_priority_char(key)
		}
	}

	// println(priority_sum)

	triplets := []string{}
	priority_sum = 0
	for _, line := range lines[:] {
		if len(triplets) == 2 {
			triplets = append(triplets, line)

			set1 := create_char_set(triplets[0])
			set2 := create_char_set(triplets[1])
			set3 := create_char_set(triplets[2])

			inter1 := intersect_sets(set1, set2)
			inter2 := intersect_sets(inter1, set3)

			for key := range inter2 {
				priority_sum += get_priority_char(key)
			}

			triplets = []string{}
		} else {
			triplets = append(triplets, line)
		}
	}

	println(priority_sum)

}

func get_priority_char(key rune) int {
	// Create a map to store the letter-to-number mapping
	lowerLetterToNumber := make(map[rune]int)
	upperLetterToNumber := make(map[rune]int)

	// Populate the map with the desired mappings
	for letter := 'a'; letter <= 'z'; letter++ {
		// Calculate the corresponding number (1 to 26)
		number := int(letter - 'a' + 1)
		lowerLetterToNumber[letter] = number
	}

	for letter := 'A'; letter <= 'Z'; letter++ {
		// Calculate the corresponding number (1 to 26)
		number := int(letter - 'A' + 27)
		upperLetterToNumber[letter] = number
	}

	//check if letter is lower or upper
	if _, ok := lowerLetterToNumber[key]; ok {
		return lowerLetterToNumber[key]
	} else if _, ok := upperLetterToNumber[key]; ok {
		return upperLetterToNumber[key]
	} else {
		return 0
	}
}
