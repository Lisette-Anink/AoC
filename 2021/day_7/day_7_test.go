package day_7

import (
	"reflect"
	"testing"

	"lisette.anink/aoc/utils"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case 1", args{"1,2,3"}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
func useTestData() string {
	input := utils.ImportFileLines("test_input_7")
	lines := utils.SplitInLines(input)
	return lines[0]
}

func useInputData() string {
	input := utils.ImportFromAoC("2021", "7")
	lines := utils.SplitInLines(input)
	return lines[0]
}

func Test_calFuel(t *testing.T) {
	test := parseInput(useTestData())
	type args struct {
		crabs []int
		pos   int
	}
	tests := []struct {
		name      string
		args      args
		wantTotal float64
	}{
		{"case 1", args{[]int{1, 2, 3}, 3}, 3},
		{"input", args{test, 3}, 39},
		{"input", args{test, 10}, 71},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := calFuel(tt.args.crabs, tt.args.pos); gotTotal != tt.wantTotal {
				t.Errorf("calFuel() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func Test_max(t *testing.T) {
	test := parseInput(useTestData())
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{[]int{1, 2, 3}}, 3},
		{"case 2", args{[]int{3, 2, 1}}, 3},
		{"case 200", args{[]int{3, 200, 1}}, 200},
		{"case test", args{test}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxInt(tt.args.a); got != tt.want {
				t.Errorf("maxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calMinFuel(t *testing.T) {
	test := parseInput(useTestData())
	part1 := parseInput(useInputData())

	type args struct {
		crabs []int
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 int
	}{
		{"case test", args{test}, 37, 2},
		{"case part 1", args{part1}, 341534, 363},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calMinFuel(tt.args.crabs)
			if got != tt.want {
				t.Errorf("calMinFuel() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calMinFuel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_minFloat(t *testing.T) {
	type args struct {
		fuel map[int]float64
	}
	tests := []struct {
		name  string
		args  args
		wantM float64
		wantP int
	}{
		{"case test", args{map[int]float64{1: 2, 3: 4, 5: 1}}, 1, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM, gotP := minFloat(tt.args.fuel)
			if gotM != tt.wantM {
				t.Errorf("minFloat() gotM = %v, want %v", gotM, tt.wantM)
			}
			if gotP != tt.wantP {
				t.Errorf("minFloat() gotP = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func TestBruteForceFuel(t *testing.T) {
	test := parseInput(useTestData())
	type args struct {
		crabs []int
		pos   int
	}
	tests := []struct {
		name      string
		args      args
		wantTotal float64
	}{
		{"case 1", args{[]int{1, 2, 3}, 3}, 4},
		{"input", args{test, 2}, 206},
		{"input", args{test, 5}, 168},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := calBruteForceFuel(tt.args.crabs, tt.args.pos); gotTotal != tt.wantTotal {
				t.Errorf("calBruteForceFuel() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func Test_calMinFuelPart2(t *testing.T) {
	test := parseInput(useTestData())
	part2 := parseInput(useInputData())
	type args struct {
		crabs []int
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 int
	}{
		{"case 1", args{[]int{1, 2, 3}}, 2, 2},
		{"test", args{test}, 168, 5},
		{"input", args{part2}, 93397632, 484},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calMinFuelPart2(tt.args.crabs)
			if got != tt.want {
				t.Errorf("calMinFuelPart2() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calMinFuelPart2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_calGaussFuel(t *testing.T) {
	test := parseInput(useTestData())
	type args struct {
		crabs []int
		pos   int
	}
	tests := []struct {
		name      string
		args      args
		wantTotal float64
	}{
		{"case 1", args{[]int{1, 2, 3}, 3}, 4},
		{"input", args{test, 2}, 206},
		{"input", args{test, 5}, 168},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := calGaussFuel(tt.args.crabs, tt.args.pos); gotTotal != tt.wantTotal {
				t.Errorf("calGaussFuel() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func Test_calcGauss(t *testing.T) {
	type args struct {
		c float64
		p float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"case 1", args{16, 5}, 66},
		{"case 2", args{1, 5}, 10},
		{"case 2", args{2, 5}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcGauss(tt.args.c, tt.args.p); got != tt.want {
				t.Errorf("calcGauss() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmark_calcGauss(in float64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		calcGauss(2, in)
	}
}

func Benchmark_calcGauss100(b *testing.B)   { benchmark_calcGauss(100, b) }
func Benchmark_calcGauss1000(b *testing.B)  { benchmark_calcGauss(1000, b) }
func Benchmark_calcGauss10000(b *testing.B) { benchmark_calcGauss(10000, b) }

func benchmark_calBruteForce(in float64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		calBruteForce(2, in)
	}
}

func Benchmark_calBruteForce100(b *testing.B)   { benchmark_calBruteForce(100, b) }
func Benchmark_calBruteForce1000(b *testing.B)  { benchmark_calBruteForce(1000, b) }
func Benchmark_calBruteForce10000(b *testing.B) { benchmark_calBruteForce(10000, b) }
