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

func (p *paper) fold(inst instruction) {
	dots := p.Dots
	if inst.direction == "x" {
		for x := inst.coordinate + 1; x <= p.MaxX; x++ {
			// y =8  distance: 1 folded: 6
			distanceX := x - inst.coordinate
			foldedX := inst.coordinate - distanceX
			for y := 0; y <= p.MaxX; y++ {
				toFold := dots[[2]int{x, y}]
				if toFold {
					dots[[2]int{foldedX, y}] = toFold
				}
				delete(dots, [2]int{x, y})
			}
		}
		p.MaxX = inst.coordinate - 1
	} else {
		for y := inst.coordinate + 1; y <= p.MaxY; y++ {
			// y =8  distance: 1 folded: 6
			distanceY := y - inst.coordinate
			foldedY := inst.coordinate - distanceY
			for x := 0; x <= p.MaxX; x++ {
				toFold := dots[[2]int{x, y}]
				// fmt.Println(x, y, toFold)
				if toFold {
					dots[[2]int{x, foldedY}] = toFold
				}
				delete(dots, [2]int{x, y})
			}
		}
		p.MaxY = inst.coordinate - 1
		//y
	}
	p.Dots = dots
}

func (p *paper) print() {
	for y := 0; y <= p.MaxX; y++ {
		for x := 0; x <= p.MaxX; x++ {
			pos := [2]int{x, y}
			if p.Dots[pos] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
			if pos[0] == p.MaxX {
				fmt.Println()
			}
		}
	}
}
