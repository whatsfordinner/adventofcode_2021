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

type priorityQueue struct {
	queue []*priorityObject
}

type priorityObject struct {
	cell     *cell
	priority int
}

func (q *priorityQueue) extract_min() *cell {
	if len(q.queue) == 0 {
		return nil
	}

	minP := math.MaxInt
	minI := 0
	for i, v := range q.queue {
		if v.priority < minP {
			minP = v.priority
			minI = i
		}
	}

	result := q.queue[minI]
	q.queue = append(q.queue[:minI], q.queue[minI+1:]...)

	return result.cell
}

func (q *priorityQueue) add_with_priority(c *cell, p int) {
	q.queue = append(q.queue, &priorityObject{c, p})
}

func (q *priorityQueue) decrease_priority(c *cell, p int) {
	for _, v := range q.queue {
		if v.cell == c {
			v.priority = p
		}
	}
}

func (c *cave) dijkstra() int {
	dist := make(map[*cell]int)
	prev := make(map[*cell]*cell)
	unvisited := new(priorityQueue)
	unvisited.queue = []*priorityObject{}

	for _, v := range c.cells {
		unvisited.add_with_priority(v, math.MaxInt)
		dist[v] = math.MaxInt
		prev[v] = nil
	}

	dist[c.start] = 0

	for len(unvisited.queue) > 0 {
		current := unvisited.extract_min()

		if current == c.end {
			break
		}

		for _, v := range current.link {
			possibleRisk := dist[current] + v.risk
			if possibleRisk < dist[v] {
				dist[v] = possibleRisk
				prev[v] = current
				unvisited.decrease_priority(v, possibleRisk)
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

	for i := 0; i < 5; i++ {
		for y, row := range input {
			newCellRow := []*cell{}
			for j := 0; j < 5; j++ {
				for x, c := range row {
					newCell := new(cell)
					newCell.risk = riskWrap(c + i + j)
					if y == 0 && x == 0 && i == 0 && j == 0 {
						result.start = newCell
					}

					if y == len(input)-1 && x == len(input[y])-1 && i == 4 && j == 4 {
						result.end = newCell
					}

					newCellRow = append(newCellRow, newCell)
				}
			}
			allCells = append(allCells, newCellRow)
		}
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

	return result
}

func riskWrap(in int) int {
	for in > 9 {
		in -= 9
	}

	return in
}
