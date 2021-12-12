package assignment

import (
	"strings"
)

type Day12 struct{}

type d12Node struct {
	name  string
	edges []*d12Node
	big   bool
}

func (n *d12Node) String() string {
	return n.name
}

type d12Graph struct {
	nodes map[string]*d12Node
}

func (d *Day12) getGraph(input string) *d12Graph {
	lines := SplitLines(input)

	g := &d12Graph{
		nodes: make(map[string]*d12Node, len(lines)),
	}
	for _, line := range lines {
		lineNodes := strings.Split(line, "-")
		if len(lineNodes) > 2 {
			panic(line)
		}
		var nds []*d12Node
		for _, name := range lineNodes {
			nd, ok := g.nodes[name]
			if !ok {
				nd = &d12Node{
					name: name,
					big:  name != strings.ToLower(name),
				}
				g.nodes[name] = nd
			}
			nds = append(nds, nd)
		}
		nds[0].edges = append(nds[0].edges, nds[1])
		nds[1].edges = append(nds[1].edges, nds[0])
	}
	return g
}

func (d *Day12) findAllPaths(g *d12Graph, start, end string) int64 {
	visited := make(map[string]struct{}, 128)

	return g.findPath(start, end, visited)
}

func (g *d12Graph) findPath(start, end string, visited map[string]struct{}) int64 {
	cur := g.nodes[start]

	if !cur.big {
		visited[cur.name] = struct{}{}
	}

	sum := int64(0)
	if cur.name == end {
		sum++
	} else {
		for _, edge := range cur.edges {
			if _, ok := visited[edge.name]; ok {
				continue
			}
			sum += g.findPath(edge.name, end, visited)
		}
	}

	delete(visited, cur.name)
	return sum
}

func (g *d12Graph) findPathsIterative(start, end string) int64 {
	// https://stackoverflow.com/a/35187404
	type item struct {
		start   string
		edgeIdx int
	}
	stack := make([]*item, 0, 1024)
	visited := make(map[string]int, 128)
	visitedTwice := ""

	stack = append(stack, &item{start, 0})
	visited[start] = 10

	sum := int64(0)
	for len(stack) > 0 {
		if len(stack) > 10_000 {
			panic("stack too long")
		}
		var cur *item
		cur = stack[len(stack)-1] // peek
		curNd := g.nodes[cur.start]

		if curNd.name == end || cur.edgeIdx == len(curNd.edges) {
			if curNd.name == end {
				sum++
			}
			visited[curNd.name]--
			stack = stack[:len(stack)-1] // pop
			if visitedTwice == curNd.name {
				visitedTwice = ""
			}
			continue
		}

		edge := curNd.edges[cur.edgeIdx]
		cur.edgeIdx++

		visits := visited[edge.name]
		if !edge.big && (visits > 0 || edge.name == visitedTwice && visits > 1) {
			if visitedTwice != "" || edge.name == start {
				continue
			}
			visitedTwice = edge.name
		}

		stack = append(stack, &item{edge.name, 0})
		if !edge.big {
			visited[edge.name]++
		}
	}
	return sum
}

func (d *Day12) SolveI(input string) int64 {
	g := d.getGraph(input)

	return d.findAllPaths(g, "start", "end")
}

func (d *Day12) SolveII(input string) int64 {
	g := d.getGraph(input)

	return g.findPathsIterative("start", "end")
}
