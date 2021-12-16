package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type literalPacket struct {
	version    int
	packetType int
	value      int
}

func (p literalPacket) getVersion() int {
	return p.version
}

func (p literalPacket) getPacketType() int {
	return p.packetType
}

func (p literalPacket) sumVersion() int {
	return p.version
}

func (p literalPacket) evaluate() int {
	return p.value
}

type operatorPacket struct {
	version    int
	packetType int
	value      []packet
}

func (p operatorPacket) getVersion() int {
	return p.version
}

func (p operatorPacket) getPacketType() int {
	return p.packetType
}

func (p operatorPacket) sumVersion() int {
	var result int

	result += p.version
	for _, v := range p.value {
		result += v.sumVersion()
	}

	return result
}

func (p operatorPacket) evaluate() int {
	var result int

	switch p.packetType {
	case 0:
		for _, v := range p.value {
			result += v.evaluate()
		}
	case 1:
		result = 1
		for _, v := range p.value {
			result *= v.evaluate()
		}
	case 2:
		result = math.MaxInt
		for _, v := range p.value {
			val := v.evaluate()
			if val < result {
				result = val
			}
		}
	case 3:
		result = 0
		for _, v := range p.value {
			val := v.evaluate()
			if val > result {
				result = val
			}
		}
	case 5:
		if p.value[0].evaluate() > p.value[1].evaluate() {
			result = 1
		} else {
			result = 0
		}
	case 6:
		if p.value[0].evaluate() < p.value[1].evaluate() {
			result = 1
		} else {
			result = 0
		}
	case 7:
		if p.value[0].evaluate() == p.value[1].evaluate() {
			result = 1
		} else {
			result = 0
		}
	}

	return result
}

type packet interface {
	getVersion() int
	getPacketType() int
	sumVersion() int
	evaluate() int
}

func main() {
	input, _ := getInput()
	log.Printf("Sum of versions: %d", input.sumVersion())
	log.Printf("Evluation: %d", input.evaluate())
}

func getInput() (packet, string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input = hexToBin(input)
	return parsePacket(input)
}

func parsePacket(in string) (packet, string) {
	version := binToDec(in[:3])
	packetType := binToDec(in[3:6])

	if packetType == 4 {
		rawContent := in[6:]
		var content string
		i := 0
		for i < len(rawContent) {
			content += rawContent[i+1 : i+5]
			if rawContent[i] == '0' {
				i += 5
				break
			}
			i += 5
		}
		return literalPacket{
			version:    version,
			packetType: packetType,
			value:      binToDec(content),
		}, in[i+6:]
	}

	lengthType := in[6]
	var remainder string
	subPackets := []packet{}
	if lengthType == '0' {
		lengthOfSubpackets := binToDec(in[7:22])
		remainder = in[22:]
		targetRemainder := len(remainder) - lengthOfSubpackets
		for len(remainder) > targetRemainder {
			newPacket, out := parsePacket(remainder)
			subPackets = append(subPackets, newPacket)
			remainder = out
		}
	} else {
		countOfSubpackets := binToDec(in[7:18])
		remainder = in[18:]
		for i := 0; i < countOfSubpackets; i++ {
			newPacket, out := parsePacket(remainder)
			subPackets = append(subPackets, newPacket)
			remainder = out
		}
	}

	return operatorPacket{
		version:    version,
		packetType: packetType,
		value:      subPackets,
	}, remainder
}

func hexToBin(in string) string {
	var result string

	for _, v := range strings.Split(in, "") {
		switch v {
		case "0":
			result += "0000"
		case "1":
			result += "0001"
		case "2":
			result += "0010"
		case "3":
			result += "0011"
		case "4":
			result += "0100"
		case "5":
			result += "0101"
		case "6":
			result += "0110"
		case "7":
			result += "0111"
		case "8":
			result += "1000"
		case "9":
			result += "1001"
		case "A":
			result += "1010"
		case "B":
			result += "1011"
		case "C":
			result += "1100"
		case "D":
			result += "1101"
		case "E":
			result += "1110"
		case "F":
			result += "1111"
		}
	}

	return result
}

func binToDec(in string) int {
	var result int

	for i, v := range strings.Split(in, "") {
		bit, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		result += bit * int(math.Pow(2, float64(len(in)-i-1)))
	}

	return result
}
