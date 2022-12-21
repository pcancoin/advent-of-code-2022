package main

import (
	"flag"
	"fmt"
	"strings"

	"pcancoin/hello/util"
)

func main() {
	var part int
	var useTestInput bool
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.BoolVar(&useTestInput, "test", false, "Use input-test.txt")
	flag.Parse()
	fmt.Println("Running part", part)

	var input string

	if useTestInput {
		input = "./input-test.txt"
	} else {
		input = "./input.txt"
	}

	if part == 1 {
		ans := part1(util.ReadFile(input))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile(input))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	data := parseInput(input)

	for i := 0; i < len(data)-3; i++ {
		slice := data[i : i+4]
		fmt.Println(slice, hasAllElementsDistincts(slice))
		if hasAllElementsDistincts(slice) {
			return i + 1 + 3
		}
	}

	return -1
}

func part2(input string) int {
	data := parseInput(input)

	for i := 0; i < len(data)-13; i++ {
		slice := data[i : i+14]
		fmt.Println(slice, hasAllElementsDistincts(slice))
		if hasAllElementsDistincts(slice) {
			return i + 1 + 13
		}
	}

	return -1
}

func parseInput(input string) (data []string) {
	data = strings.Split(input, "")
	return
}

func hasAllElementsDistincts(array []string) bool {
	s := map[string]bool{}
	for _, item := range array {
		s[item] = true
	}
	return len(s) == len(array)
}
