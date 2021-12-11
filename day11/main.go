package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type dumboOctopus struct {
	energy  int
	flashes int
}

func (d *dumboOctopus) incrementEnergy() {
	d.energy++
}

func (d *dumboOctopus) doesFlash() bool {
	if d.energy == 10 {
		d.flashes++
		return true
	}

	return false
}

func (d *dumboOctopus) finalise() bool {
	if d.energy > 9 {
		d.energy = 0
		return true
	}

	return false
}

func main() {
	input := getInput()
	for steps := 1; ; steps++ {
		for i, r := range *input {
			for j := range r {
				tick(j, i, input)
			}
		}

		synchronised := true

		for _, r := range *input {
			for _, c := range r {
				if !c.finalise() {
					synchronised = false
				}
			}
		}

		if synchronised {
			log.Printf("Synchronised flash on step %d", steps)
			break
		}
	}

	flashes := 0
	for _, r := range *input {
		for _, c := range r {
			flashes += c.flashes
		}
	}

	log.Printf("Result: %d", flashes)
}

func getInput() *[][]*dumboOctopus {
	result := new([][]*dumboOctopus)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			newRow := []*dumboOctopus{}
			tokens := strings.Split(scanner.Text(), "")
			for _, v := range tokens {
				startingEnergy, err := strconv.Atoi(string(v))
				if err != nil {
					log.Fatal(err)
				}
				newRow = append(newRow, &dumboOctopus{energy: startingEnergy, flashes: 0})
			}
			*result = append(*result, newRow)
		}
	}

	return result
}

func tick(x int, y int, in *[][]*dumboOctopus) {
	(*in)[y][x].incrementEnergy()
	if (*in)[y][x].doesFlash() {
		for dX := -1; dX <= 1; dX++ {
			if x+dX < 0 || x+dX >= len((*in)[y]) {
				continue
			}
			for dY := -1; dY <= 1; dY++ {
				if y+dY < 0 || y+dY >= len((*in)) {
					continue
				}

				tick(x+dX, y+dY, in)
			}
		}
	}
}
