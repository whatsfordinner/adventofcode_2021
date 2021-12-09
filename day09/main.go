package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
	z int
}

func (p point) equals(o point) bool {
	return p.x == o.x && p.y == o.y && p.z == o.z
}

type cave struct {
	floor     [][]point
	lowpoints []point
	basins    [][]point
}

func (c *cave) findLowpoints() {
	for _, row := range c.floor {
		for _, col := range row {
			if c.isLowpoint(col) {
				c.lowpoints = append(c.lowpoints, col)
			}
		}
	}
}

func (c *cave) isLowpoint(p point) bool {
	if p.x != 0 && c.floor[p.y][p.x-1].z <= p.z {
		return false
	}

	if p.x != len(c.floor[p.y])-1 && c.floor[p.y][p.x+1].z <= p.z {
		return false
	}

	if p.y != 0 && c.floor[p.y-1][p.x].z <= p.z {
		return false
	}

	if p.y != len(c.floor)-1 && c.floor[p.y+1][p.x].z <= p.z {
		return false
	}

	return true
}

func (c *cave) totalRisk() int {
	risk := 0
	for _, v := range c.lowpoints {
		risk += v.z + 1
	}
	return risk
}

func (c *cave) findBasins() {
	for _, v := range c.lowpoints {
		c.basins = append(c.basins, c.trackBasin(v, []point{}))
	}
}

func (c *cave) trackBasin(p point, b []point) []point {
	b = append(b, p)

	if p.x != 0 && c.floor[p.y][p.x-1].z < 9 && !containsPoint(b, c.floor[p.y][p.x-1]) {
		b = c.trackBasin(c.floor[p.y][p.x-1], b)
	}

	if p.x != len(c.floor[p.y])-1 && c.floor[p.y][p.x+1].z < 9 && !containsPoint(b, c.floor[p.y][p.x+1]) {
		b = c.trackBasin(c.floor[p.y][p.x+1], b)
	}

	if p.y != 0 && c.floor[p.y-1][p.x].z < 9 && !containsPoint(b, c.floor[p.y-1][p.x]) {
		b = c.trackBasin(c.floor[p.y-1][p.x], b)
	}

	if p.y != len(c.floor)-1 && c.floor[p.y+1][p.x].z < 9 && !containsPoint(b, c.floor[p.y+1][p.x]) {
		b = c.trackBasin(c.floor[p.y+1][p.x], b)
	}

	return b
}

func main() {
	caves := getInput()
	caves.findLowpoints()
	caves.findBasins()
	sort.SliceStable(caves.basins, func(a int, b int) bool {
		return len(caves.basins[a]) > len(caves.basins[b])
	})
	log.Printf("Found %d basins", len(caves.basins))
	result := 1
	log.Printf("Three largest basins:")
	for i := 0; i < 3; i++ {
		log.Printf("\tbasin %d: size %d", i+1, len(caves.basins[i]))
		result *= len(caves.basins[i])
	}
	log.Printf("Result: %d", result)
}

func getInput() *cave {
	result := new(cave)
	result.floor = [][]point{}
	result.lowpoints = []point{}
	result.basins = [][]point{}
	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	for scanner.Scan() {
		newRow := []point{}
		for x, v := range strings.Split(scanner.Text(), "") {
			z, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			newRow = append(newRow, point{x, y, z})
		}
		result.floor = append(result.floor, newRow)
		y++
	}
	return result
}

func containsPoint(ps []point, p point) bool {
	for _, v := range ps {
		if v.equals(p) {
			return true
		}
	}

	return false
}
