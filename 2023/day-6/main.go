package main

import (
	"fmt"
	"strings"
)

type Race struct {
	Time         uint
	BestDistance uint
}

func (r Race) CanWinIfPressedFor(time uint) bool {
	travelTime := r.Time - time
	travelSpeed := time
	travelDistance := travelTime * travelSpeed
	return travelDistance > r.BestDistance
}

func (r Race) NumberOfWaysToWin() uint {
	res := uint(0)
	for i := uint(0); i <= r.Time; i++ {
		if r.CanWinIfPressedFor(i) {
			res += 1
		}
	}
	return res
}

func RacesFromInput(input []string) []Race {
	times := strings.Fields(input[0][len("Distance:"):])
	distances := strings.Fields(input[1][len("Distance:"):])

	if len(times) != len(distances) {
		panic("times and distances don't have the same size")
	}

	var res []Race
	for i := range times {
		time := uint(MustAtoi(times[i]))
		distance := uint(MustAtoi(distances[i]))
		res = append(res, Race{Time: time, BestDistance: uint(distance)})
	}
	return res
}
func RaceFromInputV2(input []string) Race {
	time := strings.Replace(input[0][len("Distance:"):], " ", "", -1)
	distance := strings.Replace(input[1][len("Distance:"):], " ", "", -1)

	t := uint(MustAtoi(time))
	d := uint(MustAtoi(distance))
	return Race{Time: t, BestDistance: uint(d)}
}

func main() {
	input := parseInput()

	// res := 1
	// for _, race := range RacesFromInput(input) {
	// 	res *= int(race.NumberOfWaysToWin())
	// }
	// fmt.Println(res)

	race := RaceFromInputV2(input)
	fmt.Println(race.NumberOfWaysToWin())
}
