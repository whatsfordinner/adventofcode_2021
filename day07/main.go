package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type crabArmada []int

func (c *crabArmada) converge(x int) int {
	fuelUsed := 0
	for _, v := range *c {
		steps := int(math.Abs(float64(v - x)))
		fuelUsed += sumToOne(steps)
	}

	return fuelUsed
}

func (c *crabArmada) leftmost() int {
	leftmost := (*c)[0]

	for _, v := range *c {
		if v < leftmost {
			leftmost = v
		}
	}

	return leftmost
}

func (c *crabArmada) rightmost() int {
	rightmost := (*c)[0]

	for _, v := range *c {
		if v > rightmost {
			rightmost = v
		}
	}

	return rightmost
}

func (c *crabArmada) quickConverge() (int, int) {
	position := c.leftmost()
	fuel := c.converge(position)

	for i := c.leftmost() + 1; i <= c.rightmost(); i++ {
		newFuel := c.converge(i)
		if newFuel < fuel {
			position = i
			fuel = newFuel
		}
	}

	return position, fuel
}

func main() {
	inputFile := os.Args[1]
	input := getInput(inputFile)
	position, fuel := input.quickConverge()
	log.Printf("Result: Converging on %d takes %d fuel", position, fuel)
}

func getInput(filename string) *crabArmada {
	result := new(crabArmada)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			tokens := strings.Split(scanner.Text(), ",")
			for _, v := range tokens {
				position, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}

				*result = append(*result, position)
			}
		}
	}

	return result
}

func sumToOne(in int) int {
	out := 0
	for i := 1; i <= in; i++ {
		out += i
	}

	return out
}
