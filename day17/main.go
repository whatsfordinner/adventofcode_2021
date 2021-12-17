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
	s.location.x += s.vX
	s.location.y += s.vY

	if s.location.y > s.maxY {
		s.maxY = s.location.y
	}
	if s.vX > 0 {
		s.vX--
	}
	if s.vX < 0 {
		s.vX++
	}
	s.vY--
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

func (t targetArea) firstX() int {
	firstX := math.MaxInt
	for i := 1; sumOfIntegers(i) <= t.maxX; i++ {
		if sumOfIntegers(i) >= t.minX {
			if i < firstX {
				firstX = i
			}
		}
	}

	return firstX
}

func (t targetArea) lastX() int {
	return t.maxX
}

func (t targetArea) firstY() int {
	return t.minY
}

func (t targetArea) lastY() int {
	return -(t.minY + 1)
}

func main() {
	target := getInput()
	log.Printf("x: %d - %d", target.firstX(), target.lastX())
	log.Printf("y: %d - %d", target.firstY(), target.lastY())
	log.Printf("Candidates: %d", (target.lastX()-target.firstX())*(target.lastY()-target.firstY()))
	hits := 0
	for x := target.firstX(); x <= target.lastX(); x++ {
		for y := target.firstY(); y <= target.lastY(); y++ {
			testShot := &shot{
				vX:       x,
				vY:       y,
				location: coord{x: 0, y: 0},
			}
			var testResult string
			for !(testResult == "MISS" || testResult == "BULLSEYE") {
				testShot.tick()
				testResult = target.isHit(testShot.location)

			}
			if testResult == "BULLSEYE" {
				hits++
			}
		}
	}
	log.Printf("Result: %d", hits)
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
