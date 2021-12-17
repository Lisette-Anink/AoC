package day_14

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func parseInput(parts []string) ([]string, map[[2]string]string) {
	// part 0 is polymer
	polymer := strings.Split(parts[0], "")
	// part 1 is pair insertion rules
	insertionRule := map[[2]string]string{}
	rules := strings.Split(parts[1], "\n")
	r := regexp.MustCompile(`(\w)(\w) -> (\w)`)
	for _, rule := range rules {
		if len(rule) > 0 {
			m := r.FindStringSubmatch(rule)
			insertionRule[[2]string{m[1], m[2]}] = m[3]
		}
	}
	return polymer, insertionRule
}

func pairInsertion(polymer []string, insertionRules map[[2]string]string) []string {
	for i := 0; i < len(polymer)-1; i++ {
		// fmt.Println("here", polymer, len(polymer), i)
		if i < len(polymer)-1 {
			if val, ok := insertionRules[[2]string{polymer[i], polymer[i+1]}]; ok {
				// fmt.Println(polymer, len(polymer))
				polymer = insert(polymer, i+1, val)
				i++
				// fmt.Println(polymer, len(polymer))
			}
		}
	}
	return polymer
}

func insert(a []string, index int, value string) []string {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func insertSteps(polymer []string, insertionRules map[[2]string]string, steps int) []string {
	for i := 0; i < steps; i++ {
		polymer = pairInsertion(polymer, insertionRules)
	}
	return polymer
}

func count(polymer []string) (map[string]int, int, int) {
	c := map[string]int{}
	for _, el := range polymer {
		c[el] += 1
	}
	min := math.MaxInt64
	max := 0
	// fmt.Println(c)
	for i := range c {
		if c[i] < min {
			min = c[i]
		}
		if c[i] > max {
			max = c[i]
		}
		// fmt.Println(min, max)
	}
	// fmt.Println(c)
	return c, min, max
}

// pair insertion NN -> C will add NC and CN
//  insRule = map[[2]string][2][2]string > {N,N}: {{N,C}, {C,N}}
// type polymer str
var last = ""

func parseInputPairwise(parts []string) (map[[2]string]int, map[[2]string][2][2]string) {
	polymer := map[[2]string]int{}
	insRule := map[[2]string][2][2]string{}
	p := strings.Split(parts[0], "")
	for i := range p {
		if i < len(p)-1 {
			polymer[[2]string{p[i], p[i+1]}] += 1
			last = p[i+1]
		}
	}
	rules := strings.Split(parts[1], "\n")
	r := regexp.MustCompile(`(\w)(\w) -> (\w)`)
	for _, rule := range rules {
		if len(rule) > 0 {
			m := r.FindStringSubmatch(rule)
			insRule[[2]string{m[1], m[2]}] = [2][2]string{{m[1], m[3]}, {m[3], m[2]}}
		}
	}
	return polymer, insRule
}

func pairWiseInsertion(polymer map[[2]string]int, insertionRules map[[2]string][2][2]string) map[[2]string]int {
	newPol := map[[2]string]int{}
	for key, val := range polymer {
		if insPair, ok := insertionRules[key]; ok {
			newPol[insPair[0]] += val
			newPol[insPair[1]] += val
		}
	}
	return newPol
}

func insertPairwiseSteps(polymer map[[2]string]int, insertionRules map[[2]string][2][2]string, steps int) map[[2]string]int {
	for i := 0; i < steps; i++ {
		polymer = pairWiseInsertion(polymer, insertionRules)
	}
	return polymer
}

func countPair(polymer map[[2]string]int) (map[string]int, int, int) {
	fmt.Println(last)
	c := map[string]int{last: 1}
	for key, val := range polymer {
		c[key[0]] += val
	}
	min := math.MaxInt64
	max := 0
	// fmt.Println(c)
	for i := range c {
		if c[i] < min {
			min = c[i]
		}
		if c[i] > max {
			max = c[i]
		}
		// fmt.Println(min, max)
	}
	// fmt.Println(c)
	return c, min, max
}
