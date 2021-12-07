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

type line struct {
	start coord
	end   coord
}

func (l line) getCoords() []coord {
	var result []coord
	var length int

	if l.isVertical() {
		length = int(math.Abs(float64(l.end.y-l.start.y))) + 1
	} else {
		length = int(math.Abs(float64(l.end.x-l.start.x))) + 1
	}

	dX := (l.end.x - l.start.x) / (length - 1)
	dY := (l.end.y - l.start.y) / (length - 1)

	for i := 0; i < length; i++ {
		result = append(result, coord{x: l.start.x + (dX * i), y: l.start.y + (dY * i)})
	}

	return result
}

func (l line) isVertical() bool {
	return l.start.x == l.end.x
}

func (l line) isHorizontal() bool {
	return l.start.y == l.end.y
}

func main() {
	inputFile := os.Args[1]
	input := getInput(inputFile)
	seenCoords := map[coord]int{}
	overlappingPoints := 0
	for _, v := range *input {
		for _, c := range v.getCoords() {
			seenCoords[c]++
		}
	}

	for _, v := range seenCoords {
		if v > 1 {
			overlappingPoints++
		}
	}
	log.Printf("%+v", overlappingPoints)
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
