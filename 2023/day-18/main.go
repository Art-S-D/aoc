package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type DigInstruction struct {
	Dir   Direction
	Count int
	Color string
}

func (i DigInstruction) String() string {
	return fmt.Sprintf("%v %d (%s)", i.Dir, i.Count, i.Color)
}

var instructionRegexp *regexp.Regexp = regexp.MustCompile(`^(U|D|L|R) (\d+) \(#(.*)\)$`)

func InstructionFromHex(str string) DigInstruction {
	dir := str[len(str)-1]
	count := str[:len(str)-1]
	countInt, _ := strconv.ParseInt(count, 16, 32)
	return DigInstruction{
		Count: int(countInt),
		Dir:   Directions[dir-'0'],
	}
}

type DigPlan []DigInstruction

func ParsePlan() DigPlan {
	var plan DigPlan
	for _, line := range parseInput() {
		matches := instructionRegexp.FindStringSubmatch(line)
		// dir := DirectionFromByte(matches[1][0])
		// count, _ := strconv.Atoi(matches[2])
		// color := matches[3]
		// plan = append(plan, DigInstruction{Dir: dir, Count: count, Color: color})
		plan = append(plan, InstructionFromHex(matches[3]))
	}
	return plan
}

type Lagoon map[Vec]bool

func main() {
	plan := ParsePlan()

	area := 0
	pos := Vec{}
	for _, instruction := range plan {
		nextPos := pos.Add(instruction.Dir.ToVec().Scale(instruction.Count))

		area += (pos.Y+nextPos.Y)*(pos.X-nextPos.X) + instruction.Count

		pos = nextPos
	}
	fmt.Println(area/2 + 1)
}
