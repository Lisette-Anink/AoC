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

// func TestCalPositionAim(t *testing.T) {
// 	input := utils.ImportFileLines("test_input_2")
// 	lines := utils.SplitInLines(input)

// 	t.Run("test input", func(t *testing.T) {
// 		h, d := calPositionAim(lines)
// 		assert.Equal(t, 15, h)
// 		assert.Equal(t, 60, d)
// 	})
// }
// func TestPart2(t *testing.T) {
// 	input := utils.ImportFromAoC("2021", "2")
// 	lines := utils.SplitInLines(input)
// 	t.Run("test input", func(t *testing.T) {
// 		h, d := calPositionAim(lines)
// 		assert.Equal(t, 2003, h)
// 		assert.Equal(t, 869681, d)
// 		assert.Equal(t, 1741971043, d*h)
// 	})
// }
