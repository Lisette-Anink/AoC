package day_6

import (
	"regexp"

	"lisette.anink/aoc/utils"
)

type vector struct {
	x1, y1, x2, y2 int
}

func parseInput(lines []string) []vector {
	vectors := []vector{}
	for _, line := range lines {
		if len(line) > 0 {
			r, _ := regexp.Compile(`(\d*),(\d*) -> (\d*),(\d*)`)
			matched := r.FindStringSubmatch(line)
			matchedI := utils.ConvertToInt(matched[1:])
			// fmt.Println(matched, matchedI)
			v := vector{
				x1: matchedI[0],
				y1: matchedI[1],
				x2: matchedI[2],
				y2: matchedI[3],
			}
			vectors = append(vectors, v)
		}
	}
	return vectors
}

func calcDiagram(vectors []vector) map[[2]int]int {
	diagram := map[[2]int]int{}
	for _, vec := range vectors {
		if vec.isHor() {
			min, max := minmax(vec.y1, vec.y2)
			for y := min; y <= max; y++ {
				diagram[[2]int{vec.x1, y}]++
			}
		}
		if vec.isVert() {
			min, max := minmax(vec.x1, vec.x2)
			for x := min; x <= max; x++ {
				diagram[[2]int{x, vec.y1}]++
			}
		}
	}
	return diagram
}
func calcDiagramDia(vectors []vector) map[[2]int]int {
	diagram := map[[2]int]int{}
	for _, vec := range vectors {
		if vec.isHor() {
			min, max := minmax(vec.y1, vec.y2)
			for y := min; y <= max; y++ {
				diagram[[2]int{vec.x1, y}]++
			}
		} else if vec.isVert() {
			min, max := minmax(vec.x1, vec.x2)
			for x := min; x <= max; x++ {
				diagram[[2]int{x, vec.y1}]++
			}
		} else {
			switch {
			case vec.y1 < vec.y2 && vec.x1 < vec.x2:
				// both small -> big
				minX, _ := minmax(vec.x1, vec.x2)
				minY, maxY := minmax(vec.y1, vec.y2)
				x := minX
				for y := minY; y <= maxY; y++ {
					diagram[[2]int{x, y}]++
					x++
				}
			case vec.y1 > vec.y2 && vec.x1 < vec.x2:
				// only x small -> big
				x := vec.x1
				for y := vec.y1; y >= vec.y2; y-- {
					diagram[[2]int{x, y}]++
					x++
				}
			case vec.y1 < vec.y2 && vec.x1 > vec.x2:
				minY, maxY := minmax(vec.y1, vec.y2)
				x := vec.x1
				for y := minY; y <= maxY; y++ {
					diagram[[2]int{x, y}]++
					x--
				}
			case vec.y1 > vec.y2 && vec.x1 > vec.x2:
				x := vec.x1
				for y := vec.y1; y >= vec.y2; y-- {
					diagram[[2]int{x, y}]++
					x--
				}
			}
		}
	}
	return diagram
}

func (v vector) isHor() bool {
	return v.x1 == v.x2
}
func (v vector) isVert() bool {
	return v.y1 == v.y2
}

func countHighVent(diagram map[[2]int]int) int {
	total := 0
	for _, val := range diagram {
		if val > 1 {
			total++
		}
	}
	return total
}

func minmax(f1, f2 int) (int, int) {
	if f1 < f2 {
		return f1, f2
	}
	return f2, f1
}
