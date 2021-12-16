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

type packet interface {
	getVersion() int
	getPacketType() int
}

func main() {
	log.Printf("%+v", getInput())
}

func getInput() packet {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input = hexToBin(input)
	log.Printf("%s", input)
	return parsePacket(input)
}

func parsePacket(in string) packet {
	version := binToDec(in[:3])
	packetType := binToDec(in[3:6])

	log.Printf("Version: %d, Type: %d", version, packetType)

	if packetType == 4 {
		rawContent := in[6:]
		var content string
		for i := 0; i < len(rawContent); i += 5 {
			content += rawContent[i+1 : i+5]
			if rawContent[i] == '0' {
				break
			}
		}
		return literalPacket{
			version:    version,
			packetType: packetType,
			value:      binToDec(content),
		}
	}

	lengthType := in[6]
	if lengthType == '0' {
		lengthOfSubpackets := binToDec(in[7:22])
		log.Printf("Length: %d", lengthOfSubpackets)
	} else {
		countOfSubpackets := binToDec(in[7:18])
		log.Printf("Count: %d", countOfSubpackets)
	}

	return operatorPacket{}
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
