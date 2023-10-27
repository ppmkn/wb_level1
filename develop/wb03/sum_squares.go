package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int)
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			square := num * num
			results <- square
		}(num)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	/*
	   До этого момента код был взят из предыдущего задания,
	   так как для этого задания тоже необходимо для начала
	   вычислить квадрат чисел ( из массива ),
	   а теперь просто вычислим их сумму:
	*/
	sum := 0
	for square := range results {
		sum += square
	}

	fmt.Printf("Сумма квадратов чисел: %d\n", sum)
}
