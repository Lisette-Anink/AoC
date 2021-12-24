package day_22

import (
	"regexp"

	"lisette.anink/aoc/utils"
)

func parseInput(lines []string) map[int][7]int {
	r := regexp.MustCompile(`(on|off) x\=(\-?\d+)\.\.(\-?\d+),y=(\-?\d+)\.\.(\-?\d+),z\=(\-?\d+)\.\.(\-?\d+)`)
	rebootSteps := map[int][7]int{}
	for i, l := range lines {
		if len(l) > 0 {
			matches := r.FindStringSubmatch(l)
			nr := utils.ConvertToInt(matches[2:])
			onoff := 0
			if matches[1] == "on" {
				onoff = 1
			} else {
				onoff = -1

			}
			rebootSteps[i] = [7]int{onoff, nr[0], nr[1], nr[2], nr[3], nr[4], nr[5]}
		}
	}
	return rebootSteps
}

func rebootInitiation(rebootSteps map[int][7]int) int {
	cubes := map[[3]int]bool{}
	for step := 0; step <= len(rebootSteps); step++ {
		// for step, instr := range rebootSteps {
		instr := rebootSteps[step]
		if inRangeInstr(instr, -500, 500) {
			// fmt.Println(step, instr)
			if instr[0] == 1 { //on
				for x := instr[1]; x <= instr[2]; x++ {
					for y := instr[3]; y <= instr[4]; y++ {
						for z := instr[5]; z <= instr[6]; z++ {
							cubes[[3]int{x, y, z}] = true
						}
					}
				}
			} else if instr[0] == -1 { //off
				for x := instr[1]; x <= instr[2]; x++ {
					for y := instr[3]; y <= instr[4]; y++ {
						for z := instr[5]; z <= instr[6]; z++ {
							cubes[[3]int{x, y, z}] = false
						}
					}
				}
			}
		}
	}
	// fmt.Println(cubes)
	return countCubes(cubes)
}

func countCubes(cubes map[[3]int]bool) int {
	count := 0
	for c, b := range cubes {
		if b && inRange(c, -50, 50) {
			count++
		}
	}
	return count
}

func inRange(c [3]int, min, max int) bool {
	return c[0] >= min && c[0] <= max && c[1] >= min && c[1] <= max && c[2] >= min && c[2] <= max
}
func inRangeInstr(c [7]int, min, max int) bool {
	return c[1] >= min && c[2] <= max && c[3] >= min && c[4] <= max && c[5] >= min && c[6] <= max
}

func reboot(rebootSteps map[int][7]int) int {
	cubes := map[[3]int]bool{}
	// changes
	for step := 0; step <= len(rebootSteps); step++ {
		instr := rebootSteps[step]
		if instr[0] == 1 { //on
			x := instr[2] - instr[1]
			y := instr[4] - instr[3]
		}
	}
	// fmt.Println(cubes)
	return countCubes(cubes)
}
