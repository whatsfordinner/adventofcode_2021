package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func (p *point) equals(o *point) bool {
	return p.x == o.x && p.y == o.y
}

type instruction struct {
	axis     string
	location int
}

type sheet struct {
	points       []*point
	instructions []*instruction
}

func (s *sheet) fold() bool {
	if len(s.instructions) == 0 {
		return false
	}

	ins := s.instructions[0]
	s.instructions = s.instructions[1:]

	newPoints := []*point{}
	for _, v := range s.points {
		nowOverlaps := false
		foldedPoint := new(point)
		if ins.axis == "y" {
			if v.y > ins.location {
				foldedPoint.x = v.x
				foldedPoint.y = v.y - (v.y-ins.location)*2
				for _, x := range s.points {
					if x.y < ins.location && foldedPoint.equals(x) {
						nowOverlaps = true
						break
					}
				}
			} else {
				foldedPoint = v
			}
		} else {
			if v.x > ins.location {
				foldedPoint.x = v.x - (v.x-ins.location)*2
				foldedPoint.y = v.y
				for _, x := range s.points {
					if x.x < ins.location && foldedPoint.equals(x) {
						nowOverlaps = true
						break
					}
				}
			} else {
				foldedPoint = v
			}
		}

		if !nowOverlaps {
			doAppend := true
			for _, v := range newPoints {
				if v.equals(foldedPoint) {
					doAppend = false
				}
			}

			if doAppend {
				newPoints = append(newPoints, foldedPoint)
			}
		}
	}

	s.points = newPoints
	return true
}

func (s *sheet) display() {
	maxX := 0
	maxY := 0

	for _, v := range s.points {
		if v.x > maxX {
			maxX = v.x
		}

		if v.y > maxY {
			maxY = v.y
		}
	}

	field := make([][]bool, maxY+1)
	for i := 0; i <= maxY; i++ {
		field[i] = make([]bool, maxX+1)
	}

	for _, v := range s.points {
		field[v.y][v.x] = true
	}

	for _, r := range field {
		row := ""
		for _, c := range r {
			if c {
				row += "#"
			} else {
				row += "."
			}
		}
		log.Printf("%s", row)
	}
}

func main() {
	input := getInput()
	input.fold()
	log.Printf("Points after 1 fold: %d", len(input.points))
	for input.fold() {
	}
	log.Printf("Total points: %d", len(input.points))
	input.display()
}

func getInput() *sheet {
	result := new(sheet)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		tokens := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(tokens[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal(err)
		}

		result.points = append(result.points, &point{x: x, y: y})
	}

	for scanner.Scan() {
		if scanner.Text() != "" {
			rawInstruction := strings.Split(scanner.Text(), " ")[2]
			tokens := strings.Split(rawInstruction, "=")
			axis := tokens[0]
			location, err := strconv.Atoi(tokens[1])
			if err != nil {
				log.Fatal(err)
			}

			result.instructions = append(result.instructions, &instruction{axis: axis, location: location})
		}
	}

	return result
}
