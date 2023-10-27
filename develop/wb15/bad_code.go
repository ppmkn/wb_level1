package main

import (
	"fmt"
)

var justString string

// createHugeString создает большую строку заданного размера и возвращает ее
func createHugeString(size int) string {
	return "Очень большая строка размером " + fmt.Sprint(size) // Строка + её размер
}

// someFunc выполняет операции с большой строкой и сохраняет ее в глобальной переменной justString
func someFunc() {
	v := createHugeString(1 << 10)
	// Убедимся, что срез не выходит за пределы длины строки.
	if len(v) > 100 {
		justString = v[:100]
	} else { // Без этого код не работает :)
		justString = v
	}
}

func main() {
	someFunc()
	fmt.Println(justString) // Выводим результат в консоль
}
