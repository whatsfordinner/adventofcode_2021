package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

type sevenSegment map[string]string

type display string

type console struct {
	displays       *[]*display
	reading        *[]*display
	mapping        *map[int]display
	reverseMapping *map[display]int
}

func (s *sevenSegment) value(in *display) int {
	return 0
}

//  aaaa
// b    c
// b    c
//  dddd
// e    f
// e    f
//  gggg

func (c *console) mapConsole() {
	// map known knowns
	for _, v := range *c.displays {
		if len(*v) == 2 {
			(*c.mapping)[1] = *v
			(*c.reverseMapping)[*v] = 1
		}

		if len(*v) == 4 {
			(*c.mapping)[4] = *v
			(*c.reverseMapping)[*v] = 4
		}

		if len(*v) == 3 {
			(*c.mapping)[7] = *v
			(*c.reverseMapping)[*v] = 7
		}

		if len(*v) == 7 {
			(*c.mapping)[8] = *v
			(*c.reverseMapping)[*v] = 8
		}
	}

	// 3 is the only 5 segment number with 'c' and 'f'
	for _, v := range *c.displays {
		if len(*v) == 5 && len(commonChars((*c.mapping)[1], *v)) == 2 {
			(*c.mapping)[3] = *v
			(*c.reverseMapping)[*v] = 3
			break
		}
	}

	// 9 is the only 6 segment number one different to 3
	for _, v := range *c.displays {
		if len(*v) == 6 && len(uniqueChars((*c.mapping)[3], *v)) == 1 {
			(*c.mapping)[9] = *v
			(*c.reverseMapping)[*v] = 9
			break
		}
	}

	// 6 is the only 6 segment number with only one segment in common with 1
	for _, v := range *c.displays {
		if (len(*v)) == 6 && len(commonChars((*c.mapping)[1], *v)) == 1 {
			(*c.mapping)[6] = *v
			(*c.reverseMapping)[*v] = 6
		}
	}

	// 0 is the only 6 segment number left
	for _, v := range *c.displays {
		if (len(*v)) == 6 && *v != (*c.mapping)[9] && *v != (*c.mapping)[6] {
			(*c.mapping)[0] = *v
			(*c.reverseMapping)[*v] = 0
			break
		}
	}

	// 5 is only 1 segment different to 6
	for _, v := range *c.displays {
		if (len(*v)) == 5 && len(uniqueChars((*c.mapping)[6], *v)) == 1 {
			(*c.mapping)[5] = *v
			(*c.reverseMapping)[*v] = 5
			break
		}
	}

	// 2 is the only 5 segment number left unidentified
	for _, v := range *c.displays {
		if (len(*v)) == 5 && *v != (*c.mapping)[5] && *v != (*c.mapping)[3] {
			(*c.mapping)[2] = *v
			(*c.reverseMapping)[*v] = 2
			break
		}
	}
}

func (c *console) value() int {
	result := 0

	result += 1000 * (*c.reverseMapping)[*(*c.reading)[0]]
	result += 100 * (*c.reverseMapping)[*(*c.reading)[1]]
	result += 10 * (*c.reverseMapping)[*(*c.reading)[2]]
	result += 1 * (*c.reverseMapping)[*(*c.reading)[3]]

	return result
}

func main() {
	input := getInput()
	result := 0

	for _, v := range *input {
		v.mapConsole()
		result += v.value()
	}

	log.Printf("Result: %d", result)
}

func getInput() *[]console {
	result := new([]console)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			newConsole := new(console)
			newConsole.displays = new([]*display)
			newConsole.reading = new([]*display)
			newConsole.mapping = new(map[int]display)
			*(newConsole.mapping) = map[int]display{}
			newConsole.reverseMapping = new(map[display]int)
			*(newConsole.reverseMapping) = map[display]int{}
			tokens := strings.Split(scanner.Text(), " | ")
			displays := strings.Split(tokens[0], " ")
			reading := strings.Split(tokens[1], " ")
			for _, v := range displays {
				x := strings.Split(v, "")
				sort.Strings(x)
				newDisplay := display(strings.Join(x, ""))
				*newConsole.displays = append(*newConsole.displays, &newDisplay)
			}
			for _, v := range reading {
				x := strings.Split(v, "")
				sort.Strings(x)
				newReading := display(strings.Join(x, ""))
				*newConsole.reading = append(*newConsole.reading, &newReading)
			}
			*result = append(*result, *newConsole)
		}
	}

	return result
}

func uniqueChars(a display, b display) string {
	result := ""

	for _, v := range a {
		if !strings.ContainsRune(string(b), v) {
			result += string(v)
		}
	}

	for _, v := range b {
		if !strings.ContainsRune(string(a), v) {
			result += string(v)
		}
	}

	return result
}

func commonChars(a display, b display) string {
	result := ""

	for _, v := range a {
		for _, x := range b {
			if x == v {
				result += string(v)
				break
			}
		}
	}

	return result
}
