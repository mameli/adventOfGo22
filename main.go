package main

import "go_advent/utility"

func main() {
	lines := utility.Read_text_file("./inputs/input.txt")

	println("Hello world! ", len(lines))
}
