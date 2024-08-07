package main

import (
	"fmt"
	"sort"
)

type Car struct {
	position int
	speed    int
}

func main() {
	target := 12
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}
	fmt.Println(carFleet(target, position, speed)) // Output: 3
}
func carFleet(target int, position []int, speed []int) int {
	n := len(position)

	if n == 0 {
		return 0
	}
	cars := make([]Car, n)
	for i := range position {
		cars[i] = Car{position[i], speed[i]}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})
	times := make([]float64, n)
	for i := range cars {
		times[i] = float64(target-cars[i].position) / float64(cars[i].speed)
	}

	fleets := 0
	currentFleetTime := 0.0
	fmt.Println(times)
	for i := 0; i < n; i++ {
		if times[i] > currentFleetTime {
			fleets++
			currentFleetTime = times[i]
		}
	}
	return fleets
}
