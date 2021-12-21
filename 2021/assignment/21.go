package assignment

import "fmt"

type Day21 struct{}

type d21Die struct {
	rolls      int
	rolledOver int
	lastRoll   int
}

func (d *d21Die) roll() int {
	d.rolls++
	d.lastRoll++
	if d.lastRoll > 100 {
		d.rolledOver++
		d.lastRoll = 1
	}
	return d.lastRoll
}

type d21Player struct {
	id       int
	startPos int
	boardPos int
	score    int64
}

func (d *Day21) play(players []*d21Player) int64 {
	die := &d21Die{}
	players[0].boardPos = players[0].startPos
	players[1].boardPos = players[1].startPos

	for players[0].score < 1000 && players[1].score < 1000 {
		for _, p := range players {
			roll := die.roll() + die.roll() + die.roll()

			p.boardPos = (p.boardPos + roll) % 10
			if p.boardPos == 0 {
				p.boardPos = 10
			}
			p.score += int64(p.boardPos)
			if p.score >= 1000 {
				break
			}
		}
	}
	if players[0].score >= 1000 {
		return players[1].score * int64(die.rolls)
	}
	return players[0].score * int64(die.rolls)
}

func (d *Day21) getPlayers(input string) (*d21Player, *d21Player) {
	lines := SplitLines(input)

	p1 := &d21Player{}
	n, err := fmt.Sscanf(lines[0], "Player %d starting position: %d", &p1.id, &p1.startPos)
	CheckErr(err)
	if n != 2 {
		panic("invalid input")
	}

	p2 := &d21Player{}
	n, err = fmt.Sscanf(lines[1], "Player %d starting position: %d", &p2.id, &p2.startPos)
	CheckErr(err)
	if n != 2 {
		panic("invalid input")
	}
	return p1, p2
}

func (d *Day21) SolveI(input string) int64 {
	p1, p2 := d.getPlayers(input)

	return d.play([]*d21Player{p1, p2})
}

func (d *Day21) SolveII(input string) int64 {
	return 0
}
