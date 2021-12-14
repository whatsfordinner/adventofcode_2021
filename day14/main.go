package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func maxSubMin(in string) int {
	result := map[string]int{}
	for _, v := range strings.Split(string(in), "") {
		result[v]++
	}

	min, max := 0, 0

	for _, v := range result {
		if min == 0 {
			min = v
		}

		if max == 0 {
			max = v
		}

		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return max - min
}

func process(rs *rules, in string, firstProcess bool) string {
	test := strings.Split(in, "")
	quickValue := (*rs)[in]
	if quickValue != "" {
		if firstProcess {
			return test[0] + quickValue + test[len(test)-1]
		} else {
			return quickValue + test[len(test)-1]
		}
	} else {
		result := ""
		if len(in) < 1000 {
			if firstProcess {
				result += test[0]
			}
			for i := 1; i < len(test); i++ {
				result += (*rs)[test[i-1]+test[i]] + test[i]
			}
		} else {
			splits := 30
			increment := len(in) / splits
			i := increment
			result += process(rs, in[:i], firstProcess)
			i += increment
			for ; i < len(in); i += increment {
				result += process(rs, in[i-increment-1:i], false)
			}
			result += process(rs, in[i-increment-1:], false)
		}

		if firstProcess {
			(*rs)[in] = result[1 : len(result)-1]
		} else {
			(*rs)[in] = result[:len(result)-1]
		}
		return result
	}
}

type rules map[string]string

func main() {
	chain, rules := getInput()
	for i := 0; i < 40; i++ {
		log.Printf("Iteration %d", i+1)
		chain = process(rules, chain, true)
	}
	log.Printf("Result: %d", maxSubMin(chain))
}

func getInput() (string, *rules) {
	returnRules := make(rules)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	returnChain := scanner.Text()
	for scanner.Scan() {
		if scanner.Text() != "" {
			tokens := strings.Split(scanner.Text(), " -> ")
			returnRules[tokens[0]] = tokens[1]
		}
	}

	return returnChain, &returnRules
}
