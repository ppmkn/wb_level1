package main

import (
	"fmt"
	"sync"
)

// Создаем структуру Counter
type Counter struct {
	mu    sync.Mutex // Мьютекс для синхронизации доступа к переменной value
	value int        // Переменная, которую мы инкрементируем
}

// Метод Increment для структуры Counter
func (c *Counter) Increment() {
	c.mu.Lock()   // Блокировка мьютекса для предотвращения одновременного доступа к переменной value
	c.value++     // Инкрементация переменной value
	c.mu.Unlock() // Разблокировка мьютекса
}

// Метод GetValue для структуры Counter
func (c *Counter) GetValue() int {
	c.mu.Lock()         // Блокировка мьютекса для предотвращения одновременного доступа к переменной value
	defer c.mu.Unlock() // Разблокировка мьютекса после выполнения функции
	return c.value      // Возвращаем текущее значение переменной value
}

func main() {
	counter := &Counter{} // Создаем экземпляр структуры Counter
	numIncrements := 1000 // Количество инкрементаций, которые нужно выполнить

	var wg sync.WaitGroup // Структура для управления горутинами и ожидания их завершения

	for i := 0; i < numIncrements; i++ {
		wg.Add(1) // Добавляем одну горутину в WaitGroup
		go func() {
			counter.Increment() // Вызываем метод Increment структуры Counter для инкрементации значения
			wg.Done()           // Горутина завершила выполнение, уменьшаем счетчик в WaitGroup
		}()
	}

	wg.Wait() // Ожидание завершения всех инкрементаций

	finalValue := counter.GetValue()                           // Получаем итоговое значение счетчика
	fmt.Printf("Итоговое значение счетчика: %d\n", finalValue) // Выводим итоговое значение счетчика в консоль
}
