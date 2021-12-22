package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type cube struct {
	x int
	y int
	z int
}

type activeCubes []cube

func (a activeCubes) operate(o opField) {
	cs := o.getCubes()
	log.Printf("Operation affects %d cubes", len(cs))
	for _, c := range cs {
		if o.turnOn {
			a.turnOn(c)
		} else {
			a.turnOff(c)
		}
	}
}

func (a activeCubes) turnOff(c cube) {
	for i, x := range a {
		if c == x {
			a[i] = a[len(a)-1]
			a = a[:len(a)-1]
		}
	}
}

func (a activeCubes) turnOn(c cube) {
	for _, x := range a {
		if c == x {
			return
		}
	}
	a = append(a, c)
}

func (a activeCubes) length() int {
	return len(a)
}

type opField struct {
	turnOn bool
	x      [2]int
	y      [2]int
	z      [2]int
}

func (o opField) getCubes() []cube {
	result := []cube{}
	for x := o.x[0]; x <= o.x[1]; x++ {
		for y := o.y[0]; y <= o.y[1]; y++ {
			for z := o.z[0]; z <= o.z[1]; z++ {
				result = append(
					result,
					cube{
						x: x,
						y: y,
						z: z,
					},
				)
			}
		}
	}
	return result
}

func main() {
	ops := getInput()
	reactor := activeCubes{}
	for _, op := range ops {
		log.Printf("Processing %+v", op)
		reactor.operate(op)
	}
	log.Printf("Result: %d", reactor.length())
}

func getInput() []opField {
	result := []opField{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		rawRanges := strings.Split(tokens[1], ",")
		result = append(
			result,
			opField{
				turnOn: tokens[0] == "on",
				x:      parseRange(rawRanges[0]),
				y:      parseRange(rawRanges[1]),
				z:      parseRange(rawRanges[2]),
			},
		)
	}
	return result
}

func parseRange(in string) [2]int {
	tokens := strings.Split(strings.Split(in, "=")[1], "..")
	min, err := strconv.Atoi(tokens[0])
	if err != nil {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(tokens[1])
	if err != nil {
		log.Fatal(err)
	}

	return [2]int{min, max}
}
