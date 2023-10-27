package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Количество воркеров (горутин)
	numWorkers := 5

	/*
		Канал для записи данных.
		Канал позволяет безопасно обмениваться данными между горутинами
	*/
	dataChannel := make(chan string)

	// Канал для сигналов операционной системы ( в моем случаи для Ctrl+C )
	osSignals := make(chan os.Signal, 1)
	// Регистрируем обработчики сигналов
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Запускаем воркеры (горутины) для обработки данных
	for i := 0; i < numWorkers; i++ {
		go worker(i, dataChannel)
	}

	// Главная горутина записывает данные в канал
	for {
		// Слушаем сигналы операционной системы, чтобы завершить программу по Ctrl+C
		select {
		case <-osSignals:
			// Программа завершает работу по сигналу Ctrl+C
			fmt.Println("Программа завершена по сигналу Ctrl+C")
			return
		default:
			// Записываем произвольные данные в канал для воркеров
			dataChannel <- "Произвольные данные"
		}
	}
}

// Воркер (горутина) для обработки данных
func worker(id int, dataChannel <-chan string) {
	for data := range dataChannel {
		// Воркер получает и выводит данные
		fmt.Printf("Воркер %d получил данные: %s\n", id, data)
	}
}
