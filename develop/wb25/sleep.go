package main

import (
	"fmt"
	"time"
)

// Моя функция для задержки выполнения программы на указанное время
func Sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second) // Ожидаем "seconds" секунд
}

func main() {
	fmt.Println("Программа начала выполнение.")

	// Используем функцию Sleep для приостановки выполнения программы на 2 секунды
	Sleep(2)

	fmt.Println("Программа продолжает выполнение после задержки.")
}
