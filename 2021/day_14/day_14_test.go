package day_14

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func useTestData() []string {
	input := utils.ImportFileLines("test_input_14")
	parts := utils.SplitInMultiLines(input)
	return parts
}

func useInputData() []string {
	input := utils.ImportFromAoC("2021", "14")
	lines := utils.SplitInMultiLines(input)
	return lines
}

var insertRule = map[[2]string]string{{"B", "B"}: "N", {"B", "C"}: "B", {"B", "H"}: "H", {"B", "N"}: "B", {"C", "B"}: "H", {"C", "C"}: "N", {"C", "H"}: "B", {"C", "N"}: "C", {"H", "B"}: "C", {"H", "C"}: "B", {"H", "H"}: "N", {"H", "N"}: "C", {"N", "B"}: "B", {"N", "C"}: "B", {"N", "H"}: "C", {"N", "N"}: "C"}

func Test_parseInput(t *testing.T) {
	type args struct {
		parts []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 map[[2]string]string
	}{
		{"test small", args{[]string{"ABC", "AA -> B\nBC -> N"}}, []string{"A", "B", "C"}, map[[2]string]string{{"A", "A"}: "B", {"B", "C"}: "N"}},
		{"test test", args{useTestData()}, []string{"N", "N", "C", "B"}, insertRule},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseInput(tt.args.parts)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_pairInsertion(t *testing.T) {
	p, i := parseInput(useTestData())
	type args struct {
		polymer        []string
		insertionRules map[[2]string]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"small test", args{[]string{"A", "B", "C"}, map[[2]string]string{{"A", "B"}: "H", {"B", "C"}: "N"}}, []string{"A", "H", "B", "N", "C"}},
		{"small test", args{[]string{"A", "B", "C"}, map[[2]string]string{{"A", "B"}: "A", {"B", "C"}: "N"}}, []string{"A", "A", "B", "N", "C"}},
		{"bigger test", args{p, i}, []string{"N", "C", "N", "B", "C", "H", "B"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pairInsertion(tt.args.polymer, tt.args.insertionRules)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_insertSteps(t *testing.T) {
	p, i := parseInput(useTestData())
	type args struct {
		polymer        []string
		insertionRules map[[2]string]string
		steps          int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"small test", args{[]string{"A", "B", "C"}, map[[2]string]string{{"A", "B"}: "H", {"B", "C"}: "N"}, 1}, []string{"A", "H", "B", "N", "C"}},
		{"small test", args{[]string{"A", "B", "C"}, map[[2]string]string{{"A", "B"}: "A", {"B", "C"}: "N"}, 1}, []string{"A", "A", "B", "N", "C"}},
		{"small test", args{[]string{"A", "B", "C"}, map[[2]string]string{{"A", "B"}: "A", {"B", "C"}: "N"}, 2}, []string{"A", "A", "A", "B", "N", "C"}},
		{"bigger test", args{p, i, 1}, []string{"N", "C", "N", "B", "C", "H", "B"}},
		{"bigger test", args{p, i, 2}, strings.Split("NBCCNBBBCBHCB", "")},
		{"bigger test", args{p, i, 4}, strings.Split("NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", "")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertSteps(tt.args.polymer, tt.args.insertionRules, tt.args.steps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_count(t *testing.T) {
	p, i := parseInput(useTestData())
	pol := insertSteps(p, i, 10)
	pR, iR := parseInput(useInputData())
	polR := insertSteps(pR, iR, 10)
	type args struct {
		polymer []string
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]int
		want1 int
		want2 int
		want3 int
	}{
		{"test small", args{[]string{"A", "H", "H", "B", "N", "C"}}, map[string]int{"A": 1, "B": 1, "C": 1, "H": 2, "N": 1}, 1, 2, 1},

		{"test 10", args{pol}, map[string]int{"B": 1749, "C": 298, "H": 161, "N": 865}, 161, 1749, 1588},
		{"part 1", args{polR}, map[string]int{"B": 1718, "C": 4015, "F": 2859, "H": 1293, "K": 1137, "N": 3527, "O": 1195, "P": 920, "S": 1837, "V": 956}, 920, 4015, 3095},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := count(tt.args.polymer)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
			assert.Equal(t, tt.want2, got2)
			assert.Equal(t, tt.want3, got2-got1)
		})
	}
}

var insRules = map[[2]string][2][2]string{{"A", "A"}: {{"A", "B"}, {"B", "A"}}, {"B", "C"}: {{"B", "N"}, {"N", "C"}}}
var insRulesT = map[[2]string][2][2]string{{"B", "B"}: {{"B", "N"}, {"N", "B"}}, {"B", "C"}: {{"B", "B"}, {"B", "C"}}, {"B", "H"}: {{"B", "H"}, {"H", "H"}}, {"B", "N"}: {{"B", "B"}, {"B", "N"}}, {"C", "B"}: {{"C", "H"}, {"H", "B"}}, {"C", "C"}: {{"C", "N"}, {"N", "C"}}, {"C", "H"}: {{"C", "B"}, {"B", "H"}}, {"C", "N"}: {{"C", "C"}, {"C", "N"}}, {"H", "B"}: {{"H", "C"}, {"C", "B"}}, {"H", "C"}: {{"H", "B"}, {"B", "C"}}, {"H", "H"}: {{"H", "N"}, {"N", "H"}}, {"H", "N"}: {{"H", "C"}, {"C", "N"}}, {"N", "B"}: {{"N", "B"}, {"B", "B"}}, {"N", "C"}: {{"N", "B"}, {"B", "C"}}, {"N", "H"}: {{"N", "C"}, {"C", "H"}}, {"N", "N"}: {{"N", "C"}, {"C", "N"}}}

func Test_parseInputPairwise(t *testing.T) {
	type args struct {
		parts []string
	}
	tests := []struct {
		name  string
		args  args
		want  map[[2]string]int
		want1 map[[2]string][2][2]string
	}{
		{"test small", args{[]string{"ABC", "AA -> B\nBC -> N"}}, map[[2]string]int{{"A", "B"}: 1, {"B", "C"}: 1}, insRules},
		{"test test", args{useTestData()}, map[[2]string]int{{"C", "B"}: 1, {"N", "C"}: 1, {"N", "N"}: 1}, insRulesT},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseInputPairwise(tt.args.parts)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_pairWiseInsertion(t *testing.T) {
	poly := map[[2]string]int{{"A", "B"}: 1, {"B", "C"}: 1}
	insRules := map[[2]string][2][2]string{{"A", "A"}: {{"A", "B"}, {"B", "A"}}, {"B", "C"}: {{"B", "N"}, {"N", "C"}}}
	insRules2 := map[[2]string][2][2]string{{"A", "B"}: {{"A", "B"}, {"B", "A"}}, {"B", "C"}: {{"B", "N"}, {"N", "C"}}}
	type args struct {
		polymer        map[[2]string]int
		insertionRules map[[2]string][2][2]string
	}
	tests := []struct {
		name string
		args args
		want map[[2]string]int
	}{
		{"small", args{poly, insRules}, map[[2]string]int{{"A", "B"}: 1, {"B", "C"}: 0, {"B", "N"}: 1, {"N", "C"}: 1}},
		{"small2", args{map[[2]string]int{{"A", "B"}: 1, {"B", "C"}: 1}, insRules2}, map[[2]string]int{{"A", "B"}: 1, {"B", "A"}: 1, {"B", "C"}: 0, {"B", "N"}: 1, {"N", "C"}: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pairWiseInsertion(tt.args.polymer, tt.args.insertionRules)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_countPair(t *testing.T) {
	p, i := parseInputPairwise(useTestData())
	pol := insertPairwiseSteps(p, i, 10)
	pR, iR := parseInputPairwise(useInputData())
	polR := insertPairwiseSteps(pR, iR, 10)
	type args struct {
		polymer map[[2]string]int
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]int
		want1 int
		want2 int
		want3 int
	}{
		{"test 10", args{pol}, map[string]int{"B": 1749, "C": 298, "H": 161, "N": 865}, 161, 1749, 1588},
		{"part 1", args{polR}, map[string]int{"B": 1718, "C": 4015, "F": 2859, "H": 1293, "K": 1137, "N": 3527, "O": 1195, "P": 920, "S": 1837, "V": 956}, 920, 4015, 3095},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := countPair(tt.args.polymer)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
			assert.Equal(t, tt.want2, got2)
			assert.Equal(t, tt.want3, got2-got1)
		})
	}
}

func Test_part2(t *testing.T) {
	p, i := parseInputPairwise(useTestData())
	pol := insertPairwiseSteps(p, i, 40)
	pR, iR := parseInputPairwise(useInputData())
	polR := insertPairwiseSteps(pR, iR, 40)
	type args struct {
		polymer map[[2]string]int
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]int
		want1 int
		want2 int
		want3 int
	}{
		{"test 40", args{pol}, map[string]int{"B": 1749, "C": 298, "H": 161, "N": 865}, 161, 1749, 2188189693529},
		{"part 2", args{polR}, map[string]int{"B": 1749, "C": 298, "H": 161, "N": 865}, 161, 1749, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := countPair(tt.args.polymer)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
			assert.Equal(t, tt.want2, got2)
			assert.Equal(t, tt.want3, got2-got1)
		})
	}
}
