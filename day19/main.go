package main

import (
	"bufio"
	"fmt"
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

func (c coord) equals(o coord) bool {
	return c.x == o.x && c.y == o.y && c.z == o.z
}

func (c coord) rotateX() coord {
	return coord{x: c.x, y: -c.z, z: c.y}
}

func (c coord) rotateY() coord {
	return coord{x: c.z, y: c.y, z: -c.x}
}

func (c coord) rotateZ() coord {
	return coord{x: -c.y, y: c.x, z: c.z}
}

func (c coord) in(xs []coord) bool {
	for _, v := range xs {
		if c == v {
			return true
		}
	}

	return false
}

type seaScanner struct {
	location coord
	beacons  []coord
}

func (s seaScanner) toString() string {
	result := fmt.Sprintf("Location: %+v\nBeacons:\n", s.location)
	for _, v := range s.beacons {
		result += fmt.Sprintf("\t%+v\n", v)
	}
	return result + "\n"
}

func (s *seaScanner) permutations() []seaScanner {
	result := []seaScanner{}

	v := make([]coord, len(s.beacons))
	q := make([]coord, len(s.beacons))
	copy(v, s.beacons)
	copy(q, s.beacons)
	for i := 0; i < 4; i++ {
		newBeacons := make([]coord, len(v))
		copy(newBeacons, v)
		result = append(result, seaScanner{s.location, newBeacons})
		for k, x := range v {
			v[k] = x.rotateX()
		}
	}

	for i, x := range v {
		v[i] = x.rotateY()
	}

	for i := 0; i < 4; i++ {
		newBeacons := make([]coord, len(v))
		copy(newBeacons, v)
		result = append(result, seaScanner{s.location, newBeacons})
		for j, x := range v {
			v[j] = x.rotateZ()
		}
	}

	for i, x := range v {
		v[i] = x.rotateY()
	}

	for i := 0; i < 4; i++ {
		newBeacons := make([]coord, len(v))
		copy(newBeacons, v)
		result = append(result, seaScanner{s.location, newBeacons})
		for k, x := range v {
			v[k] = x.rotateX()
		}
	}

	for i, x := range v {
		v[i] = x.rotateY()
	}

	for i := 0; i < 4; i++ {
		newBeacons := make([]coord, len(v))
		copy(newBeacons, v)
		result = append(result, seaScanner{s.location, newBeacons})
		for j, x := range v {
			v[j] = x.rotateZ()
		}
	}

	for i, x := range v {
		v[i] = x.rotateY().rotateZ()
	}

	for i := 0; i < 4; i++ {
		newBeacons := make([]coord, len(v))
		copy(newBeacons, v)
		result = append(result, seaScanner{s.location, newBeacons})
		for j, x := range v {
			v[j] = x.rotateY()
		}
	}

	for i, x := range v {
		v[i] = x.rotateZ().rotateZ()
	}

	for i := 0; i < 4; i++ {
		newBeacons := make([]coord, len(v))
		copy(newBeacons, v)
		result = append(result, seaScanner{s.location, newBeacons})
		for j, x := range v {
			v[j] = x.rotateY()
		}
	}

	return result
}

func (s *seaScanner) compare(o *seaScanner) bool {
	tests := o.permutations()

	for _, test := range tests {
		for _, b := range test.beacons {
			for _, a := range s.beacons {
				test.location = s.location.add(a).subtract(b)
				result := 0
				for _, c := range test.beacons {
					if test.location.add(c).subtract(s.location).in(s.beacons) {
						result++
					}
				}
				if result >= 12 {
					*o = test
					return true
				}
			}
		}
	}

	return false
}

func main() {
	unlocated := getInput()
	located := []*seaScanner{unlocated[0]}
	unlocated = unlocated[1:]
	result := located[0].beacons
	for len(unlocated) > 0 {
		for _, v := range located {
			for _, x := range unlocated {
				if v.compare(x) {
					located = appendUniqueSeaScanner(located, x)
					unlocated = remove(unlocated, x)
					for _, y := range x.beacons {
						result = appendIfUnique(result, x.location.add(y))
					}
				}
			}
		}
	}

	max := 0
	for _, a := range located {
		for _, b := range located {
			manhattonVector := a.location.subtract(b.location)
			manhattanDistance := manhattonVector.x + manhattonVector.y + manhattonVector.z
			if manhattanDistance > max {
				max = manhattanDistance
			}
		}
	}

	log.Printf("Result: %d", len(result))
	log.Printf("Greatest Manhattan distance: %d units", max)
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
			z, err := strconv.Atoi(rawCoord[2])
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

func appendUniqueSeaScanner(xs []*seaScanner, in *seaScanner) []*seaScanner {
	for _, v := range xs {
		if v == in {
			return xs
		}
	}

	return append(xs, in)
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
