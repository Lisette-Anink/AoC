package day_11

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func useTestData() []string {
	input := utils.ImportFileLines("test_input_11")
	lines := utils.SplitInLines(input)
	return lines
}

func useInputData() []string {
	input := utils.ImportFromAoC("2021", "11")
	lines := utils.SplitInLines(input)
	return lines
}

var small_octomap = map[[2]int]int{{0, 0}: 1, {0, 1}: 1, {0, 2}: 1, {0, 3}: 1, {0, 4}: 1, {1, 0}: 1, {1, 1}: 9, {1, 2}: 9, {1, 3}: 9, {1, 4}: 1, {2, 0}: 1, {2, 1}: 9, {2, 2}: 1, {2, 3}: 9, {2, 4}: 1, {3, 0}: 1, {3, 1}: 9, {3, 2}: 9, {3, 3}: 9, {3, 4}: 1, {4, 0}: 1, {4, 1}: 1, {4, 2}: 1, {4, 3}: 1, {4, 4}: 1}

func Test_octopusEnergyModel(t *testing.T) {
	// octomap := utils.ParseIntMap(useTestData())
	type args struct {
		octomap map[[2]int]int
		steps   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"small test", args{small_octomap, 1}, 9},
		{"bigger 1 test", args{utils.ParseIntMap(useTestData()), 1}, 0},
		{"bigger 10 test", args{utils.ParseIntMap(useTestData()), 10}, 204},
		{"bigger 100 test", args{utils.ParseIntMap(useTestData()), 100}, 1656},
		{"bigger 200 part 2", args{utils.ParseIntMap(useTestData()), 200}, 1656},
		{"part 1", args{utils.ParseIntMap(useInputData()), 100}, 1656},
		{"part 2", args{utils.ParseIntMap(useInputData()), 10000}, 1656},
	}
	// assert.Equal(t, small_octomap, octomap)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := octopusEnergyModel(tt.args.octomap, tt.args.steps)
			assert.Equal(t, tt.want, got)
		})
	}
}
