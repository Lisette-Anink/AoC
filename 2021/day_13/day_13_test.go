package day_11

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func useTestData() []string {
	input := utils.ImportFileLines("test_input_13")
	parts := utils.SplitInMultiLines(input)
	return parts
}

func useInputData() []string {
	input := utils.ImportFromAoC("2021", "13")
	lines := utils.SplitInMultiLines(input)
	return lines
}

var testDots = paper{Dots: map[[2]int]bool{{0, 3}: true, {0, 13}: true, {0, 14}: true, {1, 10}: true, {2, 14}: true, {3, 0}: true, {3, 4}: true, {4, 1}: true, {4, 11}: true, {6, 0}: true, {6, 10}: true, {6, 12}: true, {8, 4}: true, {8, 10}: true, {9, 0}: true, {9, 10}: true, {10, 4}: true, {10, 12}: true}, MaxX: 10, MaxY: 14}

func Test_parseInput(t *testing.T) {
	testInput := useTestData()
	type args struct {
		lines []string
	}
	tests := []struct {
		name  string
		args  args
		want  paper
		want1 []instruction
	}{
		{"case 1", args{[]string{"6,10\n0,14\n9,10", "fold along y=713"}}, paper{Dots: map[[2]int]bool{{0, 14}: true, {6, 10}: true, {9, 10}: true}, MaxX: 9, MaxY: 14}, []instruction{{direction: "y", coordinate: 713}}},
		{"case testIn", args{testInput}, testDots, []instruction{{direction: "y", coordinate: 7}, {direction: "x", coordinate: 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseInput(tt.args.lines)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

var foldedPaper = paper{Dots: map[[2]int]bool{}, MaxX: 10, MaxY: 6}

func Test_paper_fold(t *testing.T) {
	type fields struct {
		Dots map[[2]int]bool
		MaxX int
		MaxY int
	}
	type args struct {
		inst instruction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   paper
	}{
		{"test 1", fields{testDots.Dots, testDots.MaxX, testDots.MaxY}, args{instruction{direction: "y", coordinate: 7}}, foldedPaper},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := paper{
				Dots: tt.fields.Dots,
				MaxX: tt.fields.MaxX,
				MaxY: tt.fields.MaxY,
			}
			p.fold(tt.args.inst)
			p.print()
			assert.Equal(t, tt.want, p)
			assert.Equal(t, 17, len(p.Dots))
		})
	}
}

func Test_Part1(t *testing.T) {
	data := useInputData()
	pap, instr := parseInput(data)
	pap.fold(instr[0])
	assert.Equal(t, 684, len(pap.Dots))
}

func Test_Part2(t *testing.T) {
	{
		data := useTestData()
		pap, instr := parseInput(data)
		for _, i := range instr {
			pap.fold(i)
		}
		pap.print()
		assert.Equal(t, 1, len(pap.Dots))
	}
	{
		data := useInputData()
		pap, instr := parseInput(data)
		for _, i := range instr {
			pap.fold(i)
		}
		pap.print()
		assert.Equal(t, 1, len(pap.Dots))
	}
}
