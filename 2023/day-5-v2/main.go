package main

import "fmt"

type Range struct {
	Start  uint
	Length uint
}

func (r Range) End() uint {
	return r.Start + r.Length
}
func (r Range) Overlaps(other Range) bool {
	return r.End() > other.Start && other.End() > r.Start
}

type MappingRange struct {
	Range
	Dest uint
}

// assumes r contains src
func (r *MappingRange) Map(src uint) (dest uint) {
	return (src - r.Start) + r.Dest
}

// takes a mapping and a range
// they should be overlapping
// it will return: 1: the part of r that overlaps the mapping, already mapped
// and 2: the rest of r that has not been mapped
// nonOverlapped always has at least one element
func (mapping *MappingRange) MapRange(r Range) (overlappedRange Range, nonOverlapped []Range) {
	// Check if there is no overlap
	if mapping.End() <= r.Start || r.End() <= mapping.Start {
		panic("no overlap")
	}

	// Calculate the start and end of the overlapping range
	overlapStart := max(mapping.Start, r.Start)
	overlapEnd := min(mapping.End(), r.End())

	// Create the left non-overlapping range if exists
	if r.Start < overlapStart {
		leftRange := Range{Start: r.Start, Length: overlapStart - r.Start}
		nonOverlapped = append(nonOverlapped, leftRange)
	}

	// Create the overlapping range
	overlappedRange = Range{Start: mapping.Map(overlapStart), Length: overlapEnd - overlapStart}

	// Create the right non-overlapping range if exists
	if overlapEnd < r.End() {
		rightRange := Range{Start: overlapEnd, Length: r.End() - overlapEnd}
		nonOverlapped = append(nonOverlapped, rightRange)
	}

	return overlappedRange, nonOverlapped
}

type CategoryMap struct {
	Src    string
	Dest   string
	Ranges []MappingRange
}

func (m *CategoryMap) MapRange(r Range) []Range {
	res := []Range{}
	currentNonOverlappedRanges := []Range{r}
	for _, currentRange := range m.Ranges {
		for i := 0; i < len(currentNonOverlappedRanges); i++ {
			// fmt.Println(len(res), len(currentNonOverlappedRanges))
			if currentRange.Overlaps(currentNonOverlappedRanges[i]) {
				// fmt.Printf("%+v overlaps with %+v\n", currentRange, currentNonOverlappedRanges[i])
				overlapped, notOverlapped := currentRange.MapRange(currentNonOverlappedRanges[i])
				// fmt.Printf("overlapped:%+v | notOverlapped %+v\n", overlapped, notOverlapped)
				res = append(res, overlapped)
				currentNonOverlappedRanges[i] = currentNonOverlappedRanges[len(currentNonOverlappedRanges)-1]
				currentNonOverlappedRanges = currentNonOverlappedRanges[:len(currentNonOverlappedRanges)-1]
				currentNonOverlappedRanges = append(currentNonOverlappedRanges, notOverlapped...)
				i--
			}
		}
	}
	return append(res, currentNonOverlappedRanges...)
}

type Almanach struct {
	Seeds []Range
	Maps  []CategoryMap
}

func (a *Almanach) FinalRange(r Range) []Range {
	currentRanges := []Range{r}
	nextRanges := []Range{}
	for _, m := range a.Maps {
		for _, r := range currentRanges {
			nextRanges = append(nextRanges, m.MapRange(r)...)
		}
		currentRanges = nextRanges
		nextRanges = []Range{}
	}
	return currentRanges
}

func (a *Almanach) MinimumLocation() uint {
	ranges := []Range{}
	for _, r := range a.Seeds {
		ranges = append(ranges, a.FinalRange(r)...)
	}
	minimum := ranges[0].Start
	for _, r := range ranges {
		minimum = min(r.Start, minimum)
	}
	return minimum
}

func main() {
	input := parseInput()
	almanach := AlmanacFromString(input)
	fmt.Println(almanach.MinimumLocation())
}
