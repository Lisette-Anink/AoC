package day_24

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"lisette.anink/aoc/utils"
)

type ALU struct {
	w int
	x int
	y int
	z int
}

func MONADcalc(lines []string) int {
	start := 99999999999999
	for i := start; i > 99999999099990; i-- {
		inputStr := strings.Split(fmt.Sprint(i), "")
		input := utils.ConvertToInt(inputStr)
		valid := run(lines, input)
		if valid {
			return i
		}
	}
	return 0
}

func run(lines []string, input []int) bool {
	alu := &ALU{}

	inCount := 0
	for _, line := range lines {
		if len(line) > 0 {
			instr := strings.Fields(line)
			switch instr[0] {
			case "inp":
				alu.inp(instr[1], input[inCount])
				inCount++
			case "add":
				alu.add(instr[1], instr[2])
			case "mul":
				alu.mul(instr[1], instr[2])
			case "div":
				alu.div(instr[1], instr[2])
			case "mod":
				alu.mod(instr[1], instr[2])
			case "eql":
				alu.eql(instr[1], instr[2])
			}
		}
	}
	// fmt.Println(*alu)
	return alu.z == 0
}

// inp a - Read an input value and write it to variable a.
func (alu *ALU) inp(a string, in int) {
	alu.writeToPos(in, a)
}

// add a b - Add the value of a to the value of b, then store the result in variable a.
func (alu *ALU) add(a, b string) {
	alu.writeToPos(alu.readPos(a)+alu.readPos(b), a)
}

// mul a b - Multiply the value of a by the value of b, then store the result in variable a.
func (alu *ALU) mul(a, b string) {
	alu.writeToPos(alu.readPos(a)*alu.readPos(b), a)
}

// div a b - Divide the value of a by the value of b, truncate the result to an integer, then store the result in variable a. (Here, "truncate" means to round the value toward zero.)
func (alu *ALU) div(a, b string) {
	v := math.Round(float64(alu.readPos(a) / alu.readPos(b)))
	alu.writeToPos(int(v), a)
}

// mod a b - Divide the value of a by the value of b, then store the remainder in variable a. (This is also called the modulo operation.)
func (alu *ALU) mod(a, b string) {
	alu.writeToPos(alu.readPos(a)%alu.readPos(b), a)
}

// eql a b - If the value of a and b are equal, then store the value 1 in variable a. Otherwise, store the value 0 in variable a.
func (alu *ALU) eql(a, b string) {
	if alu.readPos(a) == alu.readPos(b) {
		alu.writeToPos(1, a)
	} else {
		alu.writeToPos(0, a)
	}
}

func (alu *ALU) writeToPos(val int, pos string) {
	switch pos {
	case "w":
		alu.w = val
	case "x":
		alu.x = val
	case "y":
		alu.y = val
	case "z":
		alu.z = val
	}
}

func (alu *ALU) readPos(pos string) int {
	switch pos {
	case "w":
		return alu.w
	case "x":
		return alu.x
	case "y":
		return alu.y
	case "z":
		return alu.z
	default:
		i, _ := strconv.Atoi(pos)
		return i
	}
}
