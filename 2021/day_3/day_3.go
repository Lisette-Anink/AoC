package day_3

import (
	"strconv"
	"strings"
)

func findGammaEps(numbers []string) (g int64, e int64) {
	bits := splitInChars(numbers)
	tr := transpose(bits)
	gBin, eBin := "", ""
	for i := 0; i < len(tr); i++ {
		least, most := leastMost(tr[i])
		gBin += most
		eBin += least
	}
	g = convertBinaryToInt(gBin)
	e = convertBinaryToInt(eBin)
	return
}

func splitInChars(numbers []string) (out [][]string) {
	for _, n := range numbers {
		if len(n) > 0 {
			parts := strings.Split(n, "")
			out = append(out, parts)
		}
	}
	return
}

func transpose(array [][]string) map[int][]string {
	transposed := map[int][]string{}
	for i := 0; i < len(array[0]); i++ {
		transposed[i] = []string{}
		for _, l := range array {
			transposed[i] = append(transposed[i], l[i])
		}
	}
	return transposed
}

func leastMost(array []string) (string, string) {
	zero, one := 0, 0
	for _, n := range array {
		switch n {
		case "0":
			zero++
		case "1":
			one++
		}
	}
	if zero > one {
		return "1", "0"
	}
	return "0", "1"
}

func convertBinaryToInt(binaryString string) int64 {
	int, err := strconv.ParseInt(binaryString, 2, 0)
	if err != nil {
		panic(err)
	}
	return int
}
