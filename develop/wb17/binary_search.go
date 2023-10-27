package main

import (
	"fmt"
	"sort"
)

func main() {
	// Создаем срез
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Значение, которое мы ищем
	target := 4

	// Используем sort.Search для выполнения бинарного поиска
	index := sort.Search(len(data), func(i int) bool {
		return data[i] >= target
	})

	// Проверяем, найдено ли значение
	if index < len(data) && data[index] == target {
		fmt.Printf("Значение %d найдено на позиции %d\n", target, index)
	} else {
		fmt.Printf("Значение %d не найдено в массиве\n", target)
	}
}
