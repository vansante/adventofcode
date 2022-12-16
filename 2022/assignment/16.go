package assignment

import (
	"fmt"
	"sort"
	"strings"

	"github.com/vansante/adventofcode/2022/util"
)

type Day16 struct {
}

type d16Valve struct {
	id          uint8
	label       string
	flow        int
	connections []string
	connIDs     []uint8
}

func (d *Day16) getValves(input string) d16Valves {
	lines := util.SplitLines(input)
	valves := d16Valves{
		mp:  make(map[uint8]*d16Valve, len(lines)),
		lst: make([]*d16Valve, len(lines)),
	}

	for i, line := range lines {
		valves.idCount++
		v := &d16Valve{
			id: valves.idCount,
		}
		n, err := fmt.Sscanf(line, "Valve %s has flow rate=%d;", &v.label, &v.flow)
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
		valves.lst[i] = v
	}

	valves.setMap()
	return valves
}

type d16Valves struct {
	lst     []*d16Valve
	mp      map[uint8]*d16Valve
	lbls    map[string]uint8
	idCount uint8
}

func (v *d16Valves) getID(label string) uint8 {
	return v.lbls[label]
}

func (v *d16Valves) get(id uint8) *d16Valve {
	return v.mp[id]
}

func (v *d16Valves) setMap() {
	v.lbls = make(map[string]uint8)
	for i := range v.lst {
		valve := v.lst[i]
		v.mp[valve.id] = valve

		v.lbls[valve.label] = valve.id
	}

	for i := range v.lst {
		valve := v.lst[i]
		valve.connIDs = make([]uint8, len(valve.connections))
		for j := range valve.connections {
			valve.connIDs[j] = v.getID(valve.connections[j])
		}
	}
}

const maxValves = 20

type d16Opened struct {
	opened   [maxValves]uint8
	openSize int8
}

func (o *d16Opened) contains(id uint8) bool {
	idx := sort.Search(int(o.openSize), func(i int) bool {
		return o.opened[i] <= id
	})

	return idx < int(o.openSize) && o.opened[idx] == id
}

func (o *d16Opened) add(id uint8) {
	idx := sort.Search(int(o.openSize), func(i int) bool {
		return o.opened[i] <= id
	})
	if idx < int(o.openSize) && o.opened[idx] == id {
		return // Already in
	}

	if o.openSize >= maxValves-1 {
		panic("too many opened valves; increase array size")
	}

	if idx == int(o.openSize) {
		o.opened[idx] = id
		o.openSize++
		return
	}

	for i := int(o.openSize) + 1; i > idx; i-- {
		o.opened[i] = o.opened[i-1]
	}
	o.opened[idx] = id
	o.openSize++
}

type d16State struct {
	startID uint8
	minutes int8
	opened  d16Opened
}

type d16Release struct {
	states map[d16State]uint16
	valves d16Valves
}

func (r d16Release) releasePressure(valveID uint8, minutes int8, opened d16Opened, zeroMinute func(d16Opened) uint16) uint16 {
	if minutes <= 0 {
		if zeroMinute != nil {
			return zeroMinute(opened)
		}
		return 0
	}

	s := d16State{
		startID: valveID,
		minutes: minutes,
		opened:  opened,
	}

	if result, ok := r.states[s]; ok {
		return result
	}
	valve := r.valves.get(valveID)

	var released uint16
	for _, v := range valve.connIDs {
		// Without opening current valve:
		released = util.Max(
			released,
			r.releasePressure(v, minutes-1, opened, zeroMinute),
		)
	}

	// Is there a point in opening the current valve and is it opened already?
	if opened.contains(valve.id) || valve.flow <= 0 || minutes <= 0 {
		r.states[s] = released
		return released
	}

	// Now with opening current valve
	opened.add(valve.id)
	minutes--
	releaseSum := uint16(int(minutes) * valve.flow)
	for _, v := range valve.connIDs {
		released = util.Max(
			released,
			releaseSum+r.releasePressure(v, minutes-1, opened, zeroMinute),
		)
	}
	r.states[s] = released
	return released
}

func (d *Day16) SolveI(input string) any {
	valves := d.getValves(input)

	release := d16Release{
		valves: valves,
		states: make(map[d16State]uint16, 1_000_000),
	}
	return release.releasePressure(valves.getID("AA"), 30, d16Opened{}, nil)
}

func (d *Day16) SolveII(input string) any {
	valves := d.getValves(input)

	release := d16Release{
		valves: valves,
		states: make(map[d16State]uint16, 1_000_000),
	}

	elephantRelease := d16Release{
		valves: valves,
		states: make(map[d16State]uint16, 100_000),
	}

	return release.releasePressure(valves.getID("AA"), 26, d16Opened{}, func(opened d16Opened) uint16 {
		// Go Go elephant helper!
		released := elephantRelease.releasePressure(valves.getID("AA"), 26, opened, nil)
		return released
	})
}
