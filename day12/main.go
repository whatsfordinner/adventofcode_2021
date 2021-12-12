package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type caveSystem map[string][]string

func (c *caveSystem) connect(a string, b string) {
	aConnections, aPresent := (*c)[a]
	if !aPresent {
		(*c)[a] = []string{b}
	} else {
		found := false
		for _, v := range aConnections {
			if v == b {
				found = true
				break
			}
		}

		if !found {
			(*c)[a] = append((*c)[a], b)
		}
	}
}

func (c *caveSystem) smallCaves() []string {
	result := []string{}
	for k := range *c {
		if k != "start" && k != "end" && !isAllUpper(k) {
			result = append(result, k)
		}
	}

	return result
}

func (c *caveSystem) tracePaths() [][]string {
	return c.tracePath("start", []string{}, false)
}

func (c *caveSystem) tracePath(cave string, path []string, revisited bool) [][]string {
	path = append(path, cave)
	if cave == "end" {
		return [][]string{path}
	}

	result := [][]string{}
	for _, v := range (*c)[cave] {
		isOption := true
		if !isAllUpper(v) {
			for _, x := range path {
				if v == x && revisited || v == "start" {
					isOption = false
					break
				}
			}
		}

		if isOption {
			newRevisited := revisited

			if !newRevisited {
				for _, x := range path {
					if !isAllUpper(v) && v == x {
						newRevisited = true
					}
				}
			}

			copyPath := make([]string, len(path))
			copy(copyPath, path)
			result = append(result, c.tracePath(v, copyPath, newRevisited)...)
		}
	}

	return result
}

func main() {
	input := getInput()
	result := input.tracePaths()
	log.Printf("Total distinct paths: %d", len(result))
}

func getInput() *caveSystem {
	result := make(caveSystem)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			tokens := strings.Split(scanner.Text(), "-")
			result.connect(tokens[0], tokens[1])
			result.connect(tokens[1], tokens[0])
		}
	}

	return &result
}

func isAllUpper(s string) bool {
	re := regexp.MustCompile(`^[A-Z]+$`)
	return re.MatchString(s)
}
