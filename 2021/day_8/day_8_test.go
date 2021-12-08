package day_8

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func Test_parseInput(t *testing.T) {
	data := useTestData()

	type args struct {
		lines []string
	}
	tests := []struct {
		name             string
		args             args
		wantInputDigits  []string
		wantOutputDigits []string
	}{
		{"small", args{[]string{"acd def | ket sle", "hsg ksh | kse lsos"}}, []string{"acd", "def", "hsg", "ksh"}, []string{"ket", "sle", "kse", "lsos"}},
		{"testdata", args{data}, []string{"acd", "def", "hsg", "ksh"}, []string{"ket", "sle", "kse", "lsos"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInputDigits, gotOutputDigits := parseInput(tt.args.lines)
			assert.ElementsMatch(t, tt.wantInputDigits, gotInputDigits)
			assert.ElementsMatch(t, tt.wantOutputDigits, gotOutputDigits)

		})
	}
}

func Test_countNumbers(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{"small", args{[]string{"acd", "sedefas", "hs", "kshd"}}, 3},
		{"small 2", args{[]string{"aacd", "sedefas", "rghsf", "ksh"}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := countNumbers(tt.args.in); gotCount != tt.wantCount {
				t.Errorf("countNumbers() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func useTestData() []string {
	input := utils.ImportFileLines("test_input_8")
	lines := utils.SplitInLines(input)
	return lines
}

func useInputData() []string {
	input := utils.ImportFromAoC("2021", "8")
	lines := utils.SplitInLines(input)
	return lines
}
func Test_Part1(t *testing.T) {
	{
		_, out := parseInput(useTestData())
		count := countNumbers(out)
		assert.Equal(t, 26, count)
	}
	{
		_, out := parseInput(useInputData())
		count := countNumbers(out)
		assert.Equal(t, 26, count)
	}
}

var example_in = []string{"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"}

// acedgfb: 8
// dab: 7
// eafb: 4
// ab: 1
// cdfbe: 5
// gcdfa: 2
// fbcad: 3
// cefabd: 9
// cdfgeb: 6
// cagedb: 0

var example = map[int][]string{0: {"c", "a", "g", "e", "d", "b"}, 1: {"a", "b"}, 2: {"g", "c", "d", "f", "a"}, 3: {"c", "d", "b", "a", "f"}, 4: {"e", "a", "f", "b"}, 5: {"c", "d", "f", "e", "b"}, 7: {"d", "a", "b"}, 8: {"a", "c", "e", "d", "g", "f", "b"}, 6: {"c", "d", "f", "g", "e", "b"}, 9: {"c", "e", "f", "a", "b", "d"}}

// - top = 7-1 "d"
// - 4- 1 = mid & topleft "e""f"
// - 9 is len6 includes 4
// - 0 is other len6 contains 7
// - other len6 is 6
// - 3 is len5 includes 7
// - 5 is len5 includes mid&topleft
// - other len5 is 2

// 0:      1:      2:      3:      4:
//  aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
//  ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
//  gggg    ....    gggg    gggg    ....

//  5:      6:      7:      8:      9:
//  aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
//  dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
//  gggg    gggg    ....    gggg    gggg

func Test_identifyNumbers(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name     string
		args     args
		wantIden map[int][]string
	}{
		{"cse 1", args{[]string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab", "cdfeb", "fcadb", "cdfeb", "cdbaf"}}, example},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIden := identifyNumbers(tt.args.in)
			assert.Equal(t, tt.wantIden, gotIden)
		})
	}
}

func Test_Part2(t *testing.T) {
	{
		// data := useTestData()

	}
	{
		data := example_in
		in, out := parseInput(data)
		all := in
		all = append(all, out...)
		// identifyNumbers(all)
		iden := identifyNumbers(all)
		val := valDigits(iden, out)
		assert.Equal(t, 5353, val)
	}
}
