package main

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
// 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
// градусов. Последовательность в подмножноствах не важна.

import (
	"fmt"
	"math"
)

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := groupTemp(temp)

	printGroups(groups)
}

func groupTemp(temp []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temperature := range temp {
		groupIndex := calculateGroupIndex(temperature)
		groups[groupIndex] = append(groups[groupIndex], temperature)
	}

	return groups
}

func calculateGroupIndex(temperature float64) int {
	return int(math.Floor(temperature/10)) * 10
}

func printGroups(groups map[int][]float64) {
	for groupIndex, groupTemp := range groups {
		fmt.Printf("Group %d: %v\n", groupIndex, groupTemp)
	}
}
