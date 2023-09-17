package main

import (
	"go_advent/utility"
	"strconv"
	"strings"
)

func main() {
	lines := utility.Read_text_file("./inputs/4_cleanup.txt")

	full_overlap_count := 0
	overlap_count := 0
	for _, line := range lines[:] {
		first, second := strings.Split(line, ",")[0], strings.Split(line, ",")[1]

		s1, _ := strconv.Atoi(strings.Split(first, "-")[0])
		e1, _ := strconv.Atoi(strings.Split(first, "-")[1])
		s2, _ := strconv.Atoi(strings.Split(second, "-")[0])
		e2, _ := strconv.Atoi(strings.Split(second, "-")[1])

		if contains(s1, e1, s2, e2) || contains(s2, e2, s1, e1) {
			full_overlap_count++
		}

		if !disjoint(s1, e1, s2, e2) {
			overlap_count++
		}
	}

	println(full_overlap_count)
	println(overlap_count)
}

func contains(s1 int, e1 int, s2 int, e2 int) bool {
	return s1 <= s2 && e1 >= e2
}

func disjoint(s1 int, e1 int, s2 int, e2 int) bool {
	return s1 > e2 || e1 < s2
}
