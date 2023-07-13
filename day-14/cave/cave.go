package cave

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type element uint8

const (
	Air element = iota
	Rock
	Sand
)

func (el element) String() string {
	switch el {
	case Air:
		return "."
	case Rock:
		return "#"
	case Sand:
		return "O"
	default:
		panic(fmt.Sprintf("unknown element %d\n", el))
	}
}

type Vec2 struct {
	X int
	Y int
}

//	func (v Vec2) Add(other Vec2) Vec2 {
//		return Vec2{v.X + other.X, v.Y + other.Y}
//	}
//
//	func (v Vec2) Subtract(other Vec2) Vec2 {
//		return Vec2{v.X - other.X, v.Y - other.Y}
//	}
//
//	func (v Vec2) Unit() Vec2 {
//		res := Vec2{}
//		if v.X < 0 {
//			res.X = -1
//		} else if v.X > 0 {
//			res.X = 1
//		}
//		if v.Y < 0 {
//			res.Y = -1
//		} else if v.Y > 0 {
//			res.Y = 1
//		}
//		return res
//	}
func (v Vec2) String() string {
	return fmt.Sprintf("%d,%d", v.X, v.Y)
}
func Vec2FromString(s string) (res Vec2) {
	splitted := strings.Split(s, ",")
	fmt.Sscanf(splitted[0], "%d", &res.X)
	fmt.Sscanf(splitted[1], "%d", &res.Y)
	return res
}

// note that the origin of the cave starts on te top left
type Cave struct {
	content    map[Vec2]element
	sandSpawn  Vec2
	lowestRock int
}

func (cave *Cave) At(x, y int) element {
	if y >= 2+cave.lowestRock {
		return Rock
	}
	elem, ok := cave.content[Vec2{x, y}]
	if !ok {
		return Air
	} else {
		return elem
	}
}
func (cave *Cave) AtV(v Vec2) element {
	return cave.At(v.X, v.Y)
}

func (cave *Cave) SpawnSand() (endsAt Vec2, voids bool) {
	if cave.AtV(cave.sandSpawn) != Air {
		return cave.sandSpawn, false
	}

	grain := cave.sandSpawn
	cave.content[grain] = Sand
	// for grain.Y <= cave.lowestRock {
	for {
		if cave.At(grain.X, grain.Y+1) == Air {
			// directly under te grain
			delete(cave.content, grain)
			grain.Y += 1
			cave.content[grain] = Sand
		} else if cave.At(grain.X-1, grain.Y+1) == Air {
			// diagonally, to the left
			delete(cave.content, grain)
			grain.Y += 1
			grain.X -= 1
			cave.content[grain] = Sand
		} else if cave.At(grain.X+1, grain.Y+1) == Air {
			// diagonally, to the right
			delete(cave.content, grain)
			grain.Y += 1
			grain.X += 1
			cave.content[grain] = Sand
		} else {
			return grain, false
		}
	}
	// remove the grain to prevent filling the void
	delete(cave.content, grain)
	return grain, true
}

func (cave *Cave) Debug(xStart, xEnd, yStart, yEnd int) {
	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			fmt.Printf("%v", cave.At(x, y))
		}
		fmt.Println()
	}
}

func Parse() *Cave {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	cave := Cave{
		content:    make(map[Vec2]element),
		lowestRock: 0,
		sandSpawn:  Vec2{500, 0},
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, " -> ")
		previousPoint := Vec2FromString(points[0])
		if previousPoint.Y > cave.lowestRock {
			cave.lowestRock = previousPoint.Y
		}
		for _, p := range points[1:] {
			point := Vec2FromString(p)

			if previousPoint.X != point.X && previousPoint.Y != point.Y {
				panic(fmt.Sprintf("rock line is not straight %v -> %v\n", previousPoint, point))
			}

			i := previousPoint
			for i != point {
				cave.content[i] = Rock

				if previousPoint.X < point.X {
					i.X += 1
				} else if previousPoint.X > point.X {
					i.X -= 1
				}
				if previousPoint.Y < point.Y {
					i.Y += 1
				} else if previousPoint.Y > point.Y {
					i.Y -= 1
				}
			}

			if point.Y > cave.lowestRock {
				cave.lowestRock = point.Y
			}

			cave.content[point] = Rock
			previousPoint = point
		}
	}
	return &cave
}
