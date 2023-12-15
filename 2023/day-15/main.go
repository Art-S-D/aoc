package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func hash(s string) uint8 {
	value := 0
	for _, c := range s {
		value += int(c)
		value *= 17
		value = value % 256
	}
	return uint8(value)
}

func sumHashes(input []string) int {
	res := 0
	for _, s := range input {
		res += int(hash(s))
	}
	return res
}

type Lens struct {
	Label string
	Focal int
}
type Box []Lens

func (l Lens) Hash() uint8 {
	return hash(l.Label)
}
func (l Lens) String() string {
	return fmt.Sprintf("[%s %d]", l.Label, l.Focal)
}

func (b *Box) Add(lens Lens) {
	findLensByLabel := func(l Lens) bool { return l.Label == lens.Label }
	if slices.ContainsFunc(*b, findLensByLabel) {
		index := slices.IndexFunc(*b, findLensByLabel)
		(*b)[index] = lens
	} else {
		*b = append(*b, lens)
	}
}
func (b *Box) Remove(label string) {
	newBox := slices.DeleteFunc(*b, func(lens Lens) bool { return lens.Label == label })
	*b = newBox
}

func ParseInstruction(instruction string) (label string, focal int, operation rune) {
	if strings.Contains(instruction, "=") {
		split := strings.Split(instruction, "=")
		label := split[0]
		focal, _ := strconv.Atoi(split[1])
		return label, focal, '='
	} else if strings.Contains(instruction, "-") {
		return instruction[:len(instruction)-1], 0, '-'
	}
	panic(fmt.Sprintf("failed to parse instruction %s\n", instruction))
}

func main() {
	inputList := strings.Split(input, ",")
	boxes := make([]Box, 256)
	for _, instruction := range inputList {
		label, focal, operation := ParseInstruction(instruction)
		boxNumber := hash(label)
		if operation == '-' {
			boxes[boxNumber].Remove(label)
		} else {
			boxes[boxNumber].Add(Lens{Label: label, Focal: focal})
		}
	}
	for i, b := range boxes {
		if len(b) == 0 {
			continue
		}
		fmt.Printf("Box %d: %+v\n", i, b)
	}

	res := 0
	for boxNumber, box := range boxes {
		for slotNumber, lens := range box {
			res += (1 + boxNumber) * (slotNumber + 1) * lens.Focal
		}
	}
	fmt.Println(res)
}
