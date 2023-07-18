package main

import (
	"fmt"
	"strings"
)

type Run struct {
	// MUST BE AN ARRAY and not a slice
	At          [2]*Valve
	TimeLeft    [2]int
	AlreadyOpen string
	allValves   []*Valve
	adjacents   AdjacentMatrix
}

func (run *Run) HasAlreadyOpened(valve *Valve) bool {
	return strings.Contains(run.AlreadyOpen, valve.Label)
}
func (run *Run) Copy() *Run {
	return &Run{
		// copied since it is an array
		At:          run.At,
		TimeLeft:    run.TimeLeft,
		AlreadyOpen: run.AlreadyOpen,
		allValves:   run.allValves,
		adjacents:   run.adjacents,
	}
}

func (run *Run) MostPressureReleasable() uint {
	nextStep := run.Copy()

	maxReleasable := uint(0)

	for player := 0; player < 2; player++ {
		currentValve := run.At[player]

		timeLeft := run.TimeLeft[player]
		if timeLeft <= 0 {
			continue
		}

		for i := range run.allValves {
			targetValve := run.allValves[i]
			if targetValve.Ppm == 0 || run.HasAlreadyOpened(targetValve) {
				continue
			}
			distanceToValve := int(run.adjacents[currentValve.Label][targetValve.Label])
			if distanceToValve >= timeLeft {
				continue
			}

			// dont't use -=, we need to reset nextStep.TimeLeft from the previous loop body
			// -1 to account for te time to open the valve
			nextStep.TimeLeft[player] = timeLeft - distanceToValve
			nextStep.At[player] = targetValve
			nextStep.AlreadyOpen = fmt.Sprintf("%s,%s", run.AlreadyOpen, targetValve.Label)

			nextStep.TimeLeft[player] -= 1
			totalPressureReleased := uint(targetValve.Ppm) * uint(nextStep.TimeLeft[player])

			pressure := nextStep.MostPressureReleasable()
			releasable := pressure + totalPressureReleased
			if releasable > maxReleasable {
				maxReleasable = releasable
				// log = fmt.Sprintf(
				// 	"player %d moves from %s to %s and opens it in %d minutes(%d left). Will release %d pressure\n%s",
				// 	player,
				// 	currentValve.Label,
				// 	targetValve.Label,
				// 	distanceToValve+1,
				// 	nextStep.TimeLeft[player],
				// 	maxReleasable, nextLog,
				// )
			}
		}
	}
	return maxReleasable
}

func NewRun(from *Valve, allValves []*Valve) *Run {
	return &Run{
		At:          [2]*Valve{from, from},
		TimeLeft:    [2]int{26, 26},
		AlreadyOpen: "",
		allValves:   allValves,
		adjacents:   FloydWarshall(allValves),
	}
}
