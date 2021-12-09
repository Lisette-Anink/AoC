package day_9

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func useTestData() []string {
	input := utils.ImportFileLines("test_input_9")
	lines := utils.SplitInLines(input)
	return lines
}

func useInputData() []string {
	input := utils.ImportFromAoC("2021", "9")
	lines := utils.SplitInLines(input)
	return lines
}

func Test_parseInput(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want map[[2]int]int
	}{
		{"smal", args{[]string{"123", "456"}}, map[[2]int]int{{0, 0}: 1, {0, 1}: 2, {0, 2}: 3, {1, 0}: 4, {1, 1}: 5, {1, 2}: 6}},
		// {"test", args{useTestData()}, map[[2]int]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseInput(tt.args.lines)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findLowPoints(t *testing.T) {
	floor := parseInput(useTestData())
	floorI := parseInput(useInputData())
	type args struct {
		floormap map[[2]int]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{floor}, 15},
		{"test", args{floorI}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findLowPoints(tt.args.floormap)
			assert.Equal(t, tt.want, got)
		})
	}
}
