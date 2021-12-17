package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

type shot struct {
	vX       int
	vY       int
	maxY     int
	location coord
}

func (s *shot) tick() {
	if s.vX > 0 {
		s.vX--
	}
	if s.vX < 0 {
		s.vX++
	}
	s.vY--
	s.location.x += s.vX
	s.location.y += s.vY

	if s.location.y > s.maxY {
		s.maxY = s.location.y
	}
}

type targetArea struct {
	minX int
	maxX int
	minY int
	maxY int
}

func (t targetArea) isHit(c coord) string {
	if c.x >= t.minX && c.x <= t.maxX && c.y >= t.minY && c.y <= t.maxY {
		return "BULLSEYE"
	} else if c.y < t.minY {
		return "MISS"
	} else {
		return "UNKNOWN"
	}
}

func (t targetArea) bestX() (int, int) {
	minBest := math.MaxInt
	maxBest := 0
	for i := 1; sumOfIntegers(i) <= t.maxX; i++ {
		if sumOfIntegers(i) >= t.minX {
			if i < minBest {
				minBest = i
			}

			if i > maxBest {
				maxBest = i
			}
		}
	}

	return minBest, maxBest
}

func main() {
	target := getInput()
	log.Printf("%+v", target)
	min, max := target.bestX()
	log.Printf("%d, %d", min, max)
}

func getInput() targetArea {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rawInput := strings.Split(scanner.Text(), " ")[2:]
	rawX := strings.Split(rawInput[0][2:len(rawInput[0])-1], "..")
	rawY := strings.Split(rawInput[1][2:], "..")
	minX, err := strconv.Atoi(rawX[0])
	if err != nil {
		log.Fatal(err)
	}
	maxX, err := strconv.Atoi(rawX[1])
	if err != nil {
		log.Fatal(err)
	}
	minY, err := strconv.Atoi(rawY[0])
	if err != nil {
		log.Fatal(err)
	}
	maxY, err := strconv.Atoi(rawY[1])
	if err != nil {
		log.Fatal(err)
	}
	return targetArea{
		minX: minX,
		maxX: maxX,
		minY: minY,
		maxY: maxY,
	}
}

func sumOfIntegers(n int) int {
	return n * (n + 1) / 2
}
