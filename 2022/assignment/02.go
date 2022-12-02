package assignment

import (
	"strings"

	"github.com/vansante/adventofcode/2022/util"
)

type Day02 struct{}

type d02Weapon uint8

const (
	D02Rock     d02Weapon = 1
	D02Paper    d02Weapon = 2
	D02Scissors d02Weapon = 3
)

func (w d02Weapon) points() int64 {
	return int64(w)
}

func (w d02Weapon) needsWeapon(result d02Weapon) d02Weapon {
	switch result {
	case D02Rock: // We need to lose
		switch w {
		case D02Rock:
			return D02Scissors
		case D02Paper:
			return D02Rock
		case D02Scissors:
			return D02Paper
		}
	case D02Paper: // We need to draw
		switch w {
		case D02Rock:
			return D02Rock
		case D02Paper:
			return D02Paper
		case D02Scissors:
			return D02Scissors
		}
	case D02Scissors: // We need to win
		switch w {
		case D02Rock:
			return D02Paper
		case D02Paper:
			return D02Scissors
		case D02Scissors:
			return D02Rock
		}
	}
	panic("invalid values")
}

func (d *Day02) getWeapon(weapon string) d02Weapon {
	switch weapon {
	case "A", "X":
		return D02Rock
	case "B", "Y":
		return D02Paper
	case "C", "Z":
		return D02Scissors
	}
	panic("invalid value")
}

func (d *Day02) gamePoints(them, you d02Weapon) int64 {
	if them == you {
		return 3 // draw
	}

	switch {
	case you == D02Rock && them == D02Scissors,
		you == D02Paper && them == D02Rock,
		you == D02Scissors && them == D02Paper:
		return 6 // win
	}

	return 0 // loss
}

func (d *Day02) getGames(input string) [][]string {
	lines := util.SplitLines(input)
	games := make([][]string, len(lines))
	for i := range lines {
		games[i] = strings.Split(lines[i], " ")
	}
	return games
}

func (d *Day02) SolveI(input string) int64 {
	games := d.getGames(input)
	score := int64(0)
	for i := range games {
		them := d.getWeapon(games[i][0])
		you := d.getWeapon(games[i][1])
		score += you.points() + d.gamePoints(them, you)
	}
	return score
}

func (d *Day02) SolveII(input string) int64 {
	games := d.getGames(input)
	score := int64(0)
	for i := range games {
		them := d.getWeapon(games[i][0])
		outcome := d.getWeapon(games[i][1])
		you := them.needsWeapon(outcome)
		score += you.points() + d.gamePoints(them, you)
	}
	return score
}
