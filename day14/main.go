package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func maxSubMin(in map[string]int) int {
	min, max := 0, 0
	for _, v := range in {
		if min == 0 {
			min = v
		}

		if max == 0 {
			max = v
		}

		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return max - min
}

func process(polymers *map[string]int, rules map[string]string, elements map[string]int) {
	newPolymers := make(map[string]int)
	for k, v := range *polymers {
		newElement := rules[k]
		elements[newElement] += v
		currentElements := strings.Split(k, "")
		newPolymers[currentElements[0]+newElement] += v
		newPolymers[newElement+currentElements[1]] += v
	}
	*polymers = newPolymers
}

func main() {
	chain, rules, elements := getInput()
	for i := 0; i < 40; i++ {
		process(&chain, rules, elements)
	}
	log.Printf("Result: %d", maxSubMin(elements))
}

func getInput() (map[string]int, map[string]string, map[string]int) {
	startPolymers := make(map[string]int)
	returnRules := make(map[string]string)
	startElements := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	startChain := strings.Split(scanner.Text(), "")
	for i := 2; i <= len(startChain); i++ {
		startPolymers[strings.Join(startChain[i-2:i], "")]++
	}
	for _, v := range startChain {
		startElements[v]++
	}
	for scanner.Scan() {
		if scanner.Text() != "" {
			tokens := strings.Split(scanner.Text(), " -> ")
			returnRules[tokens[0]] = tokens[1]
		}
	}

	return startPolymers, returnRules, startElements
}
