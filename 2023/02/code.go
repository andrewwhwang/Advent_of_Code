package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func valid_set(sets string) bool {

	rule := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, set := range strings.Split(sets, "; ") {
		for _, draw := range strings.Split(set, ", ") {
			draw_split := strings.Split(draw, " ")
			num, _ := strconv.Atoi(draw_split[0])
			color := draw_split[1]
			value, ok := rule[color]
			if !ok || value < num {
				return false
			}
		}
	}
	return true
}

func get_min_set(sets string) int {
	min_set := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, set := range strings.Split(sets, "; ") {
		for _, draw := range strings.Split(set, ", ") {
			draw_split := strings.Split(draw, " ")
			num, _ := strconv.Atoi(draw_split[0])
			color := draw_split[1]

			value, _ := min_set[color]
			if num > value {
				min_set[color] = num
			}
		}
	}

	return min_set["red"] * min_set["green"] * min_set["blue"]
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
		var part2_counter int

		for _, line := range strings.Split(input, "\n") {
			split := strings.Split(line, ": ")
			sets := split[1]
			part2_counter += get_min_set(sets)
		}
		return part2_counter
	}
	var counter int
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, ": ")
		game := split[0]
		sets := split[1]
		game_num, _ := strconv.Atoi(strings.Split(game, " ")[1])
		if valid_set(sets) {
			counter += game_num
		}
	}

	return counter
}
