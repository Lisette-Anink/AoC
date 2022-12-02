package day_24

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func useTestData() []string {
	input := utils.ImportFileLines("test_input_24")
	lines := utils.SplitInLines(input)
	return lines
}

func useInputData() []string {
	input := utils.ImportFromAoC("2021", "24")
	lines := utils.SplitInLines(input)
	return lines
}

func TestALU_inp(t *testing.T) {
	type fields struct {
		w int
		x int
		y int
		z int
	}
	type args struct {
		a       string
		inCount int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ALU
	}{
		{"test 1", fields{0, 0, 0, 0}, args{"w", 0}, &ALU{9, 0, 0, 0}},
		{"test x1", fields{0, 0, 0, 0}, args{"x", 1}, &ALU{0, 9, 0, 0}},
		{"test y2", fields{0, 0, 0, 0}, args{"y", 2}, &ALU{0, 0, 9, 0}},
		{"test z3", fields{0, 0, 0, 0}, args{"z", 3}, &ALU{0, 0, 0, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alu := &ALU{
				w: tt.fields.w,
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}
			alu.inp(tt.args.a, tt.args.inCount)
			assert.Equal(t, tt.want, alu)
		})
	}
}

func TestALU_add(t *testing.T) {
	type fields struct {
		w int
		x int
		y int
		z int
	}
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ALU
	}{
		{"test 1", fields{0, 0, 0, 0}, args{"w", "1"}, &ALU{1, 0, 0, 0}},
		{"test x1", fields{0, 0, 0, 0}, args{"x", "1"}, &ALU{0, 1, 0, 0}},
		{"test y2", fields{0, 0, 0, 0}, args{"y", "2"}, &ALU{0, 0, 2, 0}},
		{"test z3", fields{0, 0, 0, 0}, args{"z", "3"}, &ALU{0, 0, 0, 3}},
		{"test zx", fields{1, 10, 5, 1}, args{"z", "x"}, &ALU{1, 10, 5, 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alu := &ALU{
				w: tt.fields.w,
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}
			alu.add(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, alu)
		})
	}
}

func TestALU_mul(t *testing.T) {
	type fields struct {
		w int
		x int
		y int
		z int
	}
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ALU
	}{
		{"test 1", fields{1, 2, 3, 4}, args{"w", "5"}, &ALU{5, 2, 3, 4}},
		{"test y2", fields{1, 2, 3, 4}, args{"y", "2"}, &ALU{1, 2, 6, 4}},
		{"test z3", fields{1, 2, 3, 4}, args{"z", "3"}, &ALU{1, 2, 3, 12}},
		{"test zx", fields{1, 10, 5, 2}, args{"z", "x"}, &ALU{1, 10, 5, 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alu := &ALU{
				w: tt.fields.w,
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}
			alu.mul(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, alu)
		})
	}
}

func TestALU_div(t *testing.T) {
	type fields struct {
		w int
		x int
		y int
		z int
	}
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ALU
	}{
		{"test 1", fields{10, 2, 3, 4}, args{"w", "5"}, &ALU{2, 2, 3, 4}},
		{"test y2", fields{1, 2, 3, 4}, args{"y", "2"}, &ALU{1, 2, 1, 4}},
		{"test z3", fields{1, 2, 3, 11}, args{"z", "3"}, &ALU{1, 2, 3, 3}},
		{"test zx", fields{1, 10, 5, 2}, args{"z", "x"}, &ALU{1, 10, 5, 0}},
		{"test zx", fields{1, 2, 5, 10}, args{"z", "x"}, &ALU{1, 2, 5, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alu := &ALU{
				w: tt.fields.w,
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}
			alu.div(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, alu)
		})
	}
}

func TestALU_mod(t *testing.T) {
	type fields struct {
		w int
		x int
		y int
		z int
	}
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ALU
	}{
		{"test 1", fields{11, 2, 3, 4}, args{"w", "5"}, &ALU{1, 2, 3, 4}},
		{"test y2", fields{1, 2, 3, 4}, args{"y", "2"}, &ALU{1, 2, 1, 4}},
		{"test z3", fields{1, 2, 3, 11}, args{"z", "3"}, &ALU{1, 2, 3, 2}},
		{"test zx", fields{1, 10, 5, 2}, args{"z", "x"}, &ALU{1, 10, 5, 2}},
		{"test zx", fields{1, 2, 5, 10}, args{"z", "x"}, &ALU{1, 2, 5, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alu := &ALU{
				w: tt.fields.w,
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}
			alu.mod(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, alu)
		})
	}
}

func TestALU_eql(t *testing.T) {
	type fields struct {
		w int
		x int
		y int
		z int
	}
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ALU
	}{
		{"test 1", fields{5, 2, 3, 4}, args{"w", "5"}, &ALU{1, 2, 3, 4}},
		{"test y2", fields{1, 2, 3, 4}, args{"y", "2"}, &ALU{1, 2, 0, 4}},
		{"test zx", fields{1, 2, 5, 2}, args{"z", "x"}, &ALU{1, 2, 5, 1}},
		{"test zy", fields{1, 2, 5, 10}, args{"z", "y"}, &ALU{1, 2, 5, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alu := &ALU{
				w: tt.fields.w,
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}
			alu.eql(tt.args.a, tt.args.b)
		})
	}
}

func Test_run(t *testing.T) {
	type args struct {
		lines []string
		input []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test", args{useTestData(), []int{9, 9, 9, 9, 9}}, false},
		{"test", args{useTestData(), []int{8, 9, 9, 9, 9}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := run(tt.args.lines, tt.args.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMONADcalc(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Part 1", args{useInputData()}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()

			if got := MONADcalc(tt.args.lines); got != tt.want {
				t.Errorf("MONADcalc() = %v, want %v", got, tt.want)
			}
			time := time.Now()
			elapsed := time.Sub(start)
			t.Logf("this took %v", elapsed)
		})
	}
}
