package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type cell struct {
	risk int
	link []*cell
}

type cave struct {
	start *cell
	end   *cell
	cells []*cell
}

func (c *cave) dijkstra() int {
	dist := make(map[*cell]int)
	prev := make(map[*cell]*cell)
	unvisited := make([]*cell, len(c.cells))
	copy(unvisited, c.cells)

	for _, v := range c.cells {
		dist[v] = math.MaxInt
		prev[v] = nil
	}

	dist[c.start] = 0

	for len(unvisited) > 0 {
		current := minCell(dist, unvisited)
		removeCell(&unvisited, current)

		if current == c.end {
			break
		}

		for _, v := range current.link {
			if containsCell(unvisited, v) {
				possibleRisk := dist[current] + v.risk
				if possibleRisk < dist[v] {
					dist[v] = possibleRisk
					prev[v] = current
				}
			}
		}
	}

	return dist[c.end]
}

func minCell(dist map[*cell]int, unvisited []*cell) *cell {
	minV := math.MaxInt
	var minC *cell
	for _, v := range unvisited {
		if dist[v] < minV {
			minC = v
			minV = dist[v]
		}
	}

	return minC
}

func removeCell(in *[]*cell, del *cell) {
	for i, v := range *in {
		if v == del {
			*in = append((*in)[:i], (*in)[i+1:]...)
			break
		}
	}
}

func containsCell(in []*cell, search *cell) bool {
	for _, v := range in {
		if v == search {
			return true
		}
	}
	return false
}

func main() {
	input := getInput()
	log.Printf("Result: %d", input.dijkstra())
}

func getInput() *cave {
	result := new(cave)
	input := [][]int{}
	allCells := [][]*cell{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			newRow := []int{}
			for _, v := range strings.Split(scanner.Text(), "") {
				r, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}
				newRow = append(newRow, r)
			}
			input = append(input, newRow)
		}
	}

	for y, row := range input {
		newCellRow := []*cell{}
		for x, c := range row {
			newCell := new(cell)
			newCell.risk = c
			if y == 0 && x == 0 {
				result.start = newCell
			}

			if y == len(input)-1 && x == len(input[y])-1 {
				result.end = newCell
			}

			newCellRow = append(newCellRow, newCell)
		}
		allCells = append(allCells, newCellRow)
	}

	for y, row := range allCells {
		for x, c := range row {
			if x+1 < len(row) {
				c.link = append(c.link, allCells[y][x+1])
			}

			if x-1 >= 0 {
				c.link = append(c.link, allCells[y][x-1])
			}

			if y+1 < len(allCells) {
				c.link = append(c.link, allCells[y+1][x])
			}

			if y-1 >= 0 {
				c.link = append(c.link, allCells[y-1][x])
			}
		}
	}

	for _, v := range allCells {
		result.cells = append(result.cells, v...)
	}

	log.Printf("Cave start: %v", result.start)
	log.Printf("Cave end: %v", result.end)
	log.Printf("Cell at (%d,%d): %v", len(allCells), len(allCells[len(allCells)-1]), allCells[len(allCells)-1][len(allCells[len(allCells)-1])-1])

	return result
}
