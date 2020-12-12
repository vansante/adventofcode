package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type instruction struct {
	action string
	value  int
}

func retrieveInstructions(file string) []instruction {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(content), "\n")

	var input []instruction
	for i := range split {
		line := strings.TrimSpace(split[i])
		if line == "" {
			continue
		}
		value, err := strconv.ParseInt(line[1:], 10, 32)
		if err != nil {
			panic(err)
		}
		input = append(input, instruction{
			action: line[:1],
			value:  int(value),
		})
	}
	return input
}

type ship struct {
	direction int
	n         int
	e         int
	waypointN int
	waypointE int
}

func (s *ship) applyPtI(i instruction) {
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

func (s *ship) rotateWaypoint(value int) {
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
func (s *ship) applyPtII(i instruction) {
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

func (s ship) printI() {
	fmt.Printf("Ship [N %d, E %d] facing %d\n", s.n, s.e, s.direction)
}

func (s ship) printII() {
	fmt.Printf("Ship [N %d, E %d] waypoint [N %d, E %d]\n", s.n, s.e, s.waypointN, s.waypointE)
}

func main() {
	wd, _ := os.Getwd()
	instr := retrieveInstructions(filepath.Join(wd, "12/input.txt"))

	s := ship{
		direction: 90,
	}

	for i := range instr {
		s.applyPtI(instr[i])
		//s.printI()
	}

	fmt.Printf("Part I: Coordinates: [N %d, E %d] Distance: %d\n\n", s.n, s.e, int(math.Abs(float64(s.n))+math.Abs(float64(s.e))))

	s = ship{
		waypointN: 1,
		waypointE: 10,
	}
	for i := range instr {
		s.applyPtII(instr[i])
		//s.printII()
	}
	fmt.Printf("Part II: Coordinates: [N %d, E %d] Distance: %d\n\n", s.n, s.e, int(math.Abs(float64(s.n))+math.Abs(float64(s.e))))
}
