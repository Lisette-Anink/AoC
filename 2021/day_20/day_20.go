package day_20

import (
	"fmt"
	"math"
	"strings"

	"lisette.anink/aoc/utils"
)

func parseInput(lines []string) (imageEnhancementAlgorithm map[int]string, inputImage map[[2]int]string) {
	imageEnhancementAlgorithm = map[int]string{}
	pos := strings.Split(lines[0], "")
	for i, p := range pos {
		imageEnhancementAlgorithm[i] = p
	}

	inputImage = map[[2]int]string{}
	horiz := strings.Split(lines[1], "\n")
	for x, l := range horiz {
		pix := strings.Split(l, "")
		for y, p := range pix {
			inputImage[[2]int{x, y}] = p
		}
	}
	return
}

func countLightPixAfterXSteps(imageEnhancementAlgorithm map[int]string, inputImage map[[2]int]string, steps int) int {
	outputImage := inputImage
	// fmt.Println(countLightPix(outputImage))
	for i := 1; i <= steps; i++ {
		outputImage = enhanceImage(imageEnhancementAlgorithm, outputImage, oddStep(i))
		// fmt.Println(countLightPix(outputImage))
		// printImage(outputImage)
	}
	printImage(outputImage)
	count := countLightPix(outputImage)
	return count
}
func oddStep(i int) bool {
	return i%2 == 0
}

func enhanceImage(imageEnhancementAlgorithm map[int]string, inputImage map[[2]int]string, odd bool) map[[2]int]string {
	output := map[[2]int]string{}
	min, max := findMinMax(inputImage)
	for x := min[0]; x <= max[0]; x++ {
		for y := min[1]; y <= max[1]; y++ {
			output[[2]int{x, y}] = valFromWindow([2]int{x, y}, inputImage, imageEnhancementAlgorithm, odd)
		}
	}
	return output
}

func findMinMax(inputImage map[[2]int]string) ([2]int, [2]int) {
	min := [2]int{math.MaxInt64, math.MaxInt64}
	max := [2]int{}
	for pos := range inputImage {
		if pos[0] <= min[0] {
			min[0] = pos[0] - 1
		}
		if pos[1] <= min[1] {
			min[1] = pos[1] - 1
		}
		if pos[0] >= max[0] {
			max[0] = pos[0] + 1
		}
		if pos[1] >= max[1] {
			max[1] = pos[1] + 1
		}
	}
	return min, max
}

func valFromWindow(pos [2]int, inputImage map[[2]int]string, imageEnhancementAlgorithm map[int]string, odd bool) string {
	newVal := ""
	for x := pos[0] - 1; x <= pos[0]+1; x++ {
		for y := pos[1] - 1; y <= pos[1]+1; y++ {
			if val, ok := inputImage[[2]int{x, y}]; ok {
				newVal += val
			} else {
				if odd && imageEnhancementAlgorithm[0] == "#" {
					newVal += "#"
				} else {
					newVal += "."
				}
			}
		}
	}
	algPos := convertToAlgPos(newVal)
	newVal = imageEnhancementAlgorithm[algPos]
	return newVal
}

func convertToAlgPos(newVal string) int {
	result := 0
	binary := ""
	for _, p := range newVal {
		if p == '.' {
			binary += "0"
		} else {
			binary += "1"
		}
	}
	result = int(utils.ConvertBinaryToInt(binary))
	return result
}

func countLightPix(inputImage map[[2]int]string) int {
	count := 0
	for _, val := range inputImage {
		if val == "#" {
			count++
		}
	}
	return count
}

func printImage(inputImage map[[2]int]string) {
	min, max := findMinMax(inputImage)
	for x := min[0]; x <= max[0]; x++ {
		for y := min[1]; y <= max[1]; y++ {
			fmt.Printf(inputImage[[2]int{x, y}])
		}
		fmt.Println()
	}
}
