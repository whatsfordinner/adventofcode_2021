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

	log.Printf("Small caves: %+v", result)

	return result
}

func (c *caveSystem) tracePaths() [][]string {
	result := [][]string{}
	for _, v := range c.smallCaves() {
		interim := c.tracePath("start", []string{}, v)
		for _, x := range interim {
			result = append(result, x)
		}
	}

	unique := [][]string{}
	for _, v := range result {
		if !containsPath(unique, v) {
			unique = append(unique, v)
		}
	}

	return unique
}

func (c *caveSystem) tracePath(cave string, path []string, specialCave string) [][]string {
	result := [][]string{}
	path = append(path, cave)
	if cave == "end" {
		return append(result, path)
	}

	timesSpecialCaveVisited := 0
	for _, v := range path {
		if v == specialCave {
			timesSpecialCaveVisited++
		}
	}
	specialCaveOption := timesSpecialCaveVisited < 2

	options := []string{}
	for _, v := range (*c)[cave] {
		isOption := true
		if !isAllUpper(v) {
			for _, x := range path {
				if v == x {
					isOption = false
				}

				if v == specialCave && specialCaveOption {
					isOption = true
				}
			}
		}

		if isOption {
			options = append(options, v)
		}
	}

	for _, v := range options {
		interim := c.tracePath(v, path, specialCave)
		for _, x := range interim {
			result = append(result, x)
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

func containsPath(ps [][]string, p []string) bool {
	for _, v := range ps {
		if len(v) == len(p) {
			possibleMatch := true
			for i := 0; i < len(v); i++ {
				if v[i] != p[i] {
					possibleMatch = false
				}
			}
			if possibleMatch {
				return true
			}
		}
	}

	return false
}
