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
	for pos, val := range floormap {
		checkBasin(floormap, pos)
	}
	// check if 9 skip
	// check in basin
	// no? put in basin /new
	// check neighbours recursive
	// del from floormap
}

func checkBasinNeighbours(floormap map[[2]int]int, pos [2]int, basinNr int) {
	neighLeft := [2]int{pos[0] - 1, pos[1]}
	if val, ok := floormap[neighLeft]; ok {
		if val != 9 {
			addToBasin(neighLeft, val, basinNr)
			checkBasin(floormap, neighLeft)
		}
	}
	neighRight := [2]int{pos[0] + 1, pos[1]}
	if val, ok := floormap[neighRight]; ok {
		if val != 9 {
			addToBasin(neighRight, val, basinNr)
			checkBasin(floormap, neighRight)
		}
	}
	neighTop := [2]int{pos[0], pos[1] - 1}
	if val, ok := floormap[neighTop]; ok {
		if val != 9 {
			addToBasin(neighTop, val, basinNr)
			checkBasin(floormap, neighTop)
		}
	}
	neighBottom := [2]int{pos[0], pos[1] + 1}
	if val, ok := floormap[neighBottom]; ok {
		if val != 9 {
			addToBasin(neighBottom, val, basinNr)
			checkBasin(floormap, neighBottom)
		}
	}
}

func checkBasin(floormap map[[2]int]int, pos [2]int) {
	if floormap[pos] != 9 {
		nr := findBasinNr(pos)
		if nr == 0 {
			addToBasin(pos, floormap[pos], nr)
		} else {
			checkBasinNeighbours(floormap, pos, nr)
		}
	}
	delete(floormap, pos)
}

func addToBasin(pos [2]int, val, nr int) {
	if nr != 0 {
		basins[nr] = append(basins[nr], pos)
	} else {
		basins[len(basins)+1] = [][2]int{pos}
	}
}

var basins = map[int][][2]int{}

func findBasinNr(pos [2]int) int {
	basinNr := 0
	for x, b := range basins {
		if contains(b, pos) {
			basinNr = x
		}
	}
	return basinNr
}

func contains(basin [][2]int, pos [2]int) bool {
	for _, i := range basin {
		if i == pos {
			return true
		}
	}
	return false
}
