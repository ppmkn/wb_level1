package main

import "fmt"

func main() {
	// Создаем мапу (множество) строк
	set := make(map[string]struct{})

	// Добавляем строки в множество
	set["cat"] = struct{}{}
	set["cat"] = struct{}{} // Дубликаты будут игнорироваться
	set["dog"] = struct{}{}
	set["cat"] = struct{}{}
	set["tree"] = struct{}{}

	// Проверяем принадлежность строки к множеству
	if _, exists := set["cat"]; exists {
		fmt.Println("Строка 'cat' принадлежит множеству.")
	}

	// Выводим все элементы множества
	fmt.Println("Элементы множества:")
	for key := range set {
		fmt.Println(key)
	}
}
