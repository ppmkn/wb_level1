package main

import (
	"fmt"
	"sort"
)

// Создаем пользовательский тип для сортировки (в данном случае, int)
type IntSlice []int

// Реализуем методы интерфейса sort.Interface для нашего пользовательского типа
func (s IntSlice) Len() int {
	return len(s)
}
func (s IntSlice) Swap(i, j int) {
	s[i],
	s[j] = s[j],
	s[i]
}
func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func main() {
	// Создаем массив для сортировки
	data := []int{6, 2, 8, 4, 1, 3, 9, 5, 7}

	// Преобразуем массив в пользовательский тип
	intSlice := IntSlice(data)

	// Используем sort.Sort для сортировки
	sort.Sort(intSlice)

	// Выводим отсортированный массив
	fmt.Println(intSlice)
}
