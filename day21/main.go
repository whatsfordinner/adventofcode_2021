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

func (p *pawn) move(n int) {
	p.position = (p.position + n) % 10
	p.score += p.position + 1
}

type die struct {
	nextSide    int
	timesRolled int
	maxSide     int
}

func (d *die) rollN(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		result += d.nextSide + 1
		d.timesRolled++
		d.nextSide = (d.nextSide + 1) % 100
	}

	return result
}

type game struct {
	d  *die
	p1 *pawn
	p2 *pawn
}

func (g *game) playTurn() *pawn {
	g.p1.move(g.d.rollN(3))
	if g.p1.score >= 1000 {
		return g.p1
	}

	g.p2.move(g.d.rollN(3))
	if g.p2.score >= 1000 {
		return g.p2
	}

	return nil
}

func (g *game) winningScore(w *pawn) int {
	if w == g.p1 {
		return g.p2.score * g.d.timesRolled
	}
	return g.p1.score * g.d.timesRolled
}

func main() {
	g := getInput()
	var winner *pawn
	for winner == nil {
		winner = g.playTurn()
	}
	log.Printf("Result: %d", g.winningScore(winner))
}

func getInput() *game {
	result := new(game)
	result.d = &die{maxSide: 100}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	p1Raw := strings.TrimSpace(strings.Split(scanner.Text(), ":")[1])
	p1Posn, err := strconv.Atoi(p1Raw)
	if err != nil {
		log.Fatal(err)
	}
	result.p1 = &pawn{position: p1Posn - 1}
	scanner.Scan()
	p2Raw := strings.TrimSpace(strings.Split(scanner.Text(), ":")[1])
	p2Posn, err := strconv.Atoi(p2Raw)
	if err != nil {
		log.Fatal(err)
	}
	result.p2 = &pawn{position: p2Posn - 1}
	return result
}
