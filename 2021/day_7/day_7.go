package day_7

import (
	"math"
	"strings"

	"lisette.anink/aoc/utils"
)

func parseInput(in string) []int {
	n := strings.Split(in, ",")
	crabs := utils.ConvertToInt(n)
	return crabs
}

func calFuel(crabs []int, pos int) (total float64) {
	for _, c := range crabs {
		f := math.Abs(float64(c) - float64(pos))
		total += f
	}
	return
}
func calBruteForceFuel(crabs []int, pos int) (total float64) {
	for _, c := range crabs {
		total += calBruteForce(float64(c), float64(pos))
	}
	return
}
func calGaussFuel(crabs []int, pos int) (total float64) {
	for _, c := range crabs {
		total += calcGauss(float64(c), float64(pos))
	}
	return
}
func calBruteForce(c, p float64) float64 {
	b := 0
	f := math.Abs(float64(c) - float64(p))
	for i := 0; i < int(f); i++ {
		b += i + 1
	}
	return float64(b)
}
func calcGauss(c, p float64) float64 {
	div := math.Abs(c - p)
	t := (div / 2) * (div + 1)
	return t
}

func calMinFuel(crabs []int) (float64, int) {
	fuel := map[int]float64{}
	for p := 0; p < maxInt(crabs); p++ {
		fuel[p] = calFuel(crabs, p)
	}
	min, pos := minFloat(fuel)
	return min, pos
}

func calMinFuelPart2(crabs []int) (float64, int) {
	fuel := map[int]float64{}
	for p := 0; p < maxInt(crabs); p++ {
		fuel[p] = calBruteForceFuel(crabs, p)
	}
	min, pos := minFloat(fuel)
	return min, pos
}

func minFloat(fuel map[int]float64) (m float64, p int) {
	m = math.MaxFloat64
	for k, v := range fuel {
		if v < m {
			m = v
			p = k
		}
	}
	return
}

func maxInt(a []int) int {
	m := a[0]
	for _, i := range a {
		if i > m {
			m = i
		}
	}
	return m
}
