package main

import "fmt"

func main() {
	// Исходный слайс
	originalSlice := []int{1, 2, 3, 4, 5}

	// Индекс элемента, который вы хотите удалить (например, третий элемент с индексом 2)
	indexToRemove := 2

	// Создаем новый слайс, в который будем копировать элементы
	var modifiedSlice []int

	// Копируем элементы до индекса indexToRemove
	modifiedSlice = append(modifiedSlice, originalSlice[:indexToRemove]...)

	// Копируем элементы после индекса indexToRemove
	modifiedSlice = append(modifiedSlice, originalSlice[indexToRemove+1:]...)

	// Выводим результат
	fmt.Println("Исходный слайс:", originalSlice)
	fmt.Println("Удаляем число 3") // Так понятнее :)
	fmt.Println("Слайс после удаления элемента:", modifiedSlice)
}
