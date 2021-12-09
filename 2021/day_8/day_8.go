package day_8

import (
	"sort"
	"strconv"
	"strings"

	"lisette.anink/aoc/utils"
)

func parseInput(lines []string) (inputDigits, outputDigits []string) {
	for _, line := range lines {
		if len(line) > 0 {
			parts := strings.Split(line, "|")
			inputDigits = append(inputDigits, strings.Fields(parts[0])...)
			outputDigits = append(outputDigits, strings.Fields(parts[1])...)
		}
	}
	return
}

func parseInput2(line string) (inputDigits, outputDigits []string) {
	if len(line) > 0 {
		parts := strings.Split(line, "|")
		inputDigits = strings.Fields(parts[0])
		outputDigits = strings.Fields(parts[1])
	}
	return
}

func countNumbers(in []string) (count int) {
	for _, i := range in {
		switch len(i) {
		case 2: // 1
			count++
		case 3: // 7
			count++
		case 4: // 4
			count++
		case 7: // 8
			count++
		}
	}
	return
}

func identifyNumbers(in []string) (iden map[string][]string) {
	inSplit := [][]string{}
	for _, i := range in {
		el := strings.Split(i, "")
		inSplit = append(inSplit, el)
	}
	iden = map[string][]string{}
	for _, i := range inSplit {
		switch len(i) {
		case 2: // 1
			iden["1"] = i
		case 3: // 7
			iden["7"] = i
		case 4: // 4
			iden["4"] = i
		case 7: // 8
			iden["8"] = i
		}
	}
	midTopLeft := utils.Difference(iden["4"], iden["1"])
	for _, i := range inSplit {
		if len(i) == 6 {
			if utils.IncludesAll(i, iden["4"]) {
				iden["9"] = i
			} else if utils.IncludesAll(i, iden["7"]) {
				iden["0"] = i
			} else {
				iden["6"] = i
			}
		}
		if len(i) == 5 {
			if utils.IncludesAll(i, iden["7"]) {
				iden["3"] = i
			} else if utils.IncludesAll(i, midTopLeft) {
				iden["5"] = i
			} else {
				iden["2"] = i
			}
		}
	}
	return
}

func valDigits(iden map[string][]string, out []string) int {
	val := ""
	inSplit := [][]string{}
	for _, i := range out {
		el := strings.Split(i, "")
		inSplit = append(inSplit, el)
	}
	for _, i := range inSplit {
		for k, v := range iden {
			if equal(v, i) {
				val += k
			}
		}
	}
	v, _ := strconv.Atoi(val)
	return v
}

func equal(A, B []string) bool {
	if len(A) != len(B) {
		return false
	}
	sort.Strings(A)
	sort.Strings(B)
	for i, e := range A {
		if e != B[i] {
			return false
		}
	}
	return true
}
