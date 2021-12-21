package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type pawn struct {
	position int
	score    int
}

func (p pawn) move(n int) pawn {
	p.position = (p.position + n) % 10
	p.score += p.position + 1
	return p
}

type game struct {
	p1 pawn
	p2 pawn
}

func (g game) playTurn(d map[int]int, p int) games {
	result := make(games)
	if g.winner() != 0 {
		result[g] = 1
		return result
	}

	for k, v := range d {
		if p == 1 {
			newP1 := g.p1.move(k)
			result[game{p1: newP1, p2: g.p2}] += v
		} else {
			newP2 := g.p2.move(k)
			result[game{p1: g.p1, p2: newP2}] += v
		}
	}

	return result
}

func (g game) winner() int {
	if g.p1.score >= 21 {
		return 1
	} else if g.p2.score >= 21 {
		return 2
	}
	return 0
}

type games map[game]int

func (g *games) playTurn(d map[int]int, p int) {
	result := make(games)

	for k, v := range *g {
		for r, c := range k.playTurn(d, p) {
			result[r] += v * c
		}
	}

	*g = result
}

func main() {
	g := new(games)
	*g = make(games)
	(*g)[getInput()] = 1
	diceDistribution := make(map[int]int)
	for d1 := 1; d1 <= 3; d1++ {
		for d2 := 1; d2 <= 3; d2++ {
			for d3 := 1; d3 <= 3; d3++ {
				diceDistribution[d1+d2+d3]++
			}
		}
	}

	finished := false
	player := 0
	for !finished {
		g.playTurn(diceDistribution, player+1)
		finished = true
		for k := range *g {
			if k.winner() == 0 {
				finished = false
				break
			}
		}
		player = (player + 1) % 2
	}

	p1Wins := 0
	p2Wins := 0
	for k, v := range *g {
		if k.winner() == 1 {
			p1Wins += v
		} else if k.winner() == 2 {
			p2Wins += v
		} else {
			log.Printf("Something has gone horribly wrong...")
		}
	}

	log.Printf("P1 won %d times", p1Wins)
	log.Printf("P2 won %d times", p2Wins)
}

func getInput() game {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	p1Raw := strings.TrimSpace(strings.Split(scanner.Text(), ":")[1])
	p1Posn, err := strconv.Atoi(p1Raw)
	if err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	p2Raw := strings.TrimSpace(strings.Split(scanner.Text(), ":")[1])
	p2Posn, err := strconv.Atoi(p2Raw)
	if err != nil {
		log.Fatal(err)
	}
	return game{
		p1: pawn{position: p1Posn - 1, score: 0},
		p2: pawn{position: p2Posn - 1, score: 0},
	}
}
