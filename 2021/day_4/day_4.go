package day_4

import (
	"strconv"
	"strings"

	"lisette.anink/aoc/utils"
)

type board struct {
	all  []string
	rows [][]string
	cols map[int][]string
}

func playBingo(numbers []string, boards []board) int {
	for i := 5; i < len(numbers); i++ {
		called := numbers[:i]
		for _, b := range boards {

			if b.bingo(called) {
				final, _ := strconv.Atoi(called[len(called)-1])
				return b.score(called) * final
			}
		}
	}
	return 0
}

func playBingoWinLast(numbers []string, boards []board) int {
	score := 0
	boardsMap := map[int]board{}
	for a, b := range boards {
		boardsMap[a] = b
	}

	for i := 5; i < len(numbers); i++ {
		called := numbers[:i]
		for x, b := range boardsMap {
			if b.bingo(called) {
				final, _ := strconv.Atoi(called[len(called)-1])
				score = b.score(called) * final
				delete(boardsMap, x)
			}
		}
	}
	return score
}

func parseInput(lines []string) ([]string, []board) {
	numbers := strings.Split(lines[0], ",")
	boards := []board{}
	for _, boa := range lines[1:] {
		b := board{}
		rows := strings.Split(boa, "\n")
		for _, row := range rows {
			if len(row) > 0 {
				vals := strings.Fields(row)
				b.rows = append(b.rows, vals)
				b.all = append(b.all, vals...)
			}
		}

		b.cols = utils.TransposeString(b.rows)
		boards = append(boards, b)
	}

	return numbers, boards
}

func (b *board) bingo(numbers []string) bool {
	for _, r := range b.rows {
		if includesAll(numbers, r) {
			return true
		}
	}
	for _, r := range b.cols {
		if includesAll(numbers, r) {
			return true
		}
	}
	return false
}

func (b *board) score(numbers []string) int {
	s := 0
	rest := difference(b.all, numbers)
	for _, n := range rest {
		i, _ := strconv.Atoi(n)
		s += i
	}
	return s
}

func includesAll(collection, test []string) bool {
	total := len(test)
	count := 0
	for i := 0; i < len(collection); i++ {
		for _, t := range test {
			if t == collection[i] {
				count++
				if count == total {
					return true
				}
			}
		}
	}
	return false
}

func difference(base, collection []string) []string {
	for _, s := range collection {
		for i, b := range base {
			if s == b {
				//remove element
				if i < len(base)-1 {
					base = append(base[:i], base[(i+1):]...)
				} else {
					base = base[:i]
				}
			}
		}
	}
	return base
}
