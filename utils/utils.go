package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertBinaryToInt(binaryString string) int64 {
	int, err := strconv.ParseInt(binaryString, 2, 0)
	if err != nil {
		panic(err)
	}
	return int
}

// ConvertToInt convert a slice of strings to slice of ints
func ConvertToInt(lines []string) []int {
	ar := []int{}
	for _, line := range lines {
		if len(line) > 0 {
			i, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("ERROR: ", err)
			}
			ar = append(ar, i)
		}
	}
	return ar
}

// Reverse order of array
func Reverse(array []string) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

// Difference returns the diff base - substract of two
// string slices
func Difference(base, substract []string) []string {
	basemap := map[string]bool{}
	for _, e := range base {
		basemap[e] = true
	}

	for _, s := range substract {
		for k, _ := range basemap {
			if s == k {
				delete(basemap, k)
			}
		}
	}
	keys := make([]string, 0, len(basemap))
	for k := range basemap {
		keys = append(keys, k)
	}
	return keys
}

// returns true if collection contains all elements of test
func IncludesAll(collection, test []string) bool {
	total := len(test)
	count := 0
	for i := 0; i < len(collection); i++ {
		for _, t := range test {
			if t == collection[i] {
				count++
				if count == total {
					return true
				}
			}
		}
	}
	return false
}

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
func Product(array []int) int {
	result := array[0]
	for _, v := range array[1:] {
		result *= v
	}
	return result
}
func ReverseInt(array []int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

func ParseIntMap(lines []string) map[[2]int]int {
	var intmap = map[[2]int]int{}
	for x, line := range lines {
		if len(line) > 0 {
			parts := strings.Split(line, "")
			for y, p := range parts {
				i, _ := strconv.Atoi(p)
				intmap[[2]int{x, y}] = i
			}
		}
	}
	return intmap
}
