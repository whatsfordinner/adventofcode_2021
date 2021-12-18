package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type snailNumber interface {
	magnitude() int
	explode(int) (snailNumber, bool, int, int)
	shouldExplode(int) bool
	split() (bool, snailNumber)
	shouldSplit() bool
	isSingle() bool
	value() int
	add(int, string) snailNumber
}

type snailSingle int

func (s snailSingle) magnitude() int {
	return int(s)
}

func (s snailSingle) add(n int, direction string) snailNumber {
	return snailSingle(int(s) + n)
}

func (s snailSingle) isSingle() bool {
	return true
}

func (s snailSingle) shouldExplode(depth int) bool {
	return false
}

func (s snailSingle) explode(depth int) (snailNumber, bool, int, int) {
	return s, false, 0, 0
}

func (s snailSingle) shouldSplit() bool {
	if int(s) > 9 {
		return true
	}

	return false
}

func (s snailSingle) split() (bool, snailNumber) {
	if s.shouldSplit() {
		if int(s)%2 == 0 {
			return true, snailPair{
				snailSingle(int(s) / 2),
				snailSingle(int(s) / 2),
			}
		}

		return true, snailPair{
			snailSingle(int(s) / 2),
			snailSingle(int(s)/2 + 1),
		}
	}
	return false, s
}

func (s snailSingle) value() int {
	return int(s)
}

type snailPair [2]snailNumber

func (s snailPair) magnitude() int {
	return 3*s[0].magnitude() + 2*s[1].magnitude()
}

func (s snailPair) add(n int, direction string) snailNumber {
	if direction == "L" {
		s[0] = s[0].add(n, "L")
	} else {
		s[1] = s[1].add(n, "R")
	}

	return s
}

func (s snailPair) isSingle() bool {
	return false
}

func (s snailPair) shouldExplode(depth int) bool {
	if s[0].isSingle() && s[1].isSingle() && depth >= 4 {
		return true
	}

	return false
}

func (s snailPair) explode(depth int) (snailNumber, bool, int, int) {
	if s.shouldExplode(depth) {
		return snailSingle(0), true, s[0].value(), s[1].value()
	}

	result, leftExplode, left, right := s[0].explode(depth + 1)
	if leftExplode {
		s[1] = s[1].add(right, "L")
		return add(result, s[1]), leftExplode, left, 0
	}

	result, rightExplode, left, right := s[1].explode(depth + 1)
	if rightExplode {
		s[0] = s[0].add(left, "R")
		return add(s[0], result), rightExplode, 0, right
	}

	return s, false, 0, 0
}

func (s snailPair) shouldSplit() bool {
	return false
}

func (s snailPair) split() (bool, snailNumber) {
	didSplit, result := s[0].split()
	if didSplit {
		return true, add(result, s[1])
	}

	didSplit, result = s[1].split()
	if didSplit {
		return true, add(s[0], result)
	}

	return false, s
}

func (s snailPair) value() int {
	return 0
}

func main() {
	input := getInput()
	max := 0
	for i, v := range input {
		for j, x := range input {
			if i != j {
				result := add(v, x)
				operated := true
				for operated {
					didExplode := false
					split := false
					result, didExplode, _, _ = result.explode(0)

					if !didExplode {
						split, result = result.split()
					}

					operated = didExplode || split
				}

				if result.magnitude() > max {
					max = result.magnitude()
				}
			}
		}
	}

	log.Printf("Result: %d", max)

}

func getInput() []snailNumber {
	result := []snailNumber{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result = append(result, readNumber(scanner.Text()))
	}
	return result
}

func readNumber(in string) snailNumber {
	var result snailNumber
	split := strings.Split(in, "")

	if len(split) == 1 {
		v, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
		return snailSingle(v)
	}

	depth := 0
	for i, v := range split {
		if v == "[" {
			depth++
		}

		if v == "]" {
			depth--
		}

		if v == "," && depth == 1 {
			result = add(readNumber(in[1:i]), readNumber(in[i+1:len(in)-1]))
		}
	}

	return result
}

func add(left snailNumber, right snailNumber) snailNumber {
	return snailPair{
		left,
		right,
	}
}
