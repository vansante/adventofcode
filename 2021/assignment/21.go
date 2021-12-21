package assignment

import "fmt"

type Day21 struct{}

const (
	d21BoardSize    = 10
	d21DiceRolls    = 3
	d21Pt1DiceSides = 100
	d21Pt2DiceSides = 3
	d21Pt1WinScore  = 1_000
	d21Pt2WinScore  = 21
)

type d21Die struct {
	rolls    int
	lastRoll int
}

func (d *d21Die) roll() int {
	d.rolls++
	d.lastRoll++
	if d.lastRoll > d21Pt1DiceSides {
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

	for players[0].score < d21Pt1WinScore && players[1].score < d21Pt1WinScore {
		for _, p := range players {
			roll := 0
			for i := 0; i < d21DiceRolls; i++ {
				roll += die.roll()
			}

			p.boardPos = (p.boardPos + roll) % d21BoardSize
			if p.boardPos == 0 {
				p.boardPos = d21BoardSize
			}
			p.score += int64(p.boardPos)
			if p.score >= d21Pt1WinScore {
				break
			}
		}
	}
	if players[0].score >= d21Pt1WinScore {
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

type d21GameState struct {
	score      [2]int
	position   [2]int8
	nextRoll   int8
	diceRolls  int8
	playerTurn int8
}

type d21Result [2]int64

type d21ResultMap map[d21GameState]d21Result

func (d *Day21) SolveI(input string) int64 {
	p1, p2 := d.getPlayers(input)

	return d.play([]*d21Player{p1, p2})
}

func (d *Day21) SolveII(input string) int64 {
	p1, p2 := d.getPlayers(input)

	states := make(d21ResultMap, 1024*1024)
	result := d.playDirac(states, d21GameState{
		position: [2]int8{int8(p1.startPos), int8(p2.startPos)},
	}, d21Pt2WinScore)

	if result[0] > result[1] {
		return result[0]
	}
	return result[1]
}

func (d *Day21) playDirac(states d21ResultMap, state d21GameState, winScore int) d21Result {
	result, ok := states[state]
	if ok {
		return result
	}

	st := state

	player := st.playerTurn
	st.position[player] = (st.position[player] + st.nextRoll) % d21BoardSize
	if st.position[player] == 0 {
		st.position[player] = d21BoardSize
	}

	if st.diceRolls == d21DiceRolls {
		st.score[player] += int(st.position[player])

		if st.score[player] >= winScore {
			res := d21Result{}
			res[player] = 1
			states[state] = res
			return res
		}
	}

	// No one won yet, play all games
	if st.diceRolls == d21DiceRolls {
		// Other players turn
		st.playerTurn += 1
		st.playerTurn %= 2

		st.diceRolls = 1
	} else {
		st.diceRolls++
	}

	for i := int8(0); i < d21Pt2DiceSides; i++ {
		st.nextRoll = i + 1
		res := d.playDirac(states, st, winScore)
		result[0] += res[0]
		result[1] += res[1]
	}
	states[state] = result
	return result
}
