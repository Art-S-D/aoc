package main

import (
	"bufio"
	"fmt"
	"os"
)

type V2 struct {
	X int
	Y int
}

func (self *V2) Add(other V2) V2 {
	return V2{self.X + other.X, self.Y + other.Y}
}

type InputDirective struct {
	Dir      rune
	Distance uint
}

func (i *InputDirective) ToVec() V2 {
	var res V2
	switch i.Dir {
	case 'U':
		res.Y += int(i.Distance)
	case 'D':
		res.Y -= int(i.Distance)
	case 'L':
		res.X -= int(i.Distance)
	case 'R':
		res.X += int(i.Distance)
	}
	return res
}

func (i *InputDirective) ToUnaryVec() V2 {
	var res V2
	switch i.Dir {
	case 'U':
		res.Y += 1
	case 'D':
		res.Y -= 1
	case 'L':
		res.X -= 1
	case 'R':
		res.X += 1
	}
	return res
}

func (i *InputDirective) String() string {
	return fmt.Sprintf("%c %d", i.Dir, i.Distance)
}

func parseInput() []InputDirective {
	var input []InputDirective

	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	buf := bufio.NewScanner(inputFile)
	for buf.Scan() {
		line := buf.Text()
		var directive InputDirective
		fmt.Sscanf(line, "%c %d", &directive.Dir, &directive.Distance)
		input = append(input, directive)
	}

	return input
}

func Sign(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func moveOnce(head, tail V2) V2 {
	if abs(head.X-tail.X) < abs(head.Y-tail.Y) {
		tail.X = head.X
		tail.Y = head.Y - Sign(head.Y-tail.Y)
	} else if abs(head.X-tail.X) > abs(head.Y-tail.Y) {
		tail.X = head.X - Sign(head.X-tail.X)
		tail.Y = head.Y
	} else {
		tail.X = head.X - Sign(head.X-tail.X)
		tail.Y = head.Y - Sign(head.Y-tail.Y)
	}

	return tail
}
func stepRope(rope []V2, instruction InputDirective, visited map[V2]bool) {
	for step := 0; step < int(instruction.Distance); step++ {
		rope[0] = rope[0].Add(instruction.ToUnaryVec())

		for i := 1; i < len(rope); i++ {
			newPosition := moveOnce(rope[i-1], rope[i])
			rope[i] = newPosition
		}

		tail := rope[len(rope)-1]
		visited[tail] = true

		// for i := 20; i >= -20; i-- {
		// 	for j := -20; j <= 20; j++ {
		// 		found := false
		// 		for index, pos := range rope {
		// 			if pos.X == j && pos.Y == i {
		// 				fmt.Printf("%d", index)
		// 				found = true
		// 				break
		// 			}
		// 		}
		// 		if !found {
		// 			fmt.Print(".")
		// 		}
		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println()
		// fmt.Scanf("%s\n", "")

	}

}

func main() {
	input := parseInput()

	rope := make([]V2, 10)
	visited := make(map[V2]bool)

	for _, instruction := range input {
		stepRope(rope, instruction, visited)
	}

	fmt.Println(len(visited))
}
