package day_11

import "fmt"

func octopusEnergyModel(octomap map[[2]int]int, steps int) int {
	flashes := 0
	diff := 0
	for i := 0; i < steps; i++ {
		for pos := range octomap {
			increase(octomap, pos)
		}
		flashes, diff = resetOctomap(octomap, flashes)
		if diff == len(octomap) {
			fmt.Println("step ", i+1)
			break
		}
	}
	return flashes
}

func resetOctomap(octomap map[[2]int]int, flashes int) (int, int) {
	// fmt.Println(octomap)
	old := flashes
	for pos, energy := range octomap {
		if energy > 9 {
			flashes += 1
			octomap[pos] = 0
			// fmt.Println(pos, flashes)
		}
	}
	// fmt.Println(octomap)
	return flashes, flashes - old
}

func increase(octomap map[[2]int]int, pos [2]int) {
	if _, ok := octomap[pos]; ok {
		if octomap[pos] == 9 {
			octomap[pos] += 1
			increaseNeighbours(octomap, pos)
		} else {
			octomap[pos] += 1
		}
	}
}

func increaseNeighbours(octomap map[[2]int]int, pos [2]int) {
	for x := pos[0] - 1; x <= pos[0]+1; x++ {
		for y := pos[1] - 1; y <= pos[1]+1; y++ {

			if !(x == pos[0] && y == pos[1]) {
				// fmt.Println("yes")
				increase(octomap, [2]int{x, y})
			}
		}
	}
}

// func printOctomap(octomap map[[2]int]int) {
// 	last := len
// }
