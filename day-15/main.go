package main

import (
	"fmt"
)

func main() {
	sensors := ParseInput()

	// for y := sensors.YMin(); y <= sensors.YMax(); y++ {
	// 	for x := sensors.XMin(); x <= sensors.XMax(); x++ {
	// 		// fmt.Println(point)
	// 		if !sensors.CanContainBeaconAt(Vec2{x, y}) {
	// 			count += 1
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	x := 0
	y := 0
	// bar := progressbar.Default(4000000)
	sensor := sensors.AnySensorNear(Vec2{x, y})
	for sensor != nil {
		// fmt.Println(x, y)
		// fmt.Println(sensor, sensor.Radius())
		// fmt.Printf("radius:%d, dy:%d, dx:%d\n", sensor.Radius(), y-sensor.Position.Y, x-sensor.Position.X)

		// skip to the end of the sensor x coordinate
		x += sensor.Radius() - absInt(y-sensor.Position.Y) - (x - sensor.Position.X) + 1

		if x > 4000000 {
			x = 0
			y += 1
			// bar.Add(1)
		}
		if y > 4000000 {
			panic("not found")
		}
		sensor = sensors.AnySensorNear(Vec2{x, y})
	}
	fmt.Println(x, y)
	fmt.Println(x*4000000 + y)
}
