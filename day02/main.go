package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type navigation struct {
	aim          int
	distance     int
	depth        int
	instructions *[]command
}

func (n *navigation) calculate() int {
	for _, v := range *n.instructions {
		n.process(v)
	}

	return n.distance * n.depth
}

func (n *navigation) process(c command) {
	switch c.operation {
	case "forward":
		n.distance += c.value
		n.depth += c.value * n.aim
	case "down":
		n.aim += c.value
	case "up":
		n.aim -= c.value
	}
}

type command struct {
	operation string
	value     int
}

type commands []command

func main() {
	log.Printf("Result: %d", getInput().calculate())
}

func getInput() *navigation {
	result := new(navigation)
	result.instructions = new([]command)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			tokens := strings.Split(scanner.Text(), " ")
			distance, err := strconv.Atoi(tokens[1])
			if err != nil {
				log.Fatal(err)
			}
			*result.instructions = append(*result.instructions, command{operation: tokens[0], value: distance})
		}
	}

	return result
}
