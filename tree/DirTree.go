package tree

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Type int

const (
	Folder Type = iota + 1
	File
)

type Node struct {
	Value    string
	Size     int
	Children []*Node
	Parent   *Node
	NodeType Type
}

type DirTree struct {
	Root       *Node
	CurrentDir *Node
}

func EmptyDirTree() *DirTree {
	initialNode := Node{
		"/",
		0,
		[]*Node{},
		nil,
		Folder,
	}
	return &DirTree{
		&initialNode,
		&initialNode,
	}
}

func ChangeDir(t *DirTree, dir string) (*DirTree, error) {
	if dir == ".." {
		if t.CurrentDir.Parent == nil {
			return t, nil
		}
		t.CurrentDir = t.CurrentDir.Parent
		return t, nil
	}
	if dir == "/" {
		t.CurrentDir = t.Root
		return t, nil
	}
	for _, child := range t.CurrentDir.Children {
		if child.Value == dir {
			t.CurrentDir = child
			return t, nil
		}
	}
	return t, errors.New("Cannot find the given folder " + dir)
}

func CreateFile(t *DirTree, fileName string, size int) (*DirTree, error) {
	for _, child := range t.CurrentDir.Children {
		if child.Value == fileName {
			return t, errors.New("File already existing " + fileName)
		}
	}
	t.CurrentDir.Children = append(t.CurrentDir.Children, &Node{
		fileName,
		size,
		[]*Node{},
		t.CurrentDir,
		File,
	})
	UpdateSize(t.CurrentDir, size)
	return t, nil
}

func CreateDir(t *DirTree, dir string) (*DirTree, error) {
	for _, child := range t.CurrentDir.Children {
		if child.Value == dir {
			return t, errors.New("Folder already existing : " + dir)
		}
	}
	t.CurrentDir.Children = append(t.CurrentDir.Children, &Node{
		dir,
		0,
		nil,
		t.CurrentDir,
		Folder,
	})
	return t, nil
}

func UpdateSize(n *Node, size int) {
	if n == nil {
		return
	}
	n.Size += size
	UpdateSize(n.Parent, size)
}

func Print(t *DirTree) {
	PrintNode(t.Root, 0)
}

func PrintNode(n *Node, depth int) {
	if n == nil {
		return
	}
	str := strings.Repeat("--", depth)
	fmt.Println(str + " " + n.Value + " (" + typeToString(n.NodeType) + ", size " + strconv.Itoa(n.Size) + ")")
	for _, child := range n.Children {
		PrintNode(child, depth+1)
	}
}

func Walker(t *DirTree) <-chan *Node {
	ch := make(chan *Node)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

func Walk(t *DirTree, ch chan *Node) {
	if t == nil {
		return
	}
	WalkNode(t.Root, ch)
}

func WalkNode(n *Node, ch chan *Node) {
	if n == nil {
		return
	}
	ch <- n
	for _, child := range n.Children {
		WalkNode(child, ch)
	}
}

// func insert(t *DirTree, dir string, v string) *DirTree {
// 	if t == nil {
// 		node := Node{v, nil}
// 		return &DirTree{&node}
// 	}
// 	if v < t.Value {
// 		t.Left = insert(t.Left, v)
// 		return t
// 	}
// 	t.Right = insert(t.Right, v)
// 	return t
// }

func typeToString(t Type) string {
	switch t {
	case File:
		return "File"
	case Folder:
		return "Folder"
	default:
		return "Unknown"
	}
}
