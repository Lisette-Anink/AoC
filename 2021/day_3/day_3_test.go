package day_3

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func TestSplitInChars(t *testing.T) {
	input := utils.ImportFileLines("test_input_3")
	lines := utils.SplitInLines(input)
	t.Logf("%v", lines)
	{
		assert.Equal(t, "00100", lines[0])
	}
	t.Run("test input", func(t *testing.T) {
		out := splitInChars(lines)
		assert.Equal(t, []string{"1", "1", "1", "1", "0"}, out[1])
	})
}
func TestTranspose(t *testing.T) {
	in := [][]string{{"0", "0"}, {"1", "1"}}
	want := map[int][]string{0: {"0", "1"}, 1: {"0", "1"}}
	t.Run("transpose", func(t *testing.T) {
		got := transpose(in)
		assert.Equal(t, want, got)
	})
}
func TestLeastMost(t *testing.T) {
	t.Run("leastmost1", func(t *testing.T) {
		in := []string{"0", "0", "1"}
		wantLeast, wantMost := "1", "0"
		l, m := leastMost(in)
		assert.Equal(t, wantLeast, l)
		assert.Equal(t, wantMost, m)
	})
	t.Run("leastmost2", func(t *testing.T) {
		in := []string{"0", "0", "1", "1", "1", "1"}
		wantLeast, wantMost := "0", "1"
		l, m := leastMost(in)
		assert.Equal(t, wantLeast, l)
		assert.Equal(t, wantMost, m)
	})
	t.Run("equal", func(t *testing.T) {
		in := []string{"0", "0", "0", "1", "1", "1"}
		wantLeast, wantMost := "0", "1"
		l, m := leastMost(in)
		assert.Equal(t, wantLeast, l)
		assert.Equal(t, wantMost, m)
	})
}
func TestFindGammaEps(t *testing.T) {
	input := utils.ImportFileLines("test_input_3")
	lines := utils.SplitInLines(input)
	// t.Logf("%v", lines)
	{
		assert.Equal(t, "00100", lines[0])
	}
	t.Run("test input", func(t *testing.T) {
		g, e := findGammaEps(lines)
		assert.Equal(t, int64(22), g)
		assert.Equal(t, int64(9), e)
	})
}

func TestPart1(t *testing.T) {
	input := utils.ImportFromAoC("2021", "3")
	lines := utils.SplitInLines(input)
	{
		assert.Equal(t, "110001010110", lines[0])
	}
	t.Run("real input", func(t *testing.T) {
		g, e := findGammaEps(lines)
		assert.Equal(t, int64(3516), g)
		assert.Equal(t, int64(579), e)

		assert.Equal(t, int64(2035764), g*e)
	})
}

// oxygen generator rating, determine the most common
//CO2 scrubber rating, determine the least common

func TestFindAll(t *testing.T) {
	in := []string{"001", "101", "000"}
	want0 := []string{"001", "000"}
	want1 := []string{"101"}
	want2 := []string{"001", "101"}

	t.Run("findall", func(t *testing.T) {
		got0 := findAll(in, rune('0'), 0)
		got1 := findAll(in, rune('1'), 0)
		got2 := findAll(in, rune('1'), 2)
		assert.Equal(t, want0, got0)
		assert.Equal(t, want1, got1)
		assert.Equal(t, want2, got2)
	})
}

func TestFindOxCoRating(t *testing.T) {
	input := utils.ImportFileLines("test_input_3")
	lines := utils.SplitInLines(input)

	t.Run("test input", func(t *testing.T) {
		o, c := findOxCoRating(lines)
		assert.Equal(t, int64(23), o)
		assert.Equal(t, int64(10), c)

		assert.Equal(t, int64(230), c*o)
	})
}

func TestPart2(t *testing.T) {
	input := utils.ImportFromAoC("2021", "3")
	lines := utils.SplitInLines(input)

	t.Run("test input", func(t *testing.T) {
		o, c := findOxCoRating(lines)
		assert.Equal(t, int64(3311), o)
		assert.Equal(t, int64(851), c)

		assert.Equal(t, int64(2817661), c*o)
	})
}

func Test_convertBinaryToInt(t *testing.T) {
	type args struct {
		binaryString string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"1", args{"01010"}, 10},
		{"1", args{"10111"}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBinaryToInt(tt.args.binaryString); got != tt.want {
				t.Errorf("convertBinaryToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
