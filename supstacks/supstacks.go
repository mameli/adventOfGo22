package main

import (
	"go_advent/utility"
	"strconv"
	"strings"
)

func main() {
	lines := utility.Read_text_file("./inputs/5_stacks.txt")

	// new list of string called stack_config
	stack_config, move_config := parse_config(lines)

	crate_indexes := []int{}
	for i := 0; i < len(stack_config[0]); i++ {
		crate_indexes = append(crate_indexes, i+1)
		i += 3
	}

	stacks := []utility.RuneStack{}
	stacks = create_stack(stack_config, crate_indexes, stacks)

	stacksN := []utility.RuneStack{}
	stacksN = create_stack(stack_config, crate_indexes, stacksN)

	for _, move := range move_config[:] {
		parsed_move := strings.Split(move, " ")

		n_crates, _ := strconv.Atoi(parsed_move[1])
		from_crate, _ := strconv.Atoi(parsed_move[3])
		to_crate, _ := strconv.Atoi(parsed_move[5])

		for i := 0; i < n_crates; i++ {
			tmp_crate, _ := stacks[from_crate-1].Pop()
			stacks[to_crate-1].Push(tmp_crate)
			continue
		}
	}

	print_stacks(stacks)
	for _, stack := range stacks[:] {
		tmp_crate, _ := stack.Pop()
		print(string(tmp_crate))
	}

	for _, move := range move_config[:] {
		parsed_move := strings.Split(move, " ")

		n_crates, _ := strconv.Atoi(parsed_move[1])
		from_crate, _ := strconv.Atoi(parsed_move[3])
		to_crate, _ := strconv.Atoi(parsed_move[5])

		tmp_crate, _ := stacksN[from_crate-1].PopN(n_crates)
		stacksN[to_crate-1].PushN(tmp_crate)
	}
	println("\n----------------------------------")
	print_stacks(stacksN)
	for _, stack := range stacksN[:] {
		tmp_crate, _ := stack.Pop()
		print(string(tmp_crate))
	}
	println("\n----------------------------------")
}

func create_stack(stack_config []string, crate_indexes []int, stacks []utility.RuneStack) []utility.RuneStack {
	for _, row := range stack_config[:len(stack_config)-1] {
		for i, index := range crate_indexes[:] {
			if len(stacks) < i+1 {
				var stack utility.RuneStack
				stacks = append(stacks, stack)
			}
			if row[index] != ' ' {
				stacks[i].Push(rune(row[index]))
			}
		}
	}

	for _, stack := range stacks[:] {
		stack.Reverse()
	}

	return stacks
}

func print_stacks(stacks []utility.RuneStack) {
	for i, stack := range stacks[:] {
		println(i+1, string(stack))
	}
}

func parse_config(config []string) ([]string, []string) {
	stack_config := []string{}
	move_config := []string{}
	check_stack := true
	for _, line := range config[:] {

		if check_stack {
			stack_config = append(stack_config, line)
		} else {
			move_config = append(move_config, line)
			continue
		}

		if line == "" {
			check_stack = false
			stack_config = stack_config[:len(stack_config)-1]
		}
	}
	return stack_config, move_config
}
