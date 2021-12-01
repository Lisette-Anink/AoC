package day_one

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func TestMain(t *testing.T) {
	input := utils.ImportFileLines("test_input_1")
	lines := utils.SplitInLines(input)
	t.Logf("%v", lines)
	ints := utils.ConvertToInt(lines)
	{
		assert.Equal(t, 124, ints[0])
	}
	{
		new := CountIncreased(ints)
		assert.Equal(t, 7, new)
	}
}

func TestPart1(t *testing.T) {
	input := utils.ImportFromAoC("2021", "1")
	lines := utils.SplitInLines(input)
	ints := utils.ConvertToInt(lines)
	{
		new := CountIncreased(ints)
		assert.Equal(t, 1559, new)
	}
}

func TestCountIncreasedSlidingWindow(t *testing.T) {
	input := utils.ImportFileLines("test_input_1")
	lines := utils.SplitInLines(input)
	t.Logf("%v", lines)
	ints := utils.ConvertToInt(lines)
	{
		new := CountIncreasedSlidingWindow(ints)
		assert.Equal(t, 5, new)
	}
}

func TestPart2(t *testing.T) {
	input := utils.ImportFromAoC("2021", "1")

	lines := utils.SplitInLines(input)
	t.Logf("%v", lines)
	ints := utils.ConvertToInt(lines)
	{
		new := CountIncreasedSlidingWindow(ints)
		assert.Equal(t, 1600, new)
	}
}
