package day_2

import (
	"strconv"
	"strings"
)

func calPosition(directions []string) (h int, d int) {
	for _, l := range directions {
		parts := strings.Split(l, " ")
		if len(parts) > 1 {
			n, _ := strconv.Atoi(parts[1])
			// fmt.Println(parts)
			switch parts[0] {
			case "forward":
				h += n
			case "up":
				d -= n
			case "down":
				d += n
			}
		}
	}

	return
}

// down X increases your aim by X units.
// up X decreases your aim by X units.
// forward X does two things:
//   It increases your horizontal position by X units.
//   It increases your depth by your aim multiplied by X.

func calPositionAim(directions []string) (h int, d int) {
	aim := 0
	for _, l := range directions {
		parts := strings.Split(l, " ")
		if len(parts) > 1 {
			n, _ := strconv.Atoi(parts[1])
			// fmt.Println(parts)
			switch parts[0] {
			case "forward":
				h += n
				d += n * aim
			case "up":
				aim -= n
			case "down":
				aim += n
			}
		}
	}

	return
}
