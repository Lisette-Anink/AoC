package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/gobuffalo/envy"
)

// SplitInLines split byte input  on '\n' in lines strings
func SplitInLines(input []byte) []string {
	return strings.Split(string(input), "\n")
}

// SplitInMultiLines split byte input on '\n\n' in lines strings
func SplitInMultiLines(input []byte) []string {
	return strings.Split(string(input), "\n\n")
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

// ImportFileLines get input from file
func ImportFileLines(path string) []byte {
	dat, err := ioutil.ReadFile(path)
	check(err)

	return dat
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetInputFromAoC get request directly from adventofcode
func GetInputFromAoC(day string) []byte {
	return ImportFromAoC("2020", day)
}

func ImportFromAoC(year, day string) []byte {
	envy.Load("./../../.env")
	filepath := fmt.Sprintf("input_%s_%s", year, day)
	body := []byte{}
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)
		req, err := http.NewRequest("GET", url, nil)
		check(err)

		sessionToken, err := envy.MustGet("session")
		check(err)

		// fmt.Println("call AoC ", sessionToken)
		req.AddCookie(&http.Cookie{Name: "session", Value: sessionToken})

		client := &http.Client{}
		resp, err := client.Do(req)
		check(err)
		defer resp.Body.Close()
		// fmt.Println("call AoC")
		body, err = ioutil.ReadAll(resp.Body)
		check(err)

		os.WriteFile(filepath, body, 0644)
	} else {
		// fmt.Println("read from file")
		body = ImportFileLines(filepath)
	}

	// fmt.Println("response", string(body))
	return body
}

// Reverse order of array
func Reverse(array []string) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
