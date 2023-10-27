package main

import (
	"fmt"
	"math"
)

func main() {
	// создаём слайс, где хранятся исходные значения
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// создаем мап для группировки значений
	temperatureGroups := make(map[int][]float64)

	// Проходим по исходным значениям
	for _, temp := range temperatures {
		// Округляем значение до ближайшего десятка
		roundedTemp := int(math.Round(temp/10) * 10)

		// Добавляем значение в соответствующую группу
		temperatureGroups[roundedTemp] = append(temperatureGroups[roundedTemp], temp)
	}

	// Выводим результат
	for group, temps := range temperatureGroups {
		fmt.Printf("%d: %v\n", group, temps)
	}
}
