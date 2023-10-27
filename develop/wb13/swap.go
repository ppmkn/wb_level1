package main

import "fmt"

func main() {
	a := 5
	b := 10

	fmt.Printf("До обмена: a = %d, b = %d\n", a, b) // Результат ДО

	a, b = b, a // Просто присваиваем переменным значения друг друга

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b) // Результат ПОСЛЕ
}
