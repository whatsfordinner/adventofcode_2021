package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	log.Printf("%+v", input)
	for i := 1; i <= 256; i++ {
		input = tick(input)
		log.Printf("%+v", input)
	}

	numFish := 0
	for _, v := range input {
		numFish += v
	}

	log.Printf("Result: %d", numFish)
}

func getInput() [9]int {
	var result [9]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			ageStrings := strings.Split(scanner.Text(), ",")
			for _, v := range ageStrings {
				age, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}
				result[age]++
			}
		}
	}

	return result
}

func tick(input [9]int) [9]int {
	var output [9]int

	for i := 7; i >= 0; i-- {
		output[i] = input[i+1]
	}
	// newly spawned fish
	output[8] = input[0]
	output[6] += input[0]

	return output
}
