package day_22

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func useTestData(s string) []string {
	input := utils.ImportFileLines("test_input_22")
	if s == "s" {
		input = utils.ImportFileLines("test_input_22_small")
	}
	if s == "b" {
		input = utils.ImportFileLines("test_input_22_big")
	}
	lines := utils.SplitInLines(input)
	return lines
}

func useInputData() []string {
	input := utils.ImportFromAoC("2021", "22")
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
		want map[int][7]int
	}{
		{"test 1", args{[]string{"on x=10..12,y=10..12,z=10..12", ""}}, map[int][7]int{0: {1, 10, 12, 10, 12, 10, 12}}},
		{"test 2", args{[]string{"on x=-27..23,y=-28..26,z=-21..29", ""}}, map[int][7]int{0: {1, -27, 23, -28, 26, -21, 29}}},
		{"test small", args{useTestData("s")}, map[int][7]int{0: {1, 10, 12, 10, 12, 10, 12}, 1: {1, 11, 13, 11, 13, 11, 13}, 2: {-1, 9, 11, 9, 11, 9, 11}, 3: {1, 10, 10, 10, 10, 10, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseInput(tt.args.lines)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_rebootInitiation(t *testing.T) {
	type args struct {
		rebootSteps map[int][7]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{map[int][7]int{0: {1, 10, 12, 10, 12, 10, 12}}}, 27},
		{"test small", args{parseInput(useTestData("s"))}, 39},
		{"test", args{parseInput(useTestData(""))}, 590784},
		{"test big", args{parseInput(useTestData("b"))}, 474140},
		{"Part 1", args{parseInput(useInputData())}, 644257},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rebootInitiation(tt.args.rebootSteps)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_inRange(t *testing.T) {
	type args struct {
		c   [3]int
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test", args{[3]int{10, 50, -50}, -50, 50}, true},
		{"test", args{[3]int{10, 51, -49}, -50, 50}, false},
		{"test", args{[3]int{-110, 50, -49}, -50, 50}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inRange(tt.args.c, tt.args.min, tt.args.max)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_reboot(t *testing.T) {
	type args struct {
		rebootSteps map[int][7]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{map[int][7]int{0: {1, 10, 12, 10, 12, 10, 12}}}, 27},
		{"test small", args{parseInput(useTestData("s"))}, 39},
		{"test", args{parseInput(useTestData(""))}, 590784},
		// {"test big", args{parseInput(useTestData("b"))}, 474140},
		// {"Part 1", args{parseInput(useInputData())}, 644257},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reboot(tt.args.rebootSteps)
			assert.Equal(t, tt.want, got)
		})
	}
}
