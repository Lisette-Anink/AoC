package day_6

import (
	"strings"

	"lisette.anink/aoc/utils"
)

type population struct {
	day        int
	sizePerAge map[int]int
}

func parseInput(in string) population {
	startPop := strings.Split(in, ",")
	startAge := utils.ConvertToInt(startPop)
	pop := population{day: 0, sizePerAge: map[int]int{}}
	for _, age := range startAge {
		pop.sizePerAge[age]++
	}
	return pop
}

func (pop *population) oneDay() {
	pop.day++
	newPop := map[int]int{}
	for age := 8; age >= 0; age-- {
		switch age {
		case 0:
			newPop[age+6] += pop.sizePerAge[age]
			newPop[8] = pop.sizePerAge[age]
		default:
			newPop[age-1] = pop.sizePerAge[age]
		}
	}
	pop.sizePerAge = newPop
}

func (pop *population) simXDays(x int) {
	for i := 0; i < x; i++ {
		pop.oneDay()
	}
}

func (pop *population) totalSize() (total int) {
	for age := 8; age >= 0; age-- {
		total += pop.sizePerAge[age]
	}
	return total
}
