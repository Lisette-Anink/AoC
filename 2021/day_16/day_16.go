package day_16

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"lisette.anink/aoc/utils"
)

func parseInput(line string) string {
	binary := ""
	chars := strings.Split(line, "")
	for _, c := range chars {
		i, _ := strconv.ParseUint(c, 16, 32)
		binary = fmt.Sprintf("%s%.4b", binary, i)
	}
	return binary
}

type packet struct {
	version  int64
	typeID   int64 //4:literal
	value    string
	valueInt int64
	subP     []*packet
}

func readPackets(in string) (*packet, string) {
	// fmt.Println("in", in)
	if len(in) < 6 {
		return nil, in
	}
	version := utils.ConvertBinaryToInt(in[:3])
	typeID := utils.ConvertBinaryToInt(in[3:6])
	remainder := ""
	pt := &packet{
		version: version,
		typeID:  typeID,
	}
	remainder = in[6:]
	// fmt.Println("p", pt.version, pt.typeID)
	if typeID == 4 { //literal
		value := ""
		for len(remainder) > 4 {
			re := regexp.MustCompile("(.)(.{4})(.*)")
			groups := re.FindStringSubmatch(remainder)
			// fmt.Println("g", groups)
			remainder = groups[3]

			value = fmt.Sprintf("%s%s", value, groups[2])
			if string(groups[1][0]) == "0" {
				break
			}
		}
		// fmt.Println("v", value)
		pt.value = value
		pt.valueInt = utils.ConvertBinaryToInt(value)
	} else {
		lenTypeID := in[6]
		switch string(lenTypeID) {
		case "0": //the next 15 bits are a number that represents the total length in bits of the sub-packets contained by this packet.
			totalLengInBits := utils.ConvertBinaryToInt(in[7:22])

			remainder = in[22 : 22+totalLengInBits]
			rest := in[22+totalLengInBits:]
			// fmt.Println("len totalLengInBits", totalLengInBits, remainder)
			for len(remainder) > 6 {
				// fmt.Println("r 3", remainder, len(remainder))
				var subP = &packet{}
				subP, remainder = readPackets(remainder)
				pt.subP = append(pt.subP, subP)
			}
			remainder = rest
		case "1": //the next 11 bits are a number that represents the number of sub-packets immediately contained by this packet
			totalNrSubpack := utils.ConvertBinaryToInt(in[7:18])
			remainder = in[18:]
			// fmt.Println("len subp", totalNrSubpack, remainder)
			for i := 0; i < int(totalNrSubpack); i++ {
				// fmt.Println("i", i)
				var subP = &packet{}
				subP, remainder = readPackets(remainder)
				if subP != nil {
					pt.subP = append(pt.subP, subP)
				}
				// fmt.Println("r5", remainder, len(remainder))
			}
		}
	}
	// fmt.Println(pt, remainder)
	return pt, remainder
}

func addVersionNr(pIn packet, n int64) int64 {
	v := n
	v += pIn.version
	// fmt.Println("packet", pIn)
	for _, p := range pIn.subP {

		v = addVersionNr(*p, v)
	}
	return v
}

func calculate(pIn packet) int64 {
	v := int64(0)
	switch pIn.typeID {
	case 0: //are sum packets
		fmt.Printf("case 0: ")
		for _, p := range pIn.subP {
			v += calculate(*p)
		}
	case 1: //are product packets
		fmt.Printf("case 1: ")
		v = 1
		if len(pIn.subP) == 1 {
			v = calculate(*pIn.subP[0])
		} else {
			for _, p := range pIn.subP {
				v *= calculate(*p)
			}
		}
	case 2: //are minimum packets
		fmt.Printf("case 2: ")
		min := int64(math.MaxInt64)
		for _, p := range pIn.subP {
			val := calculate(*p)
			if min > val {
				min = val
			}
		}
		v = min
	case 3: //are maximum packets
		fmt.Printf("case 3: ")
		for _, p := range pIn.subP {
			val := calculate(*p)
			if v < val {
				v = val
			}
		}
	case 4:
		fmt.Printf("case 4: ")
		v = pIn.valueInt
	case 5: //are greater than packets
		fmt.Printf("case 5: ")
		if calculate(*pIn.subP[0]) > calculate(*pIn.subP[1]) {
			v = 1
		} else {
			v = 0
		}
	case 6: //are less than packets
		fmt.Printf("case 6: ")
		if calculate(*pIn.subP[0]) < calculate(*pIn.subP[1]) {
			v = 1
		} else {
			v = 0
		}

	case 7: //are equal to packets
		fmt.Printf("case 7: ")
		if calculate(*pIn.subP[0]) == calculate(*pIn.subP[1]) {
			v = 1
		} else {
			v = 0
		}

	}
	fmt.Println(v)
	return v
}

// Packets with type ID 0 are sum packets - their value is the sum of the values of their sub-packets. If they only have a single sub-packet, their value is the value of the sub-packet.
// Packets with type ID 1 are product packets - their value is the result of multiplying together the values of their sub-packets. If they only have a single sub-packet, their value is the value of the sub-packet.
// Packets with type ID 2 are minimum packets - their value is the minimum of the values of their sub-packets.
// Packets with type ID 3 are maximum packets - their value is the maximum of the values of their sub-packets.
// Packets with type ID 5 are greater than packets - their value is 1 if the value of the first sub-packet is greater than the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
// Packets with type ID 6 are less than packets - their value is 1 if the value of the first sub-packet is less than the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
// Packets with type ID 7 are equal to packets - their value is 1 if the value of the first sub-packet is equal to the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
