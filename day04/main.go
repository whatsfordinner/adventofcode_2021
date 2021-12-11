package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type bingoSquare struct {
	number int
	picked bool
}

type bingoBoard [][]bingoSquare

func (b *bingoBoard) checkRow(r int) bool {
	for i := 0; i < len((*b)[r]); i++ {
		if (*b)[r][i].picked == false {
			return false
		}
	}

	return true
}

func (b *bingoBoard) checkColumn(c int) bool {
	for i := 0; i < len(*b); i++ {
		if (*b)[i][c].picked == false {
			return false
		}
	}

	return true
}

func (b *bingoBoard) pickNumber(n int) (bool, int, int) {
	for i, r := range *b {
		for j, v := range r {
			if v.number == n {
				(*b)[i][j].picked = true
				return true, i, j
			}
		}
	}

	return false, 0, 0
}

func (b *bingoBoard) calculateScore(n int) int {
	score := 0
	for _, r := range *b {
		for _, v := range r {
			if v.picked == false {
				score += v.number
			}
		}
	}

	return score * n
}

type bingoGame struct {
	numbers []int
	boards  []*bingoBoard
}

func (b *bingoGame) drawNumber() (bool, *[]int, *[]int) {
	foundBoard := false
	foundBoardNumbers := new([]int)
	scores := new([]int)
	newNumber := b.numbers[0]
	b.numbers = b.numbers[1:]

	for i, v := range b.boards {
		found, row, column := v.pickNumber(newNumber)
		if found {
			if v.checkColumn(column) || v.checkRow(row) {
				foundBoard = true
				*foundBoardNumbers = append(*foundBoardNumbers, i)
				*scores = append(*scores, v.calculateScore(newNumber))
			}
		}
	}

	return foundBoard, foundBoardNumbers, scores
}

func main() {
	input := getInput()
	for len(input.numbers) > 0 {
		win, boards, scores := input.drawNumber()
		if win {
			if len(input.boards) > 1 {
				newBoards := new([]*bingoBoard)
				for i, v := range input.boards {
					addBoard := true
					for _, b := range *boards {
						if i == b {
							addBoard = false
						}
					}

					if addBoard {
						*newBoards = append(*newBoards, v)
					}
				}

				input.boards = *newBoards
			} else {
				log.Printf("Final winning board has score: %d", (*scores)[0])
				break
			}
		}
	}
}

func getInput() *bingoGame {
	result := new(bingoGame)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbersLine := strings.Split(scanner.Text(), ",")
	for _, v := range numbersLine {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		result.numbers = append(result.numbers, n)
	}

	scanner.Scan()

	newBoard := new(bingoBoard)
	for scanner.Scan() {
		if scanner.Text() == "" {
			result.boards = append(result.boards, newBoard)
			newBoard = new(bingoBoard)
		} else {
			newRow := new([]bingoSquare)
			for _, v := range strings.Split(scanner.Text(), " ") {
				if v != "" {
					n, err := strconv.Atoi(v)
					if err != nil {
						log.Fatal(err)
					}
					*newRow = append(*newRow, bingoSquare{number: n, picked: false})
				}
			}
			*newBoard = append(*newBoard, *newRow)
		}
	}

	result.boards = append(result.boards, newBoard)

	return result
}
