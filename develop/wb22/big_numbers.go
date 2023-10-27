package main

import (
	"fmt"
)

func main() {
	a := 1 << 21 // 2 в степени 21
	b := 1 << 20 // 2 в степени 20

	// Умножение
	resultMultiply := a * b
	fmt.Printf("Умножение: %d * %d = %d\n", a, b, resultMultiply)

	// Деление
	if b != 0 { // Проверяем, чтобы b не было равно 0, чтобы избежать деления на ноль)
		resultDivide := a / b
		fmt.Printf("Деление: %d / %d = %d\n", a, b, resultDivide)
	} else {
		fmt.Println("Деление на ноль!")
	}

	// Сложение
	resultAdd := a + b
	fmt.Printf("Сложение: %d + %d = %d\n", a, b, resultAdd)

	// Вычитание
	resultSubtract := a - b
	fmt.Printf("Вычитание: %d - %d = %d\n", a, b, resultSubtract)
}
