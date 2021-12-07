package utils

import (
	"fmt"
	"strconv"
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
