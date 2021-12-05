package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

func (c coord) equals(other coord) bool {
	return c.x == other.x && c.y == other.y
}

type line struct {
	start coord
	end   coord
}

func (l line) getCoords() []coord {
	var result []coord

	if l.isVertical() {
		if l.start.y <= l.end.y {
			for i := l.start.y; i <= l.end.y; i++ {
				result = append(result, coord{x: l.start.x, y: i})
			}
		} else {
			for i := l.end.y; i <= l.start.y; i++ {
				result = append(result, coord{x: l.start.x, y: i})
			}
		}
	} else if l.isHorizontal() {
		if l.start.x <= l.end.x {
			for i := l.start.x; i <= l.end.x; i++ {
				result = append(result, coord{x: i, y: l.start.y})
			}
		} else {
			for i := l.end.x; i <= l.start.x; i++ {
				result = append(result, coord{x: i, y: l.start.y})
			}
		}
	} else {
		if l.start.x < l.end.x && l.start.y < l.end.y {
			points := l.end.x - l.start.x
			for i := 0; i <= points; i++ {
				result = append(result, coord{x: l.start.x + i, y: l.start.y + i})
			}
		} else if l.start.x > l.end.x && l.start.y < l.end.y {
			points := l.start.x - l.end.x
			for i := 0; i <= points; i++ {
				result = append(result, coord{x: l.start.x - i, y: l.start.y + i})
			}
		} else if l.start.x < l.end.x && l.start.y > l.end.y {
			points := l.end.x - l.start.x
			for i := 0; i <= points; i++ {
				result = append(result, coord{x: l.start.x + i, y: l.start.y - i})
			}
		} else {
			points := l.start.x - l.end.x
			for i := 0; i <= points; i++ {
				result = append(result, coord{x: l.start.x - i, y: l.start.y - i})
			}
		}
	}

	return result
}

func (l line) isVertical() bool {
	return l.start.x == l.end.x
}

func (l line) isHorizontal() bool {
	return l.start.y == l.end.y
}

func (l line) calculateOverlap(other line) []coord {
	var result []coord

	for _, i := range l.getCoords() {
		for _, j := range other.getCoords() {
			if i.equals(j) {
				result = append(result, i)
			}
		}
	}

	return result
}

func main() {
	inputFile := os.Args[1]
	input := getInput(inputFile)
	var overlap []coord
	for i, v := range *input {
		for _, x := range (*input)[i+1:] {
			newOverlap := v.calculateOverlap(x)
			overlap = appendNew(overlap, newOverlap)
		}
	}
	log.Printf("%+v", len(overlap))
}

func getInput(filename string) *[]line {
	result := new([]line)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			coordStrings := strings.Split(scanner.Text(), " -> ")
			startStrings := strings.Split(coordStrings[0], ",")
			endStrings := strings.Split(coordStrings[1], ",")

			x1, err := strconv.Atoi(startStrings[0])
			if err != nil {
				log.Fatal(err)
			}
			y1, err := strconv.Atoi(startStrings[1])
			if err != nil {
				log.Fatal(err)
			}
			x2, err := strconv.Atoi(endStrings[0])
			if err != nil {
				log.Fatal(err)
			}
			y2, err := strconv.Atoi(endStrings[1])
			if err != nil {
				log.Fatal(err)
			}

			*result = append(
				*result,
				line{
					start: coord{
						x: x1,
						y: y1,
					},
					end: coord{
						x: x2,
						y: y2,
					},
				},
			)
		}
	}

	return result
}

func appendNew(slice []coord, in []coord) []coord {
	for _, i := range in {
		found := false
		for _, j := range slice {
			if i.equals(j) {
				found = true
				break
			}
		}

		if !found {
			slice = append(slice, i)
		}
	}

	return slice
}
