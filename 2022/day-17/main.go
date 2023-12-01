package main

import "fmt"

type RockCycle struct {
	RockStart RockType
	JetStart  int
	Size      uint
	Height    uint
}

func (chamber *VerticalChamber) IsAtCycleStart(cycle RockCycle) bool {
	if chamber.FallingRock == nil {
		return chamber.currentJet == cycle.JetStart && cycle.RockStart == Square // Square is the last rock in the list
	}
	return chamber.currentJet == cycle.JetStart && chamber.FallingRock.Type.Next() == cycle.RockStart
}
func (chamber *VerticalChamber) Cycle(cycle RockCycle) {
	if !chamber.IsAtCycleStart(cycle) {
		panic("tried to cycle while not at the cycle start")
	}
	for i := uint(0); i < cycle.Size; i += 1 {
		chamber.RockFall()
	}
}

func isCycle(rockStart RockType, jetStart int) (RockCycle, bool) {
	chamber := NewChamber()
	chamber.FallingRock = &FallingRock{Rock{Type: rockStart}, Vec2{0, 0}}
	chamber.currentJet = jetStart

	maxCycleTest := uint(len(chamber.jetPattern) * 100)

	chamber.RockFall()
	count := uint(1)
	for count < maxCycleTest {
		chamber.RockFall()
		count += 1

		if chamber.FallingRock.Type == rockStart && chamber.currentJet == jetStart {
			height := chamber.highestChunk

			// make sure that the chunks fits well with itself
			doubleCycleChamber := NewChamber()
			doubleCycleChamber.FallingRock = &FallingRock{Rock{Type: rockStart}, Vec2{0, 0}}
			doubleCycleChamber.currentJet = jetStart
			for i := uint(0); i < count*2; i++ {
				doubleCycleChamber.RockFall()
			}
			doubleCycleHeight := doubleCycleChamber.highestChunk
			if doubleCycleHeight != 2*height {
				continue
			}

			return RockCycle{RockStart: rockStart, JetStart: jetStart, Height: height, Size: count}, true
		}
	}
	return RockCycle{}, false
}

func main() {
	chamber := NewChamber()
	count := uint64(1_000_000_000_000)
	addedViaCycles := uint64(0)
	foundCycle := false
	for i := uint64(0); i < count; i += 1 {
		chamber.RockFall()
		if foundCycle {
			continue
		}
		if cycle, ok := isCycle(chamber.FallingRock.Type.Next(), chamber.currentJet); ok {
			// ignore cycle if it leads too far
			if uint64(cycle.Height)+i >= count {
				continue
			}
			foundCycle = true
			fmt.Printf("found cycle %+v\n", cycle)
			// add one whole cycle
			chamber.Cycle(cycle)
			i += uint64(cycle.Size)
			for i < count-uint64(cycle.Height) {
				i += uint64(cycle.Size)
				addedViaCycles += uint64(cycle.Height)
			}
		}
	}
	fmt.Println(addedViaCycles, chamber.highestChunk, addedViaCycles+uint64(chamber.highestChunk))
}
