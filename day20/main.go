package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
)

type image struct {
	canvas   [][]bool
	alg      []bool
	newLight bool
}

func (i *image) enhance() {
	newHeight := len(i.canvas) + 2
	newWidth := len(i.canvas[0]) + 2
	newCanvas := make([][]bool, newHeight)
	for j := range newCanvas {
		newCanvas[j] = make([]bool, newWidth)
	}

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			algIndex := binToDec(i.getLit(x-1, y-1))
			if i.alg[algIndex] {
				newCanvas[y][x] = true
			} else {
				newCanvas[y][x] = false
			}
		}
	}

	if i.alg[0] {
		i.newLight = !i.newLight
	}

	i.canvas = newCanvas
}

func (i *image) getLit(x int, y int) []int {
	result := []int{}
	for dY := -1; dY <= 1; dY++ {
		for dX := -1; dX <= 1; dX++ {
			if x+dX < 0 || x+dX >= len(i.canvas[0]) || y+dY < 0 || y+dY >= len(i.canvas) {
				if i.newLight {
					result = append(result, 1)
				} else {
					result = append(result, 0)
				}
			} else {
				if i.canvas[y+dY][x+dX] {
					result = append(result, 1)
				} else {
					result = append(result, 0)
				}
			}
		}
	}
	return result
}

func (i *image) printImage() string {
	result := ""
	for _, r := range i.canvas {
		for _, c := range r {
			if c {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}

	return result
}

func (i *image) countLitPixels() int {
	result := 0
	for _, r := range i.canvas {
		for _, c := range r {
			if c {
				result++
			}
		}
	}

	return result
}

func main() {
	input := getInput()
	for i := 0; i < 50; i++ {
		input.enhance()
	}
	log.Printf("Lit pixels: %d", input.countLitPixels())
}

func getInput() *image {
	result := new(image)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	alg := []bool{}
	for _, v := range strings.Split(scanner.Text(), "") {
		if v == "#" {
			alg = append(alg, true)
		} else {
			alg = append(alg, false)
		}
	}
	result.alg = alg
	result.newLight = false

	scanner.Scan()

	for scanner.Scan() {
		newLine := []bool{}
		for _, v := range strings.Split(scanner.Text(), "") {
			if v == "#" {
				newLine = append(newLine, true)
			} else {
				newLine = append(newLine, false)
			}
		}
		result.canvas = append(result.canvas, newLine)
	}
	return result
}

func binToDec(in []int) int {
	result := 0
	for i := 0; i < len(in); i++ {
		result += in[len(in)-i-1] * int(math.Pow(2, float64(i)))
	}
	return result
}
