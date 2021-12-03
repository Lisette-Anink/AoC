package day_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func TestCalPosition(t *testing.T) {
	input := utils.ImportFileLines("test_input_2")
	lines := utils.SplitInLines(input)
	t.Logf("%v", lines)
	{
		assert.Equal(t, "forward 5", lines[0])
	}
	t.Run("test input", func(t *testing.T) {
		h, d := calPosition(lines)
		assert.Equal(t, 15, h)
		assert.Equal(t, 10, d)
	})
}
func TestPart1(t *testing.T) {
	input := utils.ImportFromAoC("2021", "2")
	lines := utils.SplitInLines(input)
	// t.Logf("%v", lines)
	{
		assert.Equal(t, "forward 8", lines[0])
	}
	t.Run("test input", func(t *testing.T) {
		h, d := calPosition(lines)
		assert.Equal(t, 2003, h)
		assert.Equal(t, 872, d)
		assert.Equal(t, 1746616, d*h)
	})
}
func TestCalPositionAim(t *testing.T) {
	input := utils.ImportFileLines("test_input_2")
	lines := utils.SplitInLines(input)

	t.Run("test input", func(t *testing.T) {
		h, d := calPositionAim(lines)
		assert.Equal(t, 15, h)
		assert.Equal(t, 60, d)
	})
}
func TestPart2(t *testing.T) {
	input := utils.ImportFromAoC("2021", "2")
	lines := utils.SplitInLines(input)
	t.Run("test input", func(t *testing.T) {
		h, d := calPositionAim(lines)
		assert.Equal(t, 2003, h)
		assert.Equal(t, 869681, d)
		assert.Equal(t, 1741971043, d*h)
	})
}
