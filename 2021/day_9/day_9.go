package day_9

import (
	"sort"
	"strconv"
	"strings"

	"lisette.anink/aoc/utils"
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
			totalRisk += val + 1
		}
	}
	return totalRisk
}

func solvePart2(basinSizes map[int]int) int {
	a := []int{}
	for _, v := range basinSizes {
		a = append(a, v)
	}
	sort.Ints(a)
	utils.ReverseInt(a)

	return utils.Product(a[:3])
}

func findBasin(floormap map[[2]int]int) (basinSizes map[int]int) {
	for pos := range floormap {
		checkBasin(floormap, pos)
	}
	basinSizes = map[int]int{}
	for b, p := range basins {
		basinSizes[b] = len(p)
	}
	return
}

func checkBasin(floormap map[[2]int]int, pos [2]int) {
	if floormap[pos] != 9 {
		nr := findBasinNr(pos)
		processPosition(floormap, pos, nr)
	}
	delete(floormap, pos)
}

func processPosition(floormap map[[2]int]int, pos [2]int, basinNr int) {
	if floormap[pos] != 9 {
		if !contains(basins[basinNr], pos) {
			addToBasin(pos, basinNr)
		}
		delete(floormap, pos)
		checkBasinNeighbours(floormap, pos, basinNr)
	} else {
		delete(floormap, pos)
	}
}

func checkBasinNeighbours(floormap map[[2]int]int, pos [2]int, basinNr int) {
	neighLeft := [2]int{pos[0] - 1, pos[1]}
	if _, ok := floormap[neighLeft]; ok {
		processPosition(floormap, neighLeft, basinNr)
	}
	neighRight := [2]int{pos[0] + 1, pos[1]}
	if _, ok := floormap[neighRight]; ok {
		processPosition(floormap, neighRight, basinNr)
	}
	neighTop := [2]int{pos[0], pos[1] - 1}
	if _, ok := floormap[neighTop]; ok {
		processPosition(floormap, neighTop, basinNr)
	}
	neighBottom := [2]int{pos[0], pos[1] + 1}
	if _, ok := floormap[neighBottom]; ok {
		processPosition(floormap, neighBottom, basinNr)
	}
}

func addToBasin(pos [2]int, nr int) {
	basins[nr] = append(basins[nr], pos)
}

var basins = map[int][][2]int{}

func findBasinNr(pos [2]int) int {
	basinNr := 0
	for x, b := range basins {
		if contains(b, pos) {
			basinNr = x
		}
	}
	if basinNr == 0 {
		basinNr = len(basins) + 1
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
