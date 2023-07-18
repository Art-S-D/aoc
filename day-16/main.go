package main

import "fmt"

func main() {
	initialValve, allValves := ParseInput()
	// fmt.Println(initialValve)
	// matrix := FloydWarshall(allValves)
	// matrix.Debug()

	// takes +- 600s
	run := NewRun(initialValve, allValves)
	pressure := run.MostPressureReleasable()
	fmt.Println(pressure)
}
