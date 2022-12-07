package assignment

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/vansante/adventofcode/2022/util"
)

type Day07 struct{}

var d07IDCounter int64

type d07Node struct {
	id       int64
	parent   *d07Node
	name     string
	isDir    bool
	size     int64
	children []*d07Node
}

func (n *d07Node) print(lvl int) {
	tp := "dir"
	if !n.isDir {
		tp = "file"
	}
	fmt.Printf("%s%s (%d. %s, s=%d)\n",
		strings.Repeat(" ", lvl),
		n.name,
		n.id,
		tp,
		n.size,
	)
	for i := range n.children {
		n.children[i].print(lvl + 2)
	}
}

func (n *d07Node) insert(name string) *d07Node {
	idx := sort.Search(len(n.children), func(i int) bool {
		return name <= n.children[i].name
	})

	var nd *d07Node
	if idx >= len(n.children) || n.children[idx].name != name {
		// we should insert the new node:
		d07IDCounter++
		nd = &d07Node{
			id:       d07IDCounter,
			parent:   n,
			name:     name,
			children: make([]*d07Node, 0, 64),
		}
		n.children = append(n.children[:idx], append([]*d07Node{nd}, n.children[idx:]...)...)
	}
	return n.children[idx]
}

func (n *d07Node) setDirSizes() {
	sizeMap := make(map[int64]int64, 10_000)
	n.walk(func(nd *d07Node) {
		if nd.parent == nil {
			return
		}

		if nd.isDir {
			sizeMap[nd.parent.id] += sizeMap[nd.id]
		} else {
			sizeMap[nd.parent.id] += nd.size
		}
	})

	n.walk(func(nd *d07Node) {
		if !nd.isDir {
			return
		}

		nd.size = sizeMap[nd.id]
	})
}

func (n *d07Node) walk(walker func(node *d07Node)) int64 {
	visited := make(map[int64]bool, 10_000)
	stack := make([]*d07Node, 1, 128)
	stack[0] = n
	var nd *d07Node

	var sum int64
	for len(stack) > 0 {
		nd, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if len(nd.children) > 0 && !visited[nd.id] {
			stack = append(stack, nd)
			for i := range nd.children {
				stack = append(stack, nd.children[i])
			}
			visited[nd.id] = true
			continue
		}

		walker(nd)
	}
	return sum
}

func (n *d07Node) sumDirSizes(max int64) int64 {
	var sum int64

	n.walk(func(nd *d07Node) {
		if nd.parent == nil || !nd.isDir {
			return
		}

		if nd.size <= max {
			sum += nd.size
		}
	})
	return sum
}

func (n *d07Node) findSmallestDir(min int64) int64 {
	dirSizes := make([]int64, 0, 100)
	n.walk(func(nd *d07Node) {
		if nd.parent == nil || !nd.isDir {
			return
		}

		if nd.size >= min {
			dirSizes = append(dirSizes, nd.size)
		}
	})
	return util.MinSlice(dirSizes)
}

func (d *Day07) getRootFs(lines []string) *d07Node {
	if lines[0] != "$ cd /" {
		panic("no root first")
	}

	stack := make([]*d07Node, 1, 128)
	curNode := &d07Node{
		id:       -1,
		name:     "/",
		isDir:    true,
		children: make([]*d07Node, 0, 64),
	}
	stack[0] = curNode

lineLoop:
	for _, line := range lines {
		isCmd := line[:1] == "$"

		if isCmd {
			switch line[2:4] {
			case "cd":
				switch line[5:] {
				case "..":
					if len(stack) <= 1 {
						continue lineLoop
					}
					curNode, stack = stack[len(stack)-1], stack[:len(stack)-1]
					continue lineLoop
				case "/":
					curNode, stack = stack[0], stack[:1]
					continue lineLoop
				default:
					nd := curNode.insert(line[5:])
					nd.isDir = true
					stack = append(stack, curNode)
					curNode = nd
					continue lineLoop
				}
			case "ls":
				// Nothing to do.
				continue lineLoop
			}
		}

		// We are now listing the current dir
		if strings.HasPrefix(line, "dir ") {
			nd := curNode.insert(line[4:])
			nd.isDir = true
			continue
		}

		parts := strings.Split(line, " ")
		nd := curNode.insert(parts[1])
		nd.isDir = false
		var err error
		nd.size, err = strconv.ParseInt(parts[0], 10, 64)
		util.CheckErr(err)
	}

	return stack[0]
}

func (d *Day07) SolveI(input string) any {
	lines := util.SplitLines(input)
	root := d.getRootFs(lines)
	root.setDirSizes()

	return root.sumDirSizes(100_000)
}

func (d *Day07) SolveII(input string) any {
	lines := util.SplitLines(input)
	root := d.getRootFs(lines)
	root.setDirSizes()

	total := int64(70_000_000)
	unused := total - root.size

	return root.findSmallestDir(30_000_000 - unused)
}
