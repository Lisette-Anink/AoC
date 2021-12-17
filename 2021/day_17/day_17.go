package day_17

import (
	"fmt"
	"regexp"

	"lisette.anink/aoc/utils"
)

type targetArea struct {
	xMin int
	yMin int
	xMax int
	yMax int
}
type probe struct {
	xPos int
	yPos int
	xVel int
	yVel int
}

func parseInput(line string) targetArea {
	r := regexp.MustCompile(`target area\: x\=(.*)\.\.(.*), y=(.*)\.\.(.*)`)
	fields := r.FindStringSubmatch(line)
	fmt.Println(fields)
	n := utils.ConvertToInt(fields[1:])
	return targetArea{
		xMin: n[0],
		yMin: n[2],
		xMax: n[1],
		yMax: n[3],
	}
}
func findHighest(input string, yMax int) int {
	h := 0
	target := parseInput(input)
	for x := 0; x < target.xMax; x++ {
		for y := 0; y < yMax; y++ {
			newH := shootProbe(target, x, y)
			if newH > h {
				h = newH
			}
		}
	}
	return h
}
func countHits(input string, yMax int) int {
	count := 0
	target := parseInput(input)
	for x := 0; x <= target.xMax; x++ {
		for y := target.yMin; y < yMax; y++ {
			newH := shootProbe(target, x, y)
			if newH != -1 {
				// fmt.Printf("x: %d, y: %d, h: %d\n", x, y, newH)
				count++
			}
		}
	}
	return count
}

func shootProbe(target targetArea, x, y int) int {
	// create probe at pos 0,0
	p := &probe{xVel: x, yVel: y}
	// shoot probe until target
	h := 0
	for !p.inTarget(target) {
		p.trajectoryStep()
		if p.yPos > h {
			h = p.yPos
		}
		if p.pastTarget(target) {
			h = -1
			break
		}
	}
	return h
}

func (p *probe) inTarget(t targetArea) bool {
	return p.xPos >= t.xMin && p.xPos <= t.xMax && p.yPos >= t.yMin && p.yPos <= t.yMax
}

func (p *probe) pastTarget(t targetArea) bool {
	return p.xPos > t.xMax || p.yPos < t.yMin
}

func (p *probe) trajectoryStep() {
	//The probe's x position increases by its x velocity.
	p.xPos += p.xVel
	// The probe's y position increases by its y velocity.
	p.yPos += p.yVel
	// Due to drag, the probe's x velocity changes by 1 toward the value 0; that is, it decreases by 1 if it is greater than 0, increases by 1 if it is less than 0, or does not change if it is already 0.
	if p.xVel > 0 {
		p.xVel--
	} else if p.xVel < 0 {
		p.xVel++
	}
	// Due to gravity, the probe's y velocity decreases by 1.
	p.yVel--
}
