package day_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func useTestData() []string {
	input := utils.ImportFileLines("test_input_5")
	lines := utils.SplitInLines(input)
	return lines
}

func useInputData() []string {
	input := utils.ImportFromAoC("2021", "5")
	lines := utils.SplitInLines(input)
	return lines
}

func Test_Part1(t *testing.T) {
	lines := useTestData()

	t.Run("test", func(t *testing.T) {
		score := countHighVent(calcDiagram(parseInput(lines)))

		assert.Equal(t, 5, score)
	})
	lines = useInputData()

	t.Run("real", func(t *testing.T) {
		score := countHighVent(calcDiagram(parseInput(lines)))

		assert.Equal(t, 5, score)
	})
}
func Test_Part2(t *testing.T) {
	lines := useTestData()

	t.Run("test", func(t *testing.T) {
		// diagram := calcDiagramDia(parseInput(lines))

		// assert.Equal(t, 12, diagram)
		score := countHighVent(calcDiagramDia(parseInput(lines)))

		assert.Equal(t, 12, score)
	})
	lines = useInputData()

	t.Run("real", func(t *testing.T) {
		score := countHighVent(calcDiagramDia(parseInput(lines)))

		assert.Equal(t, 5, score)
	})
}

func Test_parse(t *testing.T) {
	lines := useTestData()
	want := []vector([]vector{{x1: 0, y1: 9, x2: 5, y2: 9}, {x1: 8, y1: 0, x2: 0, y2: 8}, {x1: 9, y1: 4, x2: 3, y2: 4}, {x1: 2, y1: 2, x2: 2, y2: 1}, {x1: 7, y1: 0, x2: 7, y2: 4}, {x1: 6, y1: 4, x2: 2, y2: 0}, {x1: 0, y1: 9, x2: 2, y2: 9}, {x1: 3, y1: 4, x2: 1, y2: 4}, {x1: 0, y1: 0, x2: 8, y2: 8}, {x1: 5, y1: 5, x2: 8, y2: 2}})
	t.Run("play", func(t *testing.T) {
		vectors := parseInput(lines)

		assert.Equal(t, want, vectors)
	})
}

// calcDiagram

func Test_calcDiagram(t *testing.T) {
	tests := []struct {
		name string
		in   []vector
		want map[[2]int]int
	}{
		{"case 1", []vector{{1, 2, 1, 5}}, map[[2]int]int{[2]int{1, 2}: 1, [2]int{1, 3}: 1, [2]int{1, 4}: 1, [2]int{1, 5}: 1}},
		{"case backwards", []vector{{2, 2, 2, 1}}, map[[2]int]int{{2, 1}: 1, {2, 2}: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcDiagram(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}
func Test_calcDiagramDia(t *testing.T) {
	tests := []struct {
		name string
		in   []vector
		want map[[2]int]int
	}{
		{"case 1", []vector{{1, 2, 1, 5}}, map[[2]int]int{{1, 2}: 1, {1, 3}: 1, {1, 4}: 1, {1, 5}: 1}},
		{"case backwards", []vector{{2, 2, 2, 1}}, map[[2]int]int{{2, 1}: 1, {2, 2}: 1}},

		{"case dia 1", []vector{{8, 0, 0, 8}}, map[[2]int]int{{0, 8}: 1, {1, 7}: 1, {2, 6}: 1, {3, 5}: 1, {4, 4}: 1, {5, 3}: 1, {6, 2}: 1, {7, 1}: 1, {8, 0}: 1}},

		{"case dia 2", []vector{{6, 4, 2, 0}}, map[[2]int]int{{2, 0}: 1, {3, 1}: 1, {4, 2}: 1, {5, 3}: 1, {6, 4}: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcDiagramDia(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}
func Test_countHighVent(t *testing.T) {
	tests := []struct {
		name string
		in   map[[2]int]int
		want int
	}{
		{"case 1", map[[2]int]int{{1, 2}: 4, {1, 3}: 1, {1, 4}: 10, {1, 5}: 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countHighVent(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want []vector
	}{
		{"case 1", []string{"1,2 -> 1,5"}, []vector{{1, 2, 1, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseInput(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}
