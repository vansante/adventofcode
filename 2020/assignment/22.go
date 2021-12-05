package assignment

import (
	"fmt"
	"strings"
)

type Day22 struct{}

type d22pPlayer struct {
	cards []int64
}

func (p *d22pPlayer) deckString() string {
	return fmt.Sprintf("%v", p.cards)
}

func (p *d22pPlayer) playCard() int64 {
	var card int64
	card, p.cards = p.cards[0], p.cards[1:]
	return card
}

func (p *d22pPlayer) addCard(card int64) {
	p.cards = append(p.cards, card)
}

func (p *d22pPlayer) copy(n int) *d22pPlayer {
	cp := &d22pPlayer{
		cards: make([]int64, n),
	}
	copy(cp.cards, p.cards)
	return cp
}

func (p *d22pPlayer) hasLost() bool {
	return len(p.cards) == 0
}

func (p *d22pPlayer) score() int64 {
	var score int64
	i := int64(len(p.cards))
	for _, c := range p.cards {
		score += i * c
		i--
	}
	return score
}

func (d *Day22) createPlayer(input string) *d22pPlayer {
	lines := SplitLines(input)
	numbers := MakeIntegers(lines[1:])

	p1 := &d22pPlayer{
		cards: numbers,
	}
	return p1
}

func (d *Day22) fetchPlayers(input string) (*d22pPlayer, *d22pPlayer) {
	players := strings.Split(input, "\n\n")

	return d.createPlayer(players[0]), d.createPlayer(players[1])
}

func (d *Day22) PlayRound(prevRounds map[string]struct{}, p1, p2 *d22pPlayer, recursive bool) (winner *d22pPlayer) {
	round := fmt.Sprintf("%s / %s", p1.deckString(), p2.deckString())
	if _, ok := prevRounds[round]; ok {
		fmt.Println("Same round as before; Player 1 wins")
		fmt.Println(round)
		return p1
	}
	prevRounds[round] = struct{}{}

	card1 := p1.playCard()
	card2 := p2.playCard()

	if recursive && int64(len(p1.cards)) >= card1 && int64(len(p2.cards)) >= card2 {
		// New rules; recursive
		p1C := p1.copy(int(card1))
		p2C := p2.copy(int(card2))

		prevRounds := make(map[string]struct{})
		for {
			winner := d.PlayRound(prevRounds, p1C, p2C, recursive)
			if winner == nil {
				continue
			}

			if winner == p1C {
				p1.addCard(card1)
				p1.addCard(card2)
				break
			} else if winner == p2C {
				p2.addCard(card2)
				p2.addCard(card1)
				break
			} else {
				panic("unknown winner")
			}
		}
	} else { // Normal game
		if card1 > card2 {
			p1.addCard(card1)
			p1.addCard(card2)
		} else if card1 < card2 {
			p2.addCard(card2)
			p2.addCard(card1)
		} else {
			panic("same card value!")
		}
	}

	if p1.hasLost() {
		fmt.Println("Player 2 cards: ", p2.cards)
		return p2
	}
	if p2.hasLost() {
		fmt.Println("Player 1 cards: ", p1.cards)
		return p1
	}

	return nil // No winner yet
}

func (d *Day22) SolveI(input string) int64 {
	p1, p2 := d.fetchPlayers(input)

	prevRounds := make(map[string]struct{})
	for {
		winner := d.PlayRound(prevRounds, p1, p2, false)
		if winner != nil {
			return winner.score()
		}
	}
}

func (d *Day22) SolveII(input string) int64 {
	p1, p2 := d.fetchPlayers(input)

	prevRounds := make(map[string]struct{})
	for {
		winner := d.PlayRound(prevRounds, p1, p2, true)
		if winner != nil {
			return winner.score()
		}
	}
}
