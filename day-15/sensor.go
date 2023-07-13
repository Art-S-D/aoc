package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Sensor struct {
	Position      Vec2
	ClosestBeacon Vec2
}
type Sensors []Sensor

func (s *Sensor) Radius() int {
	return s.Position.Dist(s.ClosestBeacon)
}

func (s *Sensor) PreventsBeaconAt(pos Vec2) bool {
	if pos == s.ClosestBeacon {
		return true
	} else {
		return s.Position.Dist(pos) <= s.Radius()
	}
}

func (s Sensors) AnySensorNear(pos Vec2) *Sensor {
	for i := range s {
		if s[i].PreventsBeaconAt(pos) {
			return &s[i]
		}
	}
	return nil
}
func (s Sensors) CanContainBeaconAt(pos Vec2) bool {
	for i := range s {
		if s[i].PreventsBeaconAt(pos) {
			return false
		}
	}
	return true
}

func (s Sensors) XMin() int {
	min := s[0].Position.X - s[0].Radius()
	for i := range s {
		sensorMin := s[i].Position.X - s[i].Radius()
		if sensorMin < min {
			min = sensorMin
		}
	}
	return min
}
func (s Sensors) XMax() int {
	max := s[0].Position.X + s[0].Radius()
	for i := range s {
		sensorMax := s[i].Position.X + s[i].Radius()
		if sensorMax > max {
			max = sensorMax
		}
	}
	return max
}
func (s Sensors) YMin() int {
	min := s[0].Position.Y - s[0].Radius()
	for i := range s {
		sensorMin := s[i].Position.Y - s[i].Radius()
		if sensorMin < min {
			min = sensorMin
		}
	}
	return min
}
func (s Sensors) YMax() int {
	max := s[0].Position.Y + s[0].Radius()
	for i := range s {
		sensorMax := s[i].Position.Y + s[i].Radius()
		if sensorMax > max {
			max = sensorMax
		}
	}
	return max
}

func (s *Sensor) ListPositionsThatCannotContainABeacon(c chan<- Vec2) {
	dist := s.Radius()
	for x := s.Position.X - dist; x <= s.Position.X+dist; x++ {
		for y := s.Position.Y - dist; y <= s.Position.Y+dist; y++ {
			vec := Vec2{x, y}
			if s.Position.Dist(vec) <= dist {
				c <- vec
			}
		}
	}
}
func (s Sensors) ListPositionsThatCannotContainABeacon(c chan<- Vec2) {
	for i := range s {
		s[i].ListPositionsThatCannotContainABeacon(c)
	}
	close(c)
}

var lineRegex *regexp.Regexp = regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)

func ParseInput() Sensors {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	res := []Sensor{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if !lineRegex.MatchString(line) {
			panic(fmt.Sprintf("line <%s> does not match the regex", line))
		}
		submatches := lineRegex.FindStringSubmatch(line)
		var sensor Sensor
		fmt.Sscanf(submatches[1], "%d", &sensor.Position.X)
		fmt.Sscanf(submatches[2], "%d", &sensor.Position.Y)
		fmt.Sscanf(submatches[3], "%d", &sensor.ClosestBeacon.X)
		fmt.Sscanf(submatches[4], "%d", &sensor.ClosestBeacon.Y)
		res = append(res, sensor)
	}
	return res
}
