package main

import (
	"fmt"
	"go_advent/utility"
	"math"
	"strconv"
)

type Rope struct {
	dx int
	dy int
}

func main() {

	moves := utility.Read_text_file("./inputs/9_rope.txt")

	head := Rope{dx: 0, dy: 0}
	tail := Rope{dx: 0, dy: 0}

	positions := make(map[string]bool)

	for _, m := range moves {
		d := rune(m[0])
		n, _ := strconv.Atoi(m[2:])
		for i := 0; i < n; i++ {
			head.move(d, 1)
			nx, ny := new_tail_pos(head, tail)
			tail.dx = nx
			tail.dy = ny
			positions[fmt.Sprintf("%d,%d", nx, ny)] = true
		}
	}

	// println(len(positions))

	// part 2
	positions = make(map[string]bool)

	var knots = [10]Rope{}
	for i := 0; i < len(knots); i++ {
		knots[i] = Rope{dx: 0, dy: 0}
	}

	tailIndex := len(knots) - 1

	for _, m := range moves {
		d := rune(m[0])
		n, _ := strconv.Atoi(m[2:])
		for i := 0; i < n; i++ {
			knots[0].move(d, 1)
			for ik := 0; ik < len(knots)-1; ik++ {
				nx, ny := new_tail_pos(knots[ik], knots[ik+1])
				knots[ik+1].dx = nx
				knots[ik+1].dy = ny
			}
			positions[fmt.Sprintf("%d,%d", knots[tailIndex].dx, knots[tailIndex].dy)] = true
		}
	}
	println(len(positions))
}

func (r *Rope) move(d rune, n int) {
	switch d {
	case 'R':
		r.dy += n
	case 'L':
		r.dy -= n
	case 'U':
		r.dx -= n
	case 'D':
		r.dx += n
	}
}

func new_tail_pos(h Rope, t Rope) (int, int) {
	d := euclideanDistance(h, t)
	if d < 2 {
		return t.dx, t.dy
	}
	signX := 0
	signY := 0
	if h.dx != t.dx {
		signX = (h.dx - t.dx) / int(math.Abs(float64(h.dx-t.dx)))
	}
	if h.dy != t.dy {
		signY = (h.dy - t.dy) / int(math.Abs(float64(h.dy-t.dy)))
	}

	return t.dx + signX, t.dy + signY
}

func euclideanDistance(rope1, rope2 Rope) float64 {
	dx := rope1.dx - rope2.dx
	dy := rope1.dy - rope2.dy
	return math.Sqrt(float64(dx*dx + dy*dy))
}
