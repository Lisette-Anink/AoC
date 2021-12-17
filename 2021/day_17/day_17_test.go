package day_17

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func useTestData() string {
	input := utils.ImportFileLines("test_input_17")
	return string(input)
}

func useInputData() string {
	input := utils.ImportFromAoC("2021", "17")
	return string(input)
}

func Test_parseInput(t *testing.T) {
	type args struct {
		lines string
	}
	tests := []struct {
		name string
		args args
		want targetArea
	}{
		{"test", args{useTestData()}, targetArea{20, -10, 30, -5}},
		{"input", args{useInputData()}, targetArea{xMin: 150, yMin: -129, xMax: 171, yMax: -70}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseInput(tt.args.lines)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_probe_trajectoryStep(t *testing.T) {
	type fields struct {
		xPos int
		yPos int
		xVel int
		yVel int
	}
	tests := []struct {
		name   string
		fields fields
		probe  *probe
	}{
		{"step1", fields{0, 0, 7, 2}, &probe{xPos: 7, yPos: 2, xVel: 6, yVel: 1}},
		{"step2", fields{xPos: 7, yPos: 2, xVel: 6, yVel: 1}, &probe{xPos: 13, yPos: 3, xVel: 5, yVel: 0}},
		{"step3", fields{xPos: 13, yPos: 3, xVel: 5, yVel: 0}, &probe{xPos: 18, yPos: 3, xVel: 4, yVel: -1}},
		{"step4", fields{xPos: 18, yPos: 3, xVel: 4, yVel: -1}, &probe{xPos: 22, yPos: 2, xVel: 3, yVel: -2}},
		{"step5", fields{xPos: 22, yPos: 2, xVel: 3, yVel: -2}, &probe{xPos: 25, yPos: 0, xVel: 2, yVel: -3}},
		{"step6", fields{xPos: 25, yPos: 0, xVel: 2, yVel: -3}, &probe{xPos: 27, yPos: -3, xVel: 1, yVel: -4}},
		{"step7", fields{xPos: 27, yPos: -3, xVel: 1, yVel: -4}, &probe{xPos: 28, yPos: -7, xVel: 0, yVel: -5}},
		{"step8", fields{xPos: 28, yPos: -7, xVel: 0, yVel: -5}, &probe{xPos: 28, yPos: -12, xVel: 0, yVel: -6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &probe{
				xPos: tt.fields.xPos,
				yPos: tt.fields.yPos,
				xVel: tt.fields.xVel,
				yVel: tt.fields.yVel,
			}
			p.trajectoryStep()
			assert.Equal(t, tt.probe, p)
		})
	}
}

var target = targetArea{20, -10, 30, -5}

func Test_probe_pastTarget(t *testing.T) {
	type fields struct {
		xPos int
		yPos int
		xVel int
		yVel int
	}
	type args struct {
		t targetArea
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"before", fields{xPos: 22, yPos: 2, xVel: 3, yVel: -2}, args{target}, false},
		{"in target", fields{xPos: 28, yPos: -7, xVel: 0, yVel: -5}, args{target}, false},
		{"past true", fields{xPos: 28, yPos: -12, xVel: 0, yVel: -6}, args{target}, true},
		{"past true", fields{xPos: 35, yPos: -7, xVel: 0, yVel: -6}, args{target}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &probe{
				xPos: tt.fields.xPos,
				yPos: tt.fields.yPos,
				xVel: tt.fields.xVel,
				yVel: tt.fields.yVel,
			}
			got := p.pastTarget(tt.args.t)
			assert.Equal(t, tt.want, got)
		})
	}
}
func Test_probe_inTarget(t *testing.T) {
	type fields struct {
		xPos int
		yPos int
		xVel int
		yVel int
	}
	type args struct {
		t targetArea
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"before", fields{xPos: 22, yPos: 2, xVel: 3, yVel: -2}, args{target}, false},
		{"in target", fields{xPos: 28, yPos: -7, xVel: 0, yVel: -5}, args{target}, true},
		{"past true", fields{xPos: 28, yPos: -12, xVel: 0, yVel: -6}, args{target}, false},
		{"past true", fields{xPos: 35, yPos: -7, xVel: 0, yVel: -6}, args{target}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &probe{
				xPos: tt.fields.xPos,
				yPos: tt.fields.yPos,
				xVel: tt.fields.xVel,
				yVel: tt.fields.yVel,
			}
			got := p.inTarget(tt.args.t)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_shootProbe(t *testing.T) {
	target := parseInput(useTestData())

	type args struct {
		input targetArea
		x     int
		y     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 7,2", args{target, 7, 2}, 3},
		{"test 6,3", args{target, 6, 3}, 6},
		{"test 9,0", args{target, 9, 0}, 0},
		{"test 17,-4", args{target, 17, -4}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shootProbe(tt.args.input, tt.args.x, tt.args.y)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findHighest(t *testing.T) {
	type args struct {
		input string
		ymax  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{useTestData(), 10}, 45},
		{"solve Part 1/10", args{useInputData(), 10}, 8256},
		{"solve Part 1/100", args{useInputData(), 100}, 8256},
		{"solve Part 1/1000", args{useInputData(), 1000}, 8256},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findHighest(tt.args.input, tt.args.ymax); got != tt.want {
				t.Errorf("findHighest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countHits(t *testing.T) {
	type args struct {
		input string
		ymax  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{useTestData(), 10}, 112},
		{"solve Part 1/1000", args{useInputData(), 1000}, 2326},
		// {"solve Part 1/10000", args{useInputData(), 10000}, 2326},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countHits(tt.args.input, tt.args.ymax); got != tt.want {
				t.Errorf("countHits() = %v, want %v", got, tt.want)
			}
		})
	}
}
