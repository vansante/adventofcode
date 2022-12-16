package assignment

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/vansante/adventofcode/2022/util"
)

type Day16 struct{}

type d16Valve struct {
	id          string
	flow        int
	connections []string
}

func (d *Day16) getValves(input string) d16Valves {
	lines := util.SplitLines(input)
	m := make(map[string]*d16Valve)
	for _, line := range lines {
		v := &d16Valve{}
		n, err := fmt.Sscanf(line, "Valve %s has flow rate=%d;", &v.id, &v.flow)
		util.CheckErr(err)
		if n != 2 {
			panic("invalid matches for valve")
		}

		n = strings.Index(line, ";")
		if n < 0 {
			panic("no separator found")
		}
		connStr := strings.TrimSpace(line[n+len(" tunnels lead to valves "):])
		v.connections = strings.Split(connStr, ", ")
		m[v.id] = v
	}
	return m
}

type d16Valves map[string]*d16Valve

// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func (v d16Valves) distances(start string) map[string]int {
	visited := make(map[string]struct{}, 128)
	dist := make(map[string]int, 128)
	prev := make(map[string]string, 128)
	queue := make([]*d16Valve, 1, 128)

	const MaxDist = math.MaxInt

	for s := range v {
		dist[s] = MaxDist
	}

	dist[start] = 0
	queue[0] = v[start]

	for len(queue) > 0 {
		var cur *d16Valve
		cur, queue = queue[len(queue)-1], queue[:len(queue)-1]
		visited[cur.id] = struct{}{}

		for _, id := range cur.connections {
			neighbour := v[id]
			if _, ok := visited[id]; ok {
				continue
			}

			shortestDist := dist[cur.id] + 1 // Cost is always 1
			currentDist := dist[neighbour.id]
			if shortestDist >= currentDist {
				continue
			}

			// Insert neighbour into priority queue, lowest distance last
			idx := sort.Search(len(queue), func(i int) bool {
				return dist[queue[i].id] <= shortestDist
			})

			queue = append(queue[:idx], append([]*d16Valve{v[id]}, queue[idx:]...)...)

			dist[neighbour.id] = shortestDist
			prev[neighbour.id] = cur.id
		}
	}

	return dist
}

func (v d16Valves) distance(start, end string) (int, []string) {
	visited := make(map[string]struct{}, 128)
	dist := make(map[string]int, 128)
	prev := make(map[string]string, 128)
	queue := make([]*d16Valve, 1, 128)

	const MaxDist = math.MaxInt

	for s := range v {
		dist[s] = MaxDist
	}

	dist[start] = 0
	queue[0] = v[start]

	for len(queue) > 0 {
		var cur *d16Valve
		cur, queue = queue[len(queue)-1], queue[:len(queue)-1]
		if cur.id == end {
			break
		}
		visited[cur.id] = struct{}{}

		for _, id := range cur.connections {
			neighbour := v[id]
			if _, ok := visited[id]; ok {
				continue
			}

			shortestDist := dist[cur.id] + 1 // Cost is always 1
			currentDist := dist[neighbour.id]
			if shortestDist >= currentDist {
				continue
			}

			// Insert neighbour into priority queue, lowest distance last
			idx := sort.Search(len(queue), func(i int) bool {
				return dist[queue[i].id] <= shortestDist
			})

			queue = append(queue[:idx], append([]*d16Valve{v[id]}, queue[idx:]...)...)

			dist[neighbour.id] = shortestDist
			prev[neighbour.id] = cur.id
		}
	}

	if dist[end] == MaxDist {
		panic("no route found")
	}

	path := make([]string, 0, dist[end])
	target := end
	for prev[target] != "" || target == start {
		for target != "" {
			path = append([]string{target}, path...)
			target = prev[target]
		}
	}

	return dist[end], path
}

type d16Score struct {
	id       string
	flow     int
	distance int
}

func (v d16Valves) getValveScores(start string, skip map[string]struct{}) []d16Score {
	distances := v.distances(start)

	l := make([]d16Score, 0, len(v))
	for _, valve := range v {
		if _, ok := skip[valve.id]; ok {
			continue
		}

		l = append(l, d16Score{
			id:       valve.id,
			flow:     valve.flow,
			distance: distances[valve.id],
		})
	}
	return l
}

func (v d16Valves) releasePressure(start string) int {
	opened := make(map[string]struct{})

	var released int
	var releaseRate int
	for seconds := 0; seconds <= 30; seconds++ {
		scores := v.getValveScores(start, opened)
		sort.Slice(scores, func(i, j int) bool {
			s1 := scores[i]
			s2 := scores[j]

			if s1.flow-s1.distance == s2.flow-s2.distance {
				return s1.distance < s2.distance
			}

			return s1.flow-s1.distance > s2.flow-s2.distance
		})

		fmt.Println(scores)
		for i := 0; i < scores[0].distance; i++ {
			fmt.Println("Walking to ", scores[0], " dist ", scores[0].distance)
			released += releaseRate
			seconds++
		}

		seconds++
		releaseRate += scores[0].flow
		released += releaseRate

		fmt.Println("Open ", scores[0].id)

		opened[scores[0].id] = struct{}{} // Set valve to be opened
		start = scores[0].id
	}

	fmt.Println(releaseRate)

	return 0
}

func (d *Day16) SolveI(input string) any {
	valves := d.getValves(input)

	fmt.Println()
	fmt.Println(valves.distance("AA", "JJ"))
	fmt.Println()

	valves.releasePressure("AA")

	return 0
}

func (d *Day16) SolveII(input string) any {
	return "Not Implemented Yet"
}
