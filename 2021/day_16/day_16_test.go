package day_16

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"lisette.anink/aoc/utils"
)

func useTestData() string {
	input := utils.ImportFileLines("test_input_16")
	return string(input)
}

func useInputData() string {
	input := utils.ImportFromAoC("2021", "16")
	return string(input)
}

func Test_parseInput(t *testing.T) {
	type args struct {
		lines string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{"0"}, "0000"},
		{"1", args{"1"}, "0001"},
		{"C", args{"C"}, "1100"},
		{"F", args{"F"}, "1111"},
		{"D2FE28", args{"D2FE28"}, "110100101111111000101000"},
		{"38006F45291200", args{"38006F45291200"}, "00111000000000000110111101000101001010010001001000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseInput(tt.args.lines)
			assert.Equal(t, tt.want, got)
		})
	}
}

// 0 = 0000
// 1 = 0001
// 2 = 0010
// 3 = 0011
// 4 = 0100
// 5 = 0101
// 6 = 0110
// 7 = 0111
// 8 = 1000
// 9 = 1001
// A = 1010
// B = 1011
// C = 1100
// D = 1101
// E = 1110
// F = 1111

func Test_readPackets(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name  string
		args  args
		want  *packet
		want2 string
	}{
		{"literal", args{"110100101111111000101000"}, &packet{6, 4, "011111100101", 2021, []*packet(nil)}, "000"},
		{"operator", args{"00111000000000000110111101000101001010010001001000000000"}, &packet{1, 6, "", 0, []*packet{
			{version: 6, typeID: 4, value: "1010", valueInt: 10, subP: []*packet(nil)},
			{version: 2, typeID: 4, value: "00010100", valueInt: 20, subP: []*packet(nil)},
		}}, ""},
		{"operator", args{"11101110000000001101010000001100100000100011000001100000"}, &packet{7, 3, "", 0, []*packet{
			{version: 2, typeID: 4, value: "0001", valueInt: 1, subP: []*packet(nil)},
			{version: 4, typeID: 4, value: "0010", valueInt: 2, subP: []*packet(nil)},
			{version: 1, typeID: 4, value: "0011", valueInt: 3, subP: []*packet(nil)}},
		}, "00000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := readPackets(tt.args.in)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want2, got2)
		})
	}
}

func Test_addVersionNr(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test 38006F45291200", args{"38006F45291200"}, 9},
		{"test EE00D40C823060", args{"EE00D40C823060"}, 14},
		{"test 8A004A801A8002F478", args{"8A004A801A8002F478"}, 16},
		{"test 620080001611562C8802118E34", args{"620080001611562C8802118E34"}, 12},
		{"test C0015000016115A2E0802F182340", args{"C0015000016115A2E0802F182340"}, 23},
		{"test A0016C880162017C3686B18A3D4780", args{"A0016C880162017C3686B18A3D4780"}, 31},
		{"Part 1", args{useInputData()}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := parseInput(tt.args.in)
			p, _ := readPackets(in)
			got := addVersionNr(*p, 0)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_calculate(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test C200B40A82", args{"C200B40A82"}, 3},
		{"test 04005AC33890", args{"04005AC33890"}, 54},
		{"test 880086C3E88112", args{"880086C3E88112"}, 7},
		{"test CE00C43D881120", args{"CE00C43D881120"}, 9},
		{"test D8005AC2A8F0", args{"D8005AC2A8F0"}, 1},
		{"test F600BC2D8F", args{"F600BC2D8F"}, 0},
		{"test 9C005AC2F8F0", args{"9C005AC2F8F0"}, 0},
		{"test 9C0141080250320F1802104A08", args{"9C0141080250320F1802104A08"}, 1},
		{"Part 2", args{useInputData()}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := parseInput(tt.args.in)
			p, _ := readPackets(in)
			got := calculate(*p)
			assert.Equal(t, tt.want, got)
		})
	}
}

// C200B40A82 finds the sum of 1 and 2, resulting in the value 3.
// 04005AC33890 finds the product of 6 and 9, resulting in the value 54.
// 880086C3E88112 finds the minimum of 7, 8, and 9, resulting in the value 7.
// CE00C43D881120 finds the maximum of 7, 8, and 9, resulting in the value 9.
// D8005AC2A8F0 produces 1, because 5 is less than 15.
// F600BC2D8F produces 0, because 5 is not greater than 15.
// 9C005AC2F8F0 produces 0, because 5 is not equal to 15.
// 9C0141080250320F1802104A08 produces 1, because 1 + 3 = 2 * 2.
