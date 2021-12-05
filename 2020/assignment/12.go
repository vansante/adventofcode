package assignment

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day12 struct{}

type d12Instruction struct {
	action string
	value  int
}

func (d *Day12) retrieveInstructions(in string) []d12Instruction {
	split := strings.Split(in, "\n")

	var input []d12Instruction
	for i := range split {
		line := strings.TrimSpace(split[i])
		if line == "" {
			continue
		}
		value, err := strconv.ParseInt(line[1:], 10, 32)
		CheckErr(err)
		input = append(input, d12Instruction{
			action: line[:1],
			value:  int(value),
		})
	}
	return input
}

type d12Ship struct {
	direction int
	n         int
	e         int
	waypointN int
	waypointE int
}

func (s *d12Ship) applyPtI(i d12Instruction) {
	switch strings.ToUpper(i.action) {
	case "N":
		s.n += i.value
	case "E":
		s.e += i.value
	case "S":
		s.n -= i.value
	case "W":
		s.e -= i.value
	case "L":
		s.direction = (s.direction - i.value) % 360
	case "R":
		s.direction = (s.direction + i.value) % 360
	case "F":
		switch s.direction {
		case 0:
			s.n += i.value
		case 90, -270:
			s.e += i.value
		case 180, -180:
			s.n -= i.value
		case 270, -90:
			s.e -= i.value
		default:
			panic(fmt.Sprintf("Moving forward in direction %d with value %d (%v)", s.direction, i.value, i))
		}
	default:
		panic(i.value)
	}
}

func (s *d12Ship) rotateWaypoint(value int) {
	if value < 0 {
		value += 360
	}
	switch value {
	case 0:
	case 90:
		s.waypointN, s.waypointE = -s.waypointE, s.waypointN
	case 180:
		s.waypointN, s.waypointE = -s.waypointN, -s.waypointE
	case 270:
		s.waypointN, s.waypointE = s.waypointE, -s.waypointN
	default:
		panic(fmt.Sprintf("Rotating left with value %d", value))
	}
}
func (s *d12Ship) applyPtII(i d12Instruction) {
	switch strings.ToUpper(i.action) {
	case "N":
		s.waypointN += i.value
	case "E":
		s.waypointE += i.value
	case "S":
		s.waypointN -= i.value
	case "W":
		s.waypointE -= i.value
	case "L":
		s.rotateWaypoint(-i.value)
	case "R":
		s.rotateWaypoint(i.value)
	case "F":
		s.n += s.waypointN * i.value
		s.e += s.waypointE * i.value
	default:
		panic(i.value)
	}
}

func (s d12Ship) printI() {
	fmt.Printf("Ship [N %d, E %d] facing %d\n", s.n, s.e, s.direction)
}

func (s d12Ship) printII() {
	fmt.Printf("Ship [N %d, E %d] waypoint [N %d, E %d]\n", s.n, s.e, s.waypointN, s.waypointE)
}

func (d *Day12) SolveI(input string) int64 {
	instr := d.retrieveInstructions(input)

	s := d12Ship{
		direction: 90,
	}

	for i := range instr {
		s.applyPtI(instr[i])
	}
	s.printI()

	return int64(math.Abs(float64(s.n)) + math.Abs(float64(s.e)))
}

func (d *Day12) SolveII(input string) int64 {
	instr := d.retrieveInstructions(input)

	s := d12Ship{
		waypointN: 1,
		waypointE: 10,
	}
	for i := range instr {
		s.applyPtII(instr[i])
	}
	s.printII()

	return int64(math.Abs(float64(s.n)) + math.Abs(float64(s.e)))
}
