package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Функция reverseWords переворачивает порядок слов в строке
func reverseWords(input string) string {
	// Разбиваем строку на слова с помощью пробелов.
	words := strings.Fields(input)

	// Создаем пустой слайс для перевернутых слов
	reversedWords := make([]string, 0, len(words))

	// Переворачиваем порядок слов и добавляем их в слайс
	for i := len(words) - 1; i >= 0; i-- {
		reversedWords = append(reversedWords, words[i])
	}

	// Склеиваем перевернутые слова с пробелами и возвращаем результат
	return strings.Join(reversedWords, " ")
}

func main() {
	fmt.Print("Введите любую строку: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Пишем любую строку
	reversed := reverseWords(input)
	fmt.Println("Результат:", reversed) // Вывод результата
}
