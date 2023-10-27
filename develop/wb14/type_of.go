package main

import (
	"fmt"
)

// Функция getType принимает переменную типа interface{} и возвращает строку, представляющую тип значения.
func getType(value interface{}) string {
	// Используем оператор switch для определения типа значения в переменной value
	switch value.(type) {
	case int:
		return "int" // Если тип значения int, возвращаем строку "int"
	case string:
		return "string" // Если тип значения string, возвращаем строку "string"
	case bool:
		return "bool" // Если тип значения bool, возвращаем строку "bool"
	case chan interface{}:
		return "channel" // Если тип значения chan interface{}, возвращаем строку "channel"
	default:
		return "unknown" // В остальных случаях возвращаем "unknown"
	}
}

func main() {
	var value interface{} // Создаем переменную типа interface{}.

	// Присваиваем разные значения переменной интерфейса
	value = 42 // int
	fmt.Printf("Тип: %s\n", getType(value))

	value = "Hello, WB!" // string
	fmt.Printf("Тип: %s\n", getType(value))

	value = true // bool
	fmt.Printf("Тип: %s\n", getType(value))

	ch := make(chan interface{})
	value = ch // channel
	fmt.Printf("Тип: %s\n", getType(value))

	value = struct{ Name string }{"Пётр"} // unknown, так как это пользовательский тип данных
	fmt.Printf("Тип: %s\n", getType(value))
}
