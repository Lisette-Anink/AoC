package day_9

import (
	"strconv"
	"strings"
)

func parseInput(lines []string) map[[2]int]int {
	var floormap = map[[2]int]int{}
	for x, line := range lines {
		if len(line) > 0 {
			parts := strings.Split(line, "")
			for y, p := range parts {
				i, _ := strconv.Atoi(p)
				floormap[[2]int{x, y}] = i
			}
		}
	}
	return floormap
}

func findLowPoints(floormap map[[2]int]int) int {
	totalRisk := 0
	for pos, val := range floormap {
		lowest := true
		// pos[0] //x
		// pos[1] //y
		if neighb, ok := floormap[[2]int{pos[0] - 1, pos[1]}]; ok {
			if neighb <= val {
				lowest = false
			}
		}
		if neighb, ok := floormap[[2]int{pos[0] + 1, pos[1]}]; ok {
			if neighb <= val {
				lowest = false
			}
		}
		if neighb, ok := floormap[[2]int{pos[0], pos[1] - 1}]; ok {
			if neighb <= val {
				lowest = false
			}
		}
		if neighb, ok := floormap[[2]int{pos[0], pos[1] + 1}]; ok {
			if neighb <= val {
				lowest = false
			}
		}
		if lowest {
			// fmt.Println(pos, val)
			totalRisk += val + 1
		}
	}
	return totalRisk
}

func findBasin(floormap map[[2]int]int) {
	// for pos range floormap
	// check if 9 skip
	// check in basin
	// no? put in basin /new
	// check neighbours recursive
	// del from floormap
}
