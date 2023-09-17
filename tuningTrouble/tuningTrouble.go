package main

import (
	"go_advent/utility"
)

func create_char_set(str string) map[rune]bool {
	char_set := make(map[rune]bool)
	for _, char := range str {
		char_set[char] = true
	}
	return char_set
}

func main() {
	lines := utility.Read_text_file("./inputs/6_tuningTroubles.txt")
	msg := lines[0]

	for i := 0; i < len(msg); i++ {
		msg_slice := msg[i : i+14]
		set := create_char_set(msg_slice)
		if len(set) == 14 {
			println(msg_slice, i+14)
			break
		}
	}

}
