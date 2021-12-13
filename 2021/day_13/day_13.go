package day_11

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"lisette.anink/aoc/utils"
)

type instruction struct {
	direction  string
	coordinate int
}

type paper struct {
	Dots map[[2]int]bool
	MaxX int
	MaxY int
}

func parseInput(lines []string) (paper, []instruction) {
	dots := strings.Split(lines[0], "\n")
	newPaper := paper{}
	dotMap := map[[2]int]bool{}
	for _, line := range dots {
		parts := utils.ConvertToInt(strings.Split(line, ","))
		if newPaper.MaxX < parts[0] {
			newPaper.MaxX = parts[0]
		}
		if newPaper.MaxY < parts[1] {
			newPaper.MaxY = parts[1]
		}
		dotMap[[2]int{parts[0], parts[1]}] = true
	}
	newPaper.Dots = dotMap
	instructions := []instruction{}
	folds := strings.Split(lines[1], "\n")
	for _, fold := range folds {
		r := regexp.MustCompile(`fold along (x|y)\=(\d+)`)
		f := r.FindStringSubmatch(fold)
		if len(f) > 1 {
			fmt.Println(f)
			c, _ := strconv.Atoi(f[2])
			instructions = append(instructions, instruction{f[1], c})
		}
	}
	return newPaper, instructions
}

func foldPaper(p paper, inst instruction) {
	dots := p.Dots
	if inst.direction == "x" {

	} else {
		for i := inst.coordinate + 1; i <= p.MaxY; i++ {

		}
		//y
	}
}
