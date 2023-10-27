package main

import (
	"fmt"
	"time"
)

func main() {
	// Канал для обмена данными между горутинами
	dataChannel := make(chan int)

	// Запускаем горутину для чтения данных из канала
	go func() {
		for {
			// Читаем данные из канала
			data, ok := <-dataChannel
			if !ok {
				// Если канал закрыт, завершаем горутину чтения
				fmt.Println("Горутина чтения завершена.")
				return
			}
			// Вывод прочитанных данных
			fmt.Printf("Прочитано: %d\n", data)
		}
	}()

	// Запускаем горутину для отправки данных в канал
	go func() {
		for i := 1; i <= 10; i++ {
			// Отправляем данные в канал
			dataChannel <- i
			// Задержка в 1 секунду
			time.Sleep(1 * time.Second)
		}
		// Закрываем канал после отправки всех данных
		close(dataChannel)
	}()

	// Ждем N секунд (например, 5 секунд)
	N := 5
	time.Sleep(time.Duration(N) * time.Second)

	// Завершаем программу
	fmt.Println("Программа завершена.")
}
