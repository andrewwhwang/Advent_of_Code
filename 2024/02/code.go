package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func all_inc(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i-1] > report[i] {
			return false
		}
	}
	return true
}

func all_dec(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i-1] < report[i] {
			return false
		}
	}
	return true
}

func step(report []int) bool {
	for i := 1; i < len(report); i++ {
		diff := abs(report[i] - report[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isGood(report []int) bool {
	return (all_inc(report) || all_dec(report)) && step(report)
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
		count := 0
		for _, line := range strings.Split(input, "\n") {
			str_reports := strings.Split(line, " ")
			int_reports := make([]int, len(str_reports))
			for i, s := range str_reports {
				int_reports[i], _ = strconv.Atoi(s)
			}

			if isGood(int_reports) {
				count += 1
				continue
			}
			for i := 0; i < len(int_reports); i++ {
				sub := append(append([]int{}, int_reports[:i]...), int_reports[i+1:]...)
				if isGood(sub) {
					count += 1
					break
				}
			}

		}
		return count
	}
	count := 0
	for _, line := range strings.Split(input, "\n") {
		str_reports := strings.Split(line, " ")
		int_reports := make([]int, len(str_reports))
		for i, s := range str_reports {
			int_reports[i], _ = strconv.Atoi(s)
		}
		if isGood(int_reports) {
			count += 1
		}
	}
	// solve part 1 here
	return count
}
