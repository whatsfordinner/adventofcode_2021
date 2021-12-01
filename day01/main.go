package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type window []int

func (w *window) push(v int) {
	if len(*w) == 3 {
		*w = append((*w)[1:], v)
	} else {
		*w = append(*w, v)
	}
}

func (w *window) sum() int {
	result := 0

	for _, v := range *w {
		result += v
	}

	return result
}

func main() {
	inputFile := os.Args[1]
	input := getInput(inputFile)
	log.Printf("Result: %d", getIncreases(input))
}

func getInput(filename string) *[]int {
	result := new([]int)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			row, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}

			*result = append(*result, row)
		}
	}

	return result
}

func getIncreases(input *[]int) int {
	result := 0
	prev := 0
	w := new(window)

	for _, v := range *input {
		w.push(v)
		if len(*w) == 3 {
			if prev > 0 && w.sum() > prev {
				result++
			}
			prev = w.sum()
		}
	}

	return result
}
