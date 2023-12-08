package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

type Node struct {
	Label string
	Left  string
	Right string
}
type Map map[string]Node

func (m Map) FollowInstructions(path string) (steps uint) {
	currentNode := "AAA"
	for currentNode != "ZZZ" {
		direction := path[int(steps)%len(path)]
		if direction == 'L' {
			currentNode = m[currentNode].Left
		} else {
			currentNode = m[currentNode].Right
		}
		steps += 1
	}
	return steps
}

func (m Map) StartingNodes() []string {
	var res []string
	for k := range m {
		if k[2] == 'A' {
			res = append(res, k)
		}
	}
	return res
}

func AllEndInZ(nodes []string) bool {
	for _, n := range nodes {
		if n[2] != 'Z' {
			return false
		}
	}
	return true
}

func (m Map) FollowMultipleInstructions(path string) (steps uint) {
	currentNodes := m.StartingNodes()
	for !AllEndInZ(currentNodes) {
		direction := path[int(steps)%len(path)]
		for i, node := range currentNodes {
			if direction == 'L' {
				currentNodes[i] = m[node].Left
			} else {
				currentNodes[i] = m[node].Right
			}
		}
		steps += 1
		if steps%1_000_000 == 0 {
			fmt.Println(steps)
		}
	}
	return steps
}

type Cycle struct {
	Length      uint
	StartsAfter uint
}
type nodeEncountered struct {
	label string
	at    uint
}

func (m Map) SearchForCycle(startingAt, path string) Cycle {
	nodesEncountered := make(map[nodeEncountered]uint)
	steps := uint(0)
	currentNode := startingAt
	currentEncounter := nodeEncountered{label: currentNode, at: 0}
	for nodesEncountered[currentEncounter] == 0 {
		nodesEncountered[currentEncounter] = steps
		steps += 1

		direction := path[int(steps)%len(path)]
		if direction == 'L' {
			currentNode = m[currentNode].Left
		} else {
			currentNode = m[currentNode].Right
		}

		currentEncounter = nodeEncountered{label: currentNode, at: steps % uint(len(path))}
	}
	cycleStartsAt := nodesEncountered[currentEncounter]
	return Cycle{
		Length:      steps - cycleStartsAt,
		StartsAfter: cycleStartsAt,
	}
}

//go:embed input.txt
var input string

var nodeRegexp regexp.Regexp = *regexp.MustCompile(`([A-Z0-9]+) = \(([A-Z0-9]+), ([A-Z0-9]+)\)`)

func ParseInput() (steps string, m Map) {
	res := make(Map)
	lines := strings.Split(input, "\n")
	steps = lines[0]
	for _, line := range lines[2:] {
		if len(line) == 0 {
			continue
		}
		matches := nodeRegexp.FindStringSubmatch(line)
		node := Node{Label: matches[1], Left: matches[2], Right: matches[3]}
		res[node.Label] = node
	}

	return steps, res
}

func main() {
	path, inputMap := ParseInput()
	// steps := inputMap.FollowMultipleInstructions(path)
	// fmt.Println(steps)

	cyclesLength := []int{}
	for _, node := range inputMap.StartingNodes() {
		// yeah they all start at 1
		cyclesLength = append(cyclesLength, int(inputMap.SearchForCycle(node, path).Length))
	}

	fmt.Println(LCM(cyclesLength...))
}
