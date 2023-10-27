package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем map для хранения данных и мьютекс
	data := make(map[string]int) // Пустой map для хранения данных
	mutex := sync.Mutex{}        // Мьютекс для синхронизации доступа к map

	// Создаем канал для сигналов о завершении горутин
	done := make(chan bool)

	// Количество горутин, которые будут записывать данные
	numWorkers := 5

	// Функция, которая выполняет конкурентную запись в map
	worker := func(id int) {
		// Записываем данные в map
		for i := 0; i < 10; i++ {
			key := fmt.Sprintf("Key%d", i) // Генерируем уникальный ключ
			value := id*10 + i             // Генерируем значение

			// Выполняем захват мьютекс перед записью, чтобы предотвратить одновременный доступ нескольких горутин
			mutex.Lock()
			data[key] = value // Записываем данные в map
			// Освобождаем мьютекс после записи
			mutex.Unlock()
		}

		// Сигнализируем о завершении работы горутины
		done <- true
	}

	// Запускаем горутины для записи данных
	for i := 0; i < numWorkers; i++ {
		go worker(i)
	}

	// Ожидаем завершения всех горутин
	for i := 0; i < numWorkers; i++ {
		<-done
	}

	// Выводим результат
	fmt.Println("Данные в map:")
	for key, value := range data {
		fmt.Printf("%s: %d\n", key, value)
	}
}
