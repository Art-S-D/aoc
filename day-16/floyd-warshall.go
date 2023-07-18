package main

import "fmt"

type AdjacentMatrix map[string]map[string]uint

func initAdjacentMatrix(valves []*Valve) AdjacentMatrix {
	res := make(map[string]map[string]uint)
	for i := range valves {
		label := valves[i].Label

		// initialize to 0
		res[label] = make(map[string]uint)
		for j := range valves {
			res[label][valves[j].Label] = 0
		}

		// set direct neighbor to 1
		for _, neighbor := range valves[i].LeadsToLabel {
			res[label][neighbor] = 1
		}
	}
	return res
}

func FloydWarshall(valves []*Valve) AdjacentMatrix {
	res := initAdjacentMatrix(valves)

	for k := 0; k < len(valves); k++ {
		labelK := valves[k].Label
		for i := 0; i < len(valves); i++ {
			labelI := valves[i].Label
			for j := 0; j < len(valves); j++ {
				labelJ := valves[j].Label

				if labelI == labelJ {
					continue
				}

				// if path to i->k or k->j does not exists, do nothing
				if res[labelI][labelK] == 0 || res[labelK][labelJ] == 0 {
					continue
				}

				// if i->k+k->j is shorter than i->j, replace i->j
				oldPath := res[labelI][labelJ]
				nexPath := res[labelI][labelK] + res[labelK][labelJ]
				if oldPath == 0 || oldPath > nexPath {
					res[labelI][labelJ] = nexPath
					res[labelJ][labelI] = nexPath
				}
			}
		}
	}

	return res
}

func (matrix AdjacentMatrix) Debug() {
	labels := []string{}
	for label := range matrix {
		labels = append(labels, label)
	}

	fmt.Print("  ")
	for _, label := range labels {
		fmt.Printf("  %s", label)
	}
	fmt.Println()

	for _, label1 := range labels {
		fmt.Printf("%s  ", label1)
		for _, label2 := range labels {
			fmt.Printf("%2d  ", matrix[label1][label2])
		}
		fmt.Println()
	}
}
