package assignment

import "fmt"

type Day21 struct{}

const (
	d21Pt1DiceSides = 100
	d21Pt1WinScore  = 1_000
	d21Pt2WinScore  = 21
)

type d21Die struct {
	rolls      int
	rolledOver int
	lastRoll   int
}

func (d *d21Die) roll() int {
	d.rolls++
	d.lastRoll++
	if d.lastRoll > d21Pt1DiceSides {
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

	for players[0].score < d21Pt1WinScore && players[1].score < d21Pt1WinScore {
		for _, p := range players {
			roll := die.roll() + die.roll() + die.roll()

			p.boardPos = (p.boardPos + roll) % 10
			if p.boardPos == 0 {
				p.boardPos = 10
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

	states := make(d21ResultMap, 1024)
	result := d.playDirac(states, d21GameState{
		position:  [2]int8{int8(p1.startPos), int8(p2.startPos)},
		diceRolls: -1,
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
	st.position[player] = (st.position[player] + st.nextRoll) % 10
	if st.position[player] == 0 {
		st.position[player] = 10
	}

	if st.diceRolls == 2 {
		st.score[player] += int(st.position[player])

		if st.score[player] >= winScore {
			res := d21Result{}
			res[player] = 1
			states[state] = res
			return res
		}
	}

	// No one won yet, play all games
	var results [3]d21Result
	if st.diceRolls == 2 {
		// Other players turn
		st.playerTurn += 1
		st.playerTurn %= 2

		st.diceRolls = 0
		st.nextRoll = 1
		results[0] = d.playDirac(states, st, winScore)
		st.nextRoll = 2
		results[1] = d.playDirac(states, st, winScore)
		st.nextRoll = 3
		results[2] = d.playDirac(states, st, winScore)
	} else {
		st.nextRoll = 1
		st.diceRolls++
		results[0] = d.playDirac(states, st, winScore)
		st.nextRoll = 2
		results[1] = d.playDirac(states, st, winScore)
		st.nextRoll = 3
		results[2] = d.playDirac(states, st, winScore)
	}

	res := d21Result{}
	res[0] = results[0][0] + results[1][0] + results[2][0]
	res[1] = results[0][1] + results[1][1] + results[2][1]
	states[state] = res

	return res
}
