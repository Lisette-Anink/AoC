package day_4

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func useTestData() []string {
	input := utils.ImportFileLines("test_input_4")
	lines := utils.SplitInMultiLines(input)
	return lines
}

func useInputData() []string {
	input := utils.ImportFromAoC("2021", "4")
	lines := utils.SplitInMultiLines(input)
	return lines
}

func Test_playBingo(t *testing.T) {
	lines := useTestData()

	t.Run("play", func(t *testing.T) {
		score := playBingo(parseInput(lines))

		assert.Equal(t, 4512, score)
	})
}

func Test_Part1(t *testing.T) {
	lines := useInputData()

	t.Run("play", func(t *testing.T) {
		score := playBingo(parseInput(lines))

		assert.Equal(t, 63552, score)
	})
}
func Test_playBingoWinLast(t *testing.T) {
	lines := useTestData()

	t.Run("play", func(t *testing.T) {
		score := playBingoWinLast(parseInput(lines))

		assert.Equal(t, 192, score)
	})
}
func Test_Part2(t *testing.T) {
	lines := useInputData()

	t.Run("play", func(t *testing.T) {
		score := playBingoWinLast(parseInput(lines))

		assert.Equal(t, 9020, score)
	})
}

var parsedTestInput = []board{{all: []string{"22", "13", "17", "11", "0", "8", "2", "23", "4", "24", "21", "9", "14", "16", "7", "6", "10", "3", "18", "5", "1", "12", "20", "15", "19"}, rows: [][]string{{"22", "13", "17", "11", "0"}, {"8", "2", "23", "4", "24"}, {"21", "9", "14", "16", "7"}, {"6", "10", "3", "18", "5"}, {"1", "12", "20", "15", "19"}}, cols: map[int][]string{0: {"22", "8", "21", "6", "1"}, 1: {"13", "2", "9", "10", "12"}, 2: {"17", "23", "14", "3", "20"}, 3: {"11", "4", "16", "18", "15"}, 4: {"0", "24", "7", "5", "19"}}}, {all: []string{"3", "15", "0", "2", "22", "9", "18", "13", "17", "5", "19", "8", "7", "25", "23", "20", "11", "10", "24", "4", "14", "21", "16", "12", "6"}, rows: [][]string{{"3", "15", "0", "2", "22"}, {"9", "18", "13", "17", "5"}, {"19", "8", "7", "25", "23"}, {"20", "11", "10", "24", "4"}, {"14", "21", "16", "12", "6"}}, cols: map[int][]string{0: {"3", "9", "19", "20", "14"}, 1: {"15", "18", "8", "11", "21"}, 2: {"0", "13", "7", "10", "16"}, 3: {"2", "17", "25", "24", "12"}, 4: {"22", "5", "23", "4", "6"}}}, {all: []string{"14", "21", "17", "24", "4", "10", "16", "15", "9", "19", "18", "8", "23", "26", "20", "22", "11", "13", "6", "5", "2", "0", "12", "3", "7"}, rows: [][]string{{"14", "21", "17", "24", "4"}, {"10", "16", "15", "9", "19"}, {"18", "8", "23", "26", "20"}, {"22", "11", "13", "6", "5"}, {"2", "0", "12", "3", "7"}}, cols: map[int][]string{0: {"14", "10", "18", "22", "2"}, 1: {"21", "16", "8", "11", "0"}, 2: {"17", "15", "23", "13", "12"}, 3: {"24", "9", "26", "6", "3"}, 4: {"4", "19", "20", "5", "7"}}}}

func Test_parseInput(t *testing.T) {
	lines := useTestData()
	wN := []string{"7", "4", "9", "5", "11", "17", "23", "2", "0", "14", "21", "24", "10", "16", "13", "6", "15", "25", "12", "22", "18", "20", "8", "19", "3", "26", "1"}
	wB := parsedTestInput
	t.Run("play", func(t *testing.T) {
		numbers, boards := parseInput(lines)

		assert.Equal(t, wN, numbers)
		assert.Equal(t, wB, boards)
	})
}
func Test_IncludesAll(t *testing.T) {
	tests := []struct {
		name string
		inC  []string
		inT  []string
		want bool
	}{
		{"case 1", []string{"a", "b", "c"}, []string{"a", "b"}, true},
		{"case 2", []string{"a", "b", "c"}, []string{"a", "b", "c", "d"}, false},
		{"case 3", []string{"a", "b", "c"}, []string{"a", "d"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := includesAll(tt.inC, tt.inT)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_Difference(t *testing.T) {
	tests := []struct {
		name string
		inC  []string
		inT  []string
		want []string
	}{
		{"case 1", []string{"a", "b", "c"}, []string{"a", "b"}, []string{"c"}},
		{"case 2", []string{"a", "b", "c"}, []string{"a", "b", "c", "d"}, []string{}},
		{"case 3", []string{"a", "b", "c"}, []string{"a", "d"}, []string{"b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := difference(tt.inC, tt.inT)
			assert.Equal(t, tt.want, got)
		})
	}
}
