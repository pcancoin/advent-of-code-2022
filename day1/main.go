package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"pcancoin/hello/cast"
	"pcancoin/hello/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	ans, _ := compute(input)
	return ans
}

func part2(input string) int {
	_, ans := compute(input)
	return ans
}

func compute(input string) (ans int, indexAns int) {
	var quantities []int = make([]int, 0)
	quantities = append(quantities, 0)
	for _, l := range strings.Split(input, "\n") {
		if len(l) > 0 {
			// Still the same Elf
			quantities[len(quantities)-1] += cast.ToInt(l)
		} else {
			// New Elf
			quantities = append(quantities, 0)
		}
	}
	sort.Slice(quantities, func(i, j int) bool {
		return quantities[i] > quantities[j]
	})
	return quantities[0], quantities[0] + quantities[1] + quantities[2]
}
