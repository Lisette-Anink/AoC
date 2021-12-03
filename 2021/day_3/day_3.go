package day_3

import (
	"fmt"
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

func findOxCoRating(numbers []string) (oxGen int64, coGen int64) {
	oxGenRate := reduceArray(numbers, 0, "ox")
	coGenRate := reduceArray(numbers, 0, "co")
	fmt.Printf("o: %s, c: %s\n", oxGenRate, coGenRate)

	oxGen = convertBinaryToInt(oxGenRate)
	coGen = convertBinaryToInt(coGenRate)
	return
}

func reduceArray(arr []string, testPos int, t string) string {
	bits := splitInChars(arr)
	tr := transpose(bits)
	least, most := leastMost(tr[testPos])
	rest := []string{}
	if t == "ox" {
		rest = findAll(arr, []rune(most)[0], testPos)
	} else {
		rest = findAll(arr, []rune(least)[0], testPos)
	}

	if len(rest) > 1 {
		reduceArray(rest, testPos+1, t)
	}
	fmt.Println(rest[0])
	return rest[0]
}

func findAll(arr []string, char rune, pos int) (out []string) {
	for _, n := range arr {
		if len(n) > 0 {
			if []rune(n)[pos] == char {
				out = append(out, n)
			}
		}
	}
	// fmt.Println(out)
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
	// fmt.Printf("0: %d, 1: %d\n", zero, one)
	if zero == one {
		return "1", "1"
	} else if zero < one {
		return "0", "1"
	}
	return "1", "0"
}

func convertBinaryToInt(binaryString string) int64 {
	int, err := strconv.ParseInt(binaryString, 2, 0)
	if err != nil {
		panic(err)
	}
	return int
}
