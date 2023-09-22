package main

import (
	"go_advent/utility"
	"strconv"
	"strings"
)

func main() {
	cmds := utility.Read_text_file("./inputs/10_cathode.txt")

	x := 1
	signal := 0
	cycleNum := 0
	cycleCheck := 20
	for _, c := range cmds {
		cmdType := strings.Split(c, " ")[0]
		cycleNum++
		if cmdType != "noop" {
			cmdN := strings.Split(c, " ")[1]
			n, _ := strconv.Atoi(cmdN)
			cycleNum++
			if cycleNum >= cycleCheck {
				signal += x * cycleCheck
				cycleCheck += 40
			}
			x += n
		}
	}

	println(signal)

}
