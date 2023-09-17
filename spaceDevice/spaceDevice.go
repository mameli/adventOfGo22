package main

import (
	"go_advent/utility"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := utility.Read_text_file("/Users/filippomameli/Projects/go_advent/inputs/7_spaceDevice.txt")

	// utility.Print_array_string(lines)
	dirs := make(map[string]int)
	sub_dirs := make(map[string][]string)
	keys := []string{"/root/home"}

	ls_check := false
	current_dir := "/root"
	sub_folder := "/"
	for _, cmd := range lines[:] {
		if strings.Contains(cmd, "$ cd ..") {
			lastIndex := strings.LastIndex(current_dir, "/")
			current_dir = current_dir[:lastIndex]
			continue
		}
		if strings.Contains(cmd, "$ cd ") {
			dir := strings.Split(cmd, " ")[2]
			dirs[current_dir+"/"+dir] = 0.0
			current_dir = current_dir + "/" + dir
			ls_check = false
			continue
		}
		if strings.Contains(cmd, "$ ls") {
			ls_check = true
			continue
		}
		if ls_check {
			if strings.Contains(cmd, "dir") {
				sub_folder = current_dir + "/" + strings.Split(cmd, " ")[1]
				sub_dirs[current_dir] = append(sub_dirs[current_dir], sub_folder)
				keys = append(keys, sub_folder)
			} else {
				n_bytes, _ := strconv.Atoi(strings.Split(cmd, " ")[0])
				dirs[current_dir] += n_bytes
			}
		}
	}

	for i, j := 0, len(keys)-1; i < j; i, j = i+1, j-1 {
		keys[i], keys[j] = keys[j], keys[i]
	}

	for _, k := range keys {
		for _, v := range sub_dirs[k] {
			dirs[k] += dirs[string(v)]
		}
	}

	for k, v := range dirs {
		println(k + " -> " + strconv.Itoa(v))
	}

	sum_sub_dirs := 0
	for _, v := range dirs {
		if v <= 100000 {
			sum_sub_dirs += v
		}
	}
	println("Total size sub dirs " + strconv.Itoa(sum_sub_dirs))

	// task 2
	total_home_size := dirs["/root/home"]
	unused_size := 70000000 - total_home_size
	println("Unused size " + strconv.Itoa(unused_size))

	// all values of dirs into an array
	sizes := make([]int, 0, len(dirs))
	for _, v := range dirs {
		sizes = append(sizes, v)
	}

	// sort array
	sort.Ints(sizes)
	for i, v := range sizes {
		if 30000000 < unused_size+v {
			println("Min size " + strconv.Itoa(sizes[i]))
			break
		}
	}

}
