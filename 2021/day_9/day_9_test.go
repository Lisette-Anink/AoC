package day_9

import (
	"reflect"
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
		{"test", args{useTestData()}, map[[2]int]int{}},
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

func Test_contains(t *testing.T) {
	type args struct {
		basin [][2]int
		pos   [2]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test true", args{[][2]int{{1, 2}, {3, 4}}, [2]int{3, 4}}, true},
		{"test false", args{[][2]int{{1, 2}, {3, 4}}, [2]int{2, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := contains(tt.args.basin, tt.args.pos)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findBasinNr(t *testing.T) {
	basins[1] = [][2]int{{1, 2}, {3, 4}}
	basins[2] = [][2]int{{10, 12}, {13, 14}}
	type args struct {
		pos [2]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{[2]int{3, 4}}, 1},
		{"test 0", args{[2]int{2, 4}}, 0},
		{"test 0", args{[2]int{13, 14}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findBasinNr(tt.args.pos)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_addToBasin(t *testing.T) {
	basins[1] = [][2]int{{1, 2}, {3, 4}}
	basins[2] = [][2]int{{10, 12}, {13, 14}}
	want1 := map[int][][2]int(map[int][][2]int{1: {{1, 2}, {3, 4}, {3, 5}}, 2: {{10, 12}, {13, 14}}})
	want2 := map[int][][2]int(map[int][][2]int{1: {{1, 2}, {3, 4}, {3, 5}}, 2: {{10, 12}, {13, 14}, {2, 4}}})
	want3 := map[int][][2]int(map[int][][2]int{1: {{1, 2}, {3, 4}, {3, 5}}, 2: {{10, 12}, {13, 14}, {2, 4}}, 3: {{2, 4}}})

	type args struct {
		pos [2]int
		nr  int
	}
	tests := []struct {
		name string
		args args
		want map[int][][2]int
	}{
		{"test 1", args{[2]int{3, 5}, 1}, want1},
		{"test 2", args{[2]int{2, 4}, 2}, want2},
		{"test 3", args{[2]int{2, 4}, 0}, want3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addToBasin(tt.args.pos, tt.args.nr)
			assert.Equal(t, tt.want, basins)
		})
	}
}

func Test_processPosition(t *testing.T) {
	floor := map[[2]int]int{{0, 0}: 1}
	wantFloor := map[[2]int]int{}
	basins[1] = [][2]int{{1, 2}, {3, 4}}
	basins[2] = [][2]int{{10, 12}, {13, 14}}
	want := map[int][][2]int(map[int][][2]int{1: {{1, 2}, {3, 4}, {0, 0}}, 2: {{10, 12}, {13, 14}}})
	type args struct {
		floormap map[[2]int]int
		pos      [2]int
		basinNr  int
	}
	tests := []struct {
		name string
		args args
		want map[int][][2]int
	}{
		{"test 1", args{floor, [2]int{0, 0}, 1}, want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processPosition(tt.args.floormap, tt.args.pos, tt.args.basinNr)
			assert.Equal(t, tt.want, basins)
			assert.Equal(t, wantFloor, floor)
		})
	}
}

func Test_checkBasinNeighbours(t *testing.T) {
	floor := map[[2]int]int{{0, 0}: 1, {0, 1}: 2}
	wantFloor := map[[2]int]int{}
	basins[1] = [][2]int{{1, 2}, {3, 4}}
	basins[2] = [][2]int{{10, 12}, {13, 14}}
	want := map[int][][2]int(map[int][][2]int{1: {{1, 2}, {3, 4}, {0, 1}, {0, 0}}, 2: {{10, 12}, {13, 14}}})
	type args struct {
		floormap map[[2]int]int
		pos      [2]int
		basinNr  int
	}
	tests := []struct {
		name string
		args args
		want map[int][][2]int
	}{
		{"test 1", args{floor, [2]int{0, 0}, 1}, want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkBasinNeighbours(tt.args.floormap, tt.args.pos, tt.args.basinNr)
			assert.Equal(t, tt.want, basins)
			assert.Equal(t, wantFloor, floor)
		})
	}
}
func Test_checkBasin(t *testing.T) {
	floor := map[[2]int]int{{0, 0}: 1, {0, 1}: 2}
	floor2 := map[[2]int]int{{0, 0}: 1, {0, 1}: 2, {0, 2}: 9, {0, 3}: 2}
	floor3 := map[[2]int]int{{0, 0}: 1, {0, 1}: 2, {0, 2}: 9, {0, 3}: 2}
	floor4 := map[[2]int]int{{0, 0}: 1, {0, 1}: 2, {0, 2}: 9, {0, 3}: 2}
	// wantFloor := map[[2]int]int{}
	want := map[int][][2]int(map[int][][2]int{1: {{0, 0}, {0, 1}}, 2: {{10, 12}, {13, 14}}})
	// want2 := map[int][][2]int(map[int][][2]int{1: {{0, 0}}, 2: {{10, 12}, {13, 14}}})
	want3 := map[int][][2]int(map[int][][2]int{1: {{0, 0}, {0, 1}}, 2: {{10, 12}, {13, 14}}, 3: {{0, 3}}})
	basins[1] = [][2]int{{0, 0}}
	basins[2] = [][2]int{{10, 12}, {13, 14}}

	type args struct {
		floormap map[[2]int]int
		pos      [2]int
	}
	tests := []struct {
		name string
		args args
		want map[int][][2]int
	}{
		{"test 1", args{floor, [2]int{0, 0}}, want},
		{"test 2", args{floor2, [2]int{0, 0}}, want},
		{"test 3", args{floor3, [2]int{0, 2}}, want},
		{"test 4", args{floor4, [2]int{0, 3}}, want3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.name)
			checkBasin(tt.args.floormap, tt.args.pos)
			assert.Equal(t, tt.want, basins)
			// assert.Equal(t, wantFloor, floor)
		})
	}
}

func Test_findBasin(t *testing.T) {
	test := parseInput(useTestData())
	type args struct {
		floormap map[[2]int]int
	}
	tests := []struct {
		name      string
		args      args
		wantCount map[int]int
	}{
		{"test", args{test}, map[int]int{1: 14, 2: 9, 3: 9, 4: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := findBasin(tt.args.floormap); !reflect.DeepEqual(gotCount, tt.wantCount) {
				t.Errorf("findBasin() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_solvePart2(t *testing.T) {
	test := findBasin(parseInput(useTestData()))
	real := findBasin(parseInput(useInputData()))
	type args struct {
		count map[int]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"solve all the test things", args{test}, 113},
		{"solve all the real things", args{real}, 113},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solvePart2(tt.args.count)
			assert.Equal(t, tt.want, got)
		})
	}
}
