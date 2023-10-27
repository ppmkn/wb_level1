package main

import (
	"fmt"
)

func main() {
	// Создаем каналы для передачи данных между стадиями конвейера
	inputChannel := make(chan int)  // Канал для ввода чисел
	outputChannel := make(chan int) // Канал для вывода результатов

	// 1 стадия: пишем числа в первый канал
	go func() {
		defer close(inputChannel)       // Закрываем канал после завершения записи данных
		numbers := []int{1, 2, 3, 4, 5} // Массив чисел, которые будем обрабатывать

		for _, x := range numbers {
			inputChannel <- x // Записываем число в канал
		}
	}()

	// 2 стадия: умножаем числа на 2 и пишем во 2-й канал
	go func() {
		defer close(outputChannel) // Закрываем канал после завершения обработки данных

		for x := range inputChannel {
			result := x * 2         // Умножаем число на 2
			outputChannel <- result // Записываем результат во второй канал
		}
	}()

	// 3 стадия: выводим результат
	for result := range outputChannel {
		fmt.Println(result)
	}
}
