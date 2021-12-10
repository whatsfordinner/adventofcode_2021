package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

type opener rune

func (o opener) matches(r rune) bool {
	return o == '(' && r == ')' ||
		o == '[' && r == ']' ||
		o == '{' && r == '}' ||
		o == '<' && r == '>'
}

type navStack []opener

func (s *navStack) push(r opener) {
	*s = append(*s, r)
}

func (s *navStack) pop() opener {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return result
}

func (s *navStack) isEmpty() bool {
	return len(*s) == 0
}

func main() {
	input := getInput()
	incompleteScores := []int{}
	for _, v := range *input {
		corrupt, stack := isCorrupted(v)
		if !corrupt {
			incompleteScores = append(incompleteScores, scoreIncomplete(v, stack))
		}
	}
	sort.Ints(incompleteScores)
	log.Printf("Result: %d", incompleteScores[len(incompleteScores)/2])
}

func scoreCorrupted(in string) int {
	stack := new(navStack)
	for _, v := range in {
		if v == '(' || v == '[' || v == '{' || v == '<' {
			stack.push(opener(v))
		} else {
			currentOpener := stack.pop()
			if !currentOpener.matches(v) {
				switch v {
				case ')':
					return 3
				case ']':
					return 57
				case '}':
					return 1197
				case '>':
					return 25137
				}
			}
		}
	}

	return 0
}

func isCorrupted(in string) (bool, *navStack) {
	stack := new(navStack)
	for _, v := range in {
		if v == '(' || v == '[' || v == '{' || v == '<' {
			stack.push(opener(v))
		} else {
			currentOpener := stack.pop()
			if !currentOpener.matches(v) {
				return true, stack
			}
		}
	}

	return false, stack
}

func scoreIncomplete(in string, stack *navStack) int {
	result := 0

	for !stack.isEmpty() {
		result *= 5
		switch stack.pop() {
		case '(':
			result += 1
		case '[':
			result += 2
		case '{':
			result += 3
		case '<':
			result += 4
		}
	}

	return result
}

func getInput() *[]string {
	result := new([]string)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			*result = append(*result, scanner.Text())
		}
	}

	return result
}
