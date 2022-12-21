package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"

	"pcancoin/hello/tree"
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

	dirTree := tree.EmptyDirTree()
	for _, command := range data {
		commandName := strings.Split(strings.Split(command, "\n")[0], " ")[0]
		switch commandName {
		case "cd":
			handleCd(dirTree, command)
			break
		case "ls":
			handleLs(dirTree, command)
			break
		}
	}

	total := 0
	for node := range tree.Walker(dirTree) {
		if node.NodeType == tree.Folder && node.Size < 100000 {
			fmt.Println(node)
			total += node.Size
		}
	}
	return total
}

func part2(input string) int {
	data := parseInput(input)

	dirTree := tree.EmptyDirTree()
	for _, command := range data {
		commandName := strings.Split(strings.Split(command, "\n")[0], " ")[0]
		switch commandName {
		case "cd":
			handleCd(dirTree, command)
			break
		case "ls":
			handleLs(dirTree, command)
			break
		}
	}

	min := 0

	free := 70000000 - dirTree.Root.Size
	toFree := 30000000 - free
	fmt.Printf("Free : %d, to free : %d\n", free, toFree)
	for node := range tree.Walker(dirTree) {
		if node.NodeType == tree.Folder && node.Size >= toFree {
			fmt.Printf("node %s size %d\n", node.Value, node.Size)
			if math.Abs(float64(toFree)-float64(node.Size)) < math.Abs(float64(toFree)-float64(min)) {
				min = node.Size
			}
		}
	}
	return min
	// 44795677 to high
}

func parseInput(input string) (data []string) {
	data = strings.Split(strings.Trim(input, "$ "), "\n$ ")
	return
}

func handleCd(t *tree.DirTree, command string) {
	dir := strings.Split(command, " ")[1]
	t, err := tree.ChangeDir(t, dir)
	if err != nil {
		fmt.Println("Error :" + err.Error())
	}
}

func handleLs(dirTree *tree.DirTree, command string) {
	list := strings.Split(command, "\n")[1:]
	for _, line := range list {
		infos := strings.Split(line, " ")
		if infos[0] == "dir" {
			t, err := tree.CreateDir(dirTree, infos[1])
			if err != nil {
				panic(err)
			}
			dirTree = t
		} else {
			size, err := strconv.Atoi(infos[0])
			t, err := tree.CreateFile(dirTree, infos[1], size)
			if err != nil {
				panic(err)
			}
			dirTree = t
		}
	}
}
