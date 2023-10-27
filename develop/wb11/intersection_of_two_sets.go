package main

import "fmt"

// Функция intersection будет принимать 2 неупорядоченных множества (мапы) set1 и set2 и возвращать их пересечение
func intersection(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool) // Создаем пустой мап для результата

	// Проходим по элементам первого множества (set1)
	for element := range set1 {
		// Если элемент также присутствует во втором множестве (set2), добавляем его в результат
		if set2[element] {
			result[element] = true
		}
	}

	return result // Возвращаем мап, представляющий пересечение множеств
}

func main() {
	// Создаем два неупорядоченных множества в виде мап (set1 и set2)
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	// Заполняем первое множество элементами
	set1[1] = true
	set1[2] = true
	set1[3] = true

	// Заполняем второе множество элементами
	set2[2] = true
	set2[3] = true
	set2[4] = true

	// Вызываем функцию intersection для нахождения пересечения множеств
	result := intersection(set1, set2)

	// Выводим результат
	fmt.Println("Пересечение множеств:")
	for element := range result {
		fmt.Println(element)
	}
}
