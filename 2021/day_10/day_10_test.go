package day_10

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func useTestData() []string {
	input := utils.ImportFileLines("test_input_10")
	lines := utils.SplitInLines(input)
	return lines
}

func useInputData() []string {
	input := utils.ImportFromAoC("2021", "10")
	lines := utils.SplitInLines(input)
	return lines
}

// func Test_parseInput(t *testing.T) {
// 	type args struct {
// 		lines []string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want map[[2]int]int
// 	}{
// 		{"smal", args{[]string{"123", "456"}}, map[[2]int]int{{0, 0}: 1, {0, 1}: 2, {0, 2}: 3, {1, 0}: 4, {1, 1}: 5, {1, 2}: 6}},
// 		{"test", args{useTestData()}, map[[2]int]int{}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := parseInput(tt.args.lines)
// 			assert.Equal(t, tt.want, got)
// 		})
// 	}
// }

// func Test_parseLine(t *testing.T) {
// 	// lines := useTestData()
// 	type args struct {
// 		line string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		{"line 1", args{"((}))"}},
// 		{"line 1", args{"(({})[{}])"}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			parseLine(tt.args.line)
// 			assert.Equal(t, 3, 4)
// 		})
// 	}
// }

func Test_isOpen(t *testing.T) {
	// lines := useTestData()
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"line (", args{"("}, true},
		{"line {", args{"{"}, true},
		{"line }", args{"}"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isOpen(tt.args.line)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_isClose(t *testing.T) {
	type args struct {
		open  string
		close string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"line (", args{"(", ")"}, true},
		{"line {", args{"{", "}"}, true},
		{"line {", args{"<", ">"}, true},
		{"line {", args{"[", "]"}, true},
		{"line }", args{"[", "}"}, false},
		{"line }", args{"[", ")"}, false},
		{"line }", args{"[", ">"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isClose(tt.args.open, tt.args.close); got != tt.want {
				t.Errorf("isClose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCorrect(t *testing.T) {

	type args struct {
		elements map[int]string
		current  int
		close    int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		{"small test", args{map[int]string{0: "(", 1: "{", 2: "}", 3: ")"}, 0, 1}, true, "{"},
		{"small test fail", args{map[int]string{0: "(", 1: "{", 2: "}", 3: "[", 4: "}"}, 0, 1}, false, "]"},
		{"bigger test true", args{map[int]string{0: "[", 1: "(", 2: "(", 3: ")", 4: "[", 5: "<", 6: ">", 7: "]", 8: ")", 9: "]", 10: "(", 11: "{", 12: "[", 13: "<", 14: "{", 15: "<", 16: "<", 17: "[", 18: "]", 19: ">", 20: ">", 21: "("}, 0, 1}, true, "("},
		{"bigger test fail", args{map[int]string{0: "{", 1: "(", 2: "[", 3: "(", 4: "<", 5: "{", 6: "}", 7: "[", 8: "<", 9: ">", 10: "[", 11: "]", 12: "}", 13: ">", 14: "{", 15: "[", 16: "]", 17: "{", 18: "[", 19: "(", 20: "<", 21: "(", 22: ")", 23: ">"}, 0, 1}, false, "}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findCorrect(tt.args.elements, tt.args.current, tt.args.close)
			if got != tt.want {
				t.Errorf("findCorrect() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findCorrect() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_dumbWay(t *testing.T) {
	test := useTestData()
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test", args{test}, []string{")"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dumbWay(tt.args.lines)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_mapFromString(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want map[int]string
	}{
		{"test 1", args{"[(()[<>])]({[<{<<[]>>("}, map[int]string{0: "[", 1: "(", 2: "(", 3: ")", 4: "[", 5: "<", 6: ">", 7: "]", 8: ")", 9: "]", 10: "(", 11: "{", 12: "[", 13: "<", 14: "{", 15: "<", 16: "<", 17: "[", 18: "]", 19: ">", 20: ">", 21: "("}},
		{"test 2", args{"{([(<{}[<>[]}>{[]{[(<()>"}, map[int]string{0: "{", 1: "(", 2: "[", 3: "(", 4: "<", 5: "{", 6: "}", 7: "[", 8: "<", 9: ">", 10: "[", 11: "]", 12: "}", 13: ">", 14: "{", 15: "[", 16: "]", 17: "{", 18: "[", 19: "(", 20: "<", 21: "(", 22: ")", 23: ">"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapFromString(tt.args.in)
			assert.Equal(t, tt.want, got)
		})
	}
}
