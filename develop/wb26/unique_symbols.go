package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func areAllCharactersUnique(s string) bool {
	s = strings.ToLower(s)         // Приводим строку к нижнему регистру
	charSet := make(map[rune]bool) // мапа для отслеживания уникальных символов в строке (символы строки : bool)

	for _, char := range s {
		// Проверяем, если символ уже встречался
		if charSet[char] {
			return false // Если символ уже встречался, строка не имеет уникальных символов
		}
		// Помечаем символ как встреченный
		charSet[char] = true
	}

	return true // Если ни один символ не повторяется, строка имеет уникальные символы
}

func main() {
	fmt.Print("Введите любую строку: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')     // Пишем любую строку
	input = strings.TrimSpace(input)        // Убираем перенос строки, чтобы корректно отображалось
	result := areAllCharactersUnique(input) // Кидаем в переменную result нашу функцию с нашей строкой
	fmt.Printf("%s — %t\n", input, result)  // Вывод результата
}
