package day_6

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func useTestData() []string {
	input := utils.ImportFileLines("test_input_6")
	lines := utils.SplitInLines(input)
	return lines
}

func useInputData() []string {
	input := utils.ImportFromAoC("2021", "6")
	lines := utils.SplitInLines(input)
	return lines
}

func Test_Part1(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		lines := useTestData()
		pop := parseInput(lines[0])
		pop.simXDays(80)
		total := pop.totalSize()

		assert.Equal(t, 5934, total)
	})

	t.Run("real", func(t *testing.T) {
		lines := useInputData()
		pop := parseInput(lines[0])
		pop.simXDays(80)
		total := pop.totalSize()

		assert.Equal(t, 383160, total)
	})
}
func Test_Part2(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		lines := useTestData()
		pop := parseInput(lines[0])
		pop.simXDays(256)
		total := pop.totalSize()

		assert.Equal(t, 26984457539, total)
	})

	t.Run("real", func(t *testing.T) {
		lines := useInputData()
		pop := parseInput(lines[0])
		pop.simXDays(256)
		total := pop.totalSize()

		assert.Equal(t, 1721148811504, total)
	})
}

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want population
	}{
		{"case 1", "3,4,3,1,2", population{day: 0, sizePerAge: map[int]int{1: 1, 2: 1, 3: 2, 4: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseInput(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}
func Test_oneDay(t *testing.T) {
	tests := []struct {
		name string
		in   population
		want population
	}{
		{"case 1", population{day: 0, sizePerAge: map[int]int{1: 1, 2: 1, 3: 2, 4: 1}}, population{day: 1, sizePerAge: map[int]int{0: 1, 1: 1, 2: 2, 3: 1, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.oneDay()
			assert.Equal(t, tt.want, tt.in)
		})
	}
}
func Test_simXDays(t *testing.T) {
	tests := []struct {
		name string
		in   population
		x    int
		want population
	}{
		{"case 1", population{day: 0, sizePerAge: map[int]int{1: 1, 2: 1, 3: 2, 4: 1}}, 1, population{day: 1, sizePerAge: map[int]int{0: 1, 1: 1, 2: 2, 3: 1, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}}},
		{"case 2", population{day: 0, sizePerAge: map[int]int{1: 1, 2: 1, 3: 2, 4: 1}}, 2, population{day: 2, sizePerAge: map[int]int{0: 1, 1: 2, 2: 1, 3: 0, 4: 0, 5: 0, 6: 1, 7: 0, 8: 1}}},
		{"case 3", population{day: 0, sizePerAge: map[int]int{1: 1, 2: 1, 3: 2, 4: 1}}, 3, population{day: 3, sizePerAge: map[int]int{0: 2, 1: 1, 2: 0, 3: 0, 4: 0, 5: 1, 6: 1, 7: 1, 8: 1}}},
		{"case 4", population{day: 0, sizePerAge: map[int]int{1: 1, 2: 1, 3: 2, 4: 1}}, 4, population{day: 4, sizePerAge: map[int]int{0: 1, 1: 0, 2: 0, 3: 0, 4: 1, 5: 1, 6: 3, 7: 1, 8: 2}}},
		{"case 5", population{day: 0, sizePerAge: map[int]int{1: 1, 2: 1, 3: 2, 4: 1}}, 5, population{day: 5, sizePerAge: map[int]int{0: 0, 1: 0, 2: 0, 3: 1, 4: 1, 5: 3, 6: 2, 7: 2, 8: 1}}},
		{"case 10", population{day: 0, sizePerAge: map[int]int{1: 1, 2: 1, 3: 2, 4: 1}}, 10, population{day: 10, sizePerAge: map[int]int{0: 3, 1: 2, 2: 2, 3: 1, 4: 0, 5: 1, 6: 1, 7: 1, 8: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.simXDays(tt.x)
			assert.Equal(t, tt.want, tt.in)
		})
	}
}

func Test_totalSize(t *testing.T) {
	tests := []struct {
		name string
		in   population
		x    int
		want int
	}{
		{"case 1", population{day: 0, sizePerAge: map[int]int{1: 1, 2: 1, 3: 2, 4: 1}}, 1, 5},
		{"case 2", population{day: 0, sizePerAge: map[int]int{1: 1, 2: 1, 3: 2, 4: 1}}, 2, 6},
		{"case 10", population{day: 0, sizePerAge: map[int]int{1: 1, 2: 1, 3: 2, 4: 1}}, 10, 12},
		{"case 10", population{day: 0, sizePerAge: map[int]int{1: 1, 2: 1, 3: 2, 4: 1}}, 80, 5934},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.simXDays(tt.x)
			got := tt.in.totalSize()
			assert.Equal(t, tt.want, got)
		})
	}
}
