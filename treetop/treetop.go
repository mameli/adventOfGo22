package main

import (
	"go_advent/utility"
	"strconv"
)

func main() {
	lines := utility.Read_text_file("./inputs/8_treetop.txt")

	forest_x := len(lines[0])
	forest_y := len(lines)

	trees := make([][]int, forest_y)
	for i := 0; i < forest_y; i++ {
		for _, c := range lines[i] {
			tmpInt, _ := strconv.Atoi(string(c))
			trees[i] = append(trees[i], tmpInt)
		}
	}

	// pivot trees array
	treesPivot := make([][]int, forest_x)
	for i := range treesPivot {
		treesPivot[i] = make([]int, forest_x)
	}

	for i := 0; i < forest_x; i++ {
		for j := 0; j < forest_y; j++ {
			treesPivot[j][i] = trees[i][j]
		}
	}

	visible_trees := 0
	// edge trees
	visible_trees += forest_x * 2
	visible_trees += forest_y * 2
	visible_trees -= 4

	for i := 1; i < forest_y-1; i++ {
		for j := 1; j < forest_x-1; j++ {
			visible_trees += checkVisibleTrees(trees[i], treesPivot[j], i, j, trees[i][j])
		}
	}

	println(visible_trees)
	max_scenic_score := 0
	tmp_scenic_score := 0
	for i := 1; i < forest_y-1; i++ {
		for j := 1; j < forest_x-1; j++ {
			tmp_scenic_score = calcScenicScore(trees[i], treesPivot[j], i, j, trees[i][j])
			if max_scenic_score < tmp_scenic_score {
				max_scenic_score = tmp_scenic_score
			}
		}
	}

	println("--------------------")
	println(max_scenic_score)
}

func checkVisibleTrees(row []int, col []int, i, j, tree int) int {
	tree_left := row[:j]
	tree_right := row[j+1:]
	tree_up := col[:i]
	tree_down := col[i+1:]
	sl := checkSight(tree_left, tree)
	sr := checkSight(tree_right, tree)
	su := checkSight(tree_up, tree)
	sd := checkSight(tree_down, tree)
	if sl || sr || su || sd {
		return 1
	}
	return 0
}

func checkSight(treesOnSight []int, tree int) bool {
	for _, t := range treesOnSight {
		if t >= tree {
			return false
		}
	}
	return true
}

func calcScenicScore(row []int, col []int, i, j, tree int) int {
	tree_left := row[:j]
	tree_right := row[j+1:]
	tree_up := col[:i]
	tree_down := col[i+1:]
	tree_left = reverseTrees(tree_left)
	tree_up = reverseTrees(tree_up)

	sl := calcTreeInSight(tree_left, tree)
	sr := calcTreeInSight(tree_right, tree)
	su := calcTreeInSight(tree_up, tree)
	sd := calcTreeInSight(tree_down, tree)
	return sl * sr * su * sd
}

func calcTreeInSight(treesOnSight []int, tree int) int {
	treeCounter := 0
	for _, t := range treesOnSight {
		if t >= tree {
			treeCounter++
			return treeCounter
		} else {
			treeCounter++
		}
	}
	return treeCounter
}

func reverseTrees(trees []int) []int {
	reversedTrees := make([]int, len(trees))
	for i := len(trees) - 1; i >= 0; i-- {
		reversedTrees[len(trees)-1-i] = trees[i]
	}
	return reversedTrees
}
