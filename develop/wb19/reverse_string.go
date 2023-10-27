package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Функция для переворачивания строки с учетом символов Unicode
func reverseString(input string) string {
	// Преобразуем строку в слайс рун
	runes := []rune(input)
	length := len(runes)

	// Переворачиваем слайс рун
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	return string(runes) // Преобразуем слайс рун обратно в строку
}

func main() {
	fmt.Print("Введите любую строку: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Пишем любое слово/строку
	input = strings.TrimSpace(input)    // Убираем перенос строки, чтобы корректно отобразить
	reversed := reverseString(input)    // Вызываем нашу функцию
	fmt.Println("Перевернутая строка:", reversed)
}
