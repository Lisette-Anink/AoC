package day_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func TestMain(t *testing.T) {
	input := utils.ImportFileLines("test_input_2")
	lines := utils.SplitInLines(input)
	t.Logf("%v", lines)
	{
		assert.Equal(t, "abc", lines[0])
	}
	t.Run("add first", func(t *testing.T) {
		new := AddW(lines[1])
		assert.Equal(t, "defw", new)
	})
}
