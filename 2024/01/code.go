package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		counter := make(map[int]int)

		for _, line := range strings.Split(input, "\n") {
			split := strings.Split(line, "   ")
			i2, _ := strconv.Atoi(split[1])
			counter[i2] += 1
		}

		var count int
		for _, line := range strings.Split(input, "\n") {
			split := strings.Split(line, "   ")
			i1, _ := strconv.Atoi(split[0])
			count += i1 * counter[i1]
		}
		return count
	}
	var list1 []int
	var list2 []int

	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, "   ")
		i1, _ := strconv.Atoi(split[0])
		list1 = append(list1, i1)

		i2, _ := strconv.Atoi(split[1])
		list2 = append(list2, i2)

	}
	slices.Sort(list1)
	slices.Sort(list2)

	var counter = 0
	for i := range list1 {
		counter += max(list1[i]-list2[i], list2[i]-list1[i])
	}
	// solve part 1 here
	return counter
}
