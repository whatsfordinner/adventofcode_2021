package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
	z int
}

func (c coord) add(o coord) coord {
	return coord{x: c.x + o.x, y: c.y + o.y, z: c.z + o.z}
}

func (c coord) subtract(o coord) coord {
	return coord{x: c.x - o.x, y: c.y - o.y, z: c.z - o.z}
}

func (c coord) rotateX() coord {
	return coord{x: c.x, y: -c.z, z: c.y}
}

func (c coord) rotateY() coord {
	return coord{x: c.z, y: c.y, z: -c.x}
}

func (c coord) rotateZ() coord {
	return coord{x: c.y, y: -c.x, z: c.z}
}

func (c coord) permutations() []coord {
	result := []coord{c}
	v := c
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				v = v.rotateX()
				result = append(result, v)
			}
			v = v.rotateY()
		}
		v = v.rotateZ()
	}
	return result
}

func (c coord) equals(o coord) bool {
	return c.x == o.x && c.y == o.y && c.z == o.z
}

type seaScanner struct {
	location coord
	beacons  []coord
}

func (s *seaScanner) permutations() []seaScanner {
	result := []seaScanner{}
	permutations := [][]coord{}
	for _, v := range s.beacons {
		beaconPermutations := v.permutations()
		permutations = append(permutations, beaconPermutations)
	}

	for i := 0; i < len(permutations[0]); i++ {
		rotatedScanner := seaScanner{location: s.location}
		for _, v := range permutations {
			rotatedScanner.beacons = append(rotatedScanner.beacons, v[i])
		}
		result = append(result, rotatedScanner)
	}
	return result
}

func (s *seaScanner) compare(o *seaScanner) []coord {
	result := []coord{}
	tests := o.permutations()

	for _, test := range tests {
		for _, a := range s.beacons {
			aAbsolute := s.location.add(a)
			test.location = aAbsolute.subtract(test.beacons[0])
			result = []coord{}
			for _, b := range test.beacons {
				for _, c := range s.beacons {
					if s.location.add(c).equals(test.location.add(b)) {
						result = append(result, c)
					}
				}
			}
			if len(result) >= 12 {
				*o = test
				return result
			}
		}
	}

	return result
}

func main() {
	unlocated := getInput()
	located := []*seaScanner{unlocated[0]}
	unlocated = unlocated[1:]
	result := []coord{}
	for len(unlocated) > 0 {
		for _, v := range located {
			log.Printf("%+v", v)
			for _, x := range unlocated {
				common := v.compare(x)
				if len(common) >= 12 {
					located = append(located, x)
					unlocated = remove(unlocated, x)
					for _, v := range common {
						result = appendIfUnique(result, v)
					}
				}
			}
		}
	}

	log.Printf("%d", len(result))
}

func getInput() []*seaScanner {
	result := []*seaScanner{}
	scanner := bufio.NewScanner(os.Stdin)
	seaScannerRegex := regexp.MustCompile("--- scanner [0-9]+ ---")
	inSeaScanner := false
	newSeaScanner := new(seaScanner)
	for scanner.Scan() {
		if seaScannerRegex.Match([]byte(scanner.Text())) {
			inSeaScanner = true
			continue
		}

		if scanner.Text() == "" {
			inSeaScanner = false
			result = append(result, newSeaScanner)
			newSeaScanner = new(seaScanner)
			continue
		}

		if inSeaScanner {
			rawCoord := strings.Split(scanner.Text(), ",")
			x, err := strconv.Atoi(rawCoord[0])
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(rawCoord[1])
			if err != nil {
				log.Fatal(err)
			}
			z, err := strconv.Atoi(rawCoord[1])
			newSeaScanner.beacons = append(newSeaScanner.beacons, coord{x: x, y: y, z: z})
		}
	}
	result = append(result, newSeaScanner)
	return result
}

func appendIfUnique(cs []coord, c coord) []coord {
	for _, v := range cs {
		if v.equals(c) {
			return cs
		}
	}
	return append(cs, c)
}

func remove(xs []*seaScanner, x *seaScanner) []*seaScanner {
	result := []*seaScanner{}
	for _, v := range xs {
		if v != x {
			result = append(result, v)
		}
	}

	return result
}
