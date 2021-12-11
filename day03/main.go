package main

import (
	"bufio"
	"log"
	"math"
	"os"
)

func main() {
	input := getInput()
	gamma, epsilon := findGammaAndEpsilon(input)
	log.Printf("Gamme and Epsilon: %d", binaryToDecimal(gamma)*binaryToDecimal(epsilon))
	oxygen := findMeetsCriteria(0, input, pickMostCommon)
	co2 := findMeetsCriteria(0, input, pickLeastCommon)
	log.Printf("Oxygen and CO2: %d", binaryToDecimal(&oxygen)*binaryToDecimal(&co2))
}

func getInput() *[][]int {
	result := new([][]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			newLine := new([]int)
			for _, v := range scanner.Text() {
				bit := int(v - '0')
				*newLine = append(*newLine, bit)
			}
			*result = append(*result, *newLine)
		}
	}

	return result
}

func findGammaAndEpsilon(in *[][]int) (*[]int, *[]int) {
	gamma := new([]int)
	epsilon := new([]int)
	numRows := len(*in)
	numCols := len((*in)[0])

	for i := 0; i < numCols; i++ {
		count := 0
		for j := 0; j < numRows; j++ {
			if (*in)[j][i] == 1 {
				count++
			}
		}
		if count > numRows/2 {
			*gamma = append(*gamma, 1)
			*epsilon = append(*epsilon, 0)
		} else {
			*gamma = append(*gamma, 0)
			*epsilon = append(*epsilon, 1)
		}
	}

	return gamma, epsilon
}

func binaryToDecimal(in *[]int) int {
	result := 0

	for i := 0; i < len(*in); i++ {
		result += (*in)[len(*in)-i-1] * int(math.Pow(2, float64(i)))
	}

	return result
}

func findMeetsCriteria(index int, in *[][]int, decider func(*[][]int, *[][]int) *[][]int) []int {
	if len(*in) == 1 {
		return (*in)[0]
	} else {
		zeros := new([][]int)
		ones := new([][]int)

		for i := 0; i < len(*in); i++ {
			if (*in)[i][index] == 1 {
				*ones = append(*ones, (*in)[i])
			} else {
				*zeros = append(*zeros, (*in)[i])
			}
		}

		return findMeetsCriteria(index+1, decider(ones, zeros), decider)
	}
}

func pickMostCommon(ones *[][]int, zeros *[][]int) *[][]int {
	if len(*ones) >= len(*zeros) {
		return ones
	}

	return zeros
}

func pickLeastCommon(ones *[][]int, zeros *[][]int) *[][]int {
	if len(*ones) < len(*zeros) {
		return ones
	}

	return zeros
}
