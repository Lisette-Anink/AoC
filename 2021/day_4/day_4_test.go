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

		assert.Equal(t, 4512, score)
	})
}

func Test_parseInput(t *testing.T) {
	lines := useTestData()
	wN := []string{"7", "4", "9", "5", "11", "17", "23", "2", "0", "14", "21", "24", "10", "16", "13", "6", "15", "25", "12", "22", "18", "20", "8", "19", "3", "26", "1"}
	wB := []board{}
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

func Test_Intersection(t *testing.T) {
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
			got := intersection(tt.inC, tt.inT)
			assert.Equal(t, tt.want, got)
		})
	}
}
