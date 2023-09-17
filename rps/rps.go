package rps

import (
	"fmt"
	"go_advent/utility"
	"strings"
)

func get_score(play string) int {
	// map with scores
	scores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	// switch case with all possible combinations
	switch play {
	case "A A":
		return 3 + scores["X"]
	case "A B":
		return 6 + scores["Y"]
	case "A C":
		return 0 + scores["Z"]
	case "B A":
		return 0 + scores["X"]
	case "B B":
		return 3 + scores["Y"]
	case "B C":
		return 6 + scores["Z"]
	case "C A":
		return 6 + scores["X"]
	case "C B":
		return 0 + scores["Y"]
	case "C C":
		return 3 + scores["Z"]
	}

	return 0
}

func Get_final_score() {
	lines := utility.Read_text_file("./inputs/2_rps_games_test.txt")

	final_score := 0

	// loop through the lines
	for _, line := range lines {
		// get the score
		score := get_score(convert_play(line))

		final_score += score
	}

	fmt.Println(final_score)

	final_score = 0

	// loop through the lines
	for _, line := range lines {
		// get the score
		fmt.Println(find_play(line))
		score := get_score(find_play(line))

		final_score += score
	}

	fmt.Println(final_score)

}

func convert_play(play string) string {
	play = strings.ReplaceAll(play, "X", "A")
	play = strings.ReplaceAll(play, "Y", "B")
	play = strings.ReplaceAll(play, "Z", "C")
	return play
}

func find_play(play string) string {
	switch play {
	case "A X":
		return "A C"
	case "A Y":
		return "A A"
	case "A Z":
		return "A B"
	case "B X":
		return "B A"
	case "B Y":
		return "B B"
	case "B Z":
		return "B C"
	case "C X":
		return "C B"
	case "C Y":
		return "C C"
	case "C Z":
		return "C A"
	}
	return ""
}
