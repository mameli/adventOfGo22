package utility

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Read_text_file(filePath string) []string {
	// read the file
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return []string{}
	}

	// Convert the byte slice to a string
	content := string(contentBytes)

	// split the string with empty line and store it in a slice
	lines := strings.Split(content, "\n")

	return lines
}

func Print_array_rune(arr []rune) {
	for _, line := range arr[:] {
		print(string(line) + " ")
	}
}

func Print_array_string(arr []string) {
	for _, line := range arr[:] {
		println(line)
	}
}

type RuneStack []rune

// Push adds a rune to the top of the stack.
func (s *RuneStack) Push(item rune) {
	*s = append(*s, item)
}

// Pop removes and returns the top rune from the stack.
func (s *RuneStack) Pop() (rune, error) {
	if s.IsEmpty() {
		return 0, errors.New("rune stack is empty")
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top, nil
}

// pop multiple items
func (s *RuneStack) PopN(n int) ([]rune, error) {
	if s.IsEmpty() {
		return []rune{}, errors.New("rune stack is empty")
	}
	if n > len(*s) {
		return []rune{}, errors.New("n is larger than the stack size")
	}
	top := (*s)[len(*s)-n:]
	*s = (*s)[:len(*s)-n]
	return top, nil
}

// push multiple items
func (s *RuneStack) PushN(items []rune) {
	for _, item := range items[:] {
		s.Push(item)
	}
}

// IsEmpty checks if the stack is empty.
func (s *RuneStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *RuneStack) Reverse() {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
