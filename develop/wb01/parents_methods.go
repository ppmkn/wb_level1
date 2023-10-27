package main

import "fmt"

// Структура Human
type Human struct {
	Name    string //Тут просто разные поля внутри структуры
	Age     int    //
	Address string //
}

// Метод SayHello для структуры Human
func (h *Human) SayHello() {
	fmt.Printf("Привет, меня зовут %s!\n", h.Name)
}

// Структура Action с встраиванием Human
type Action struct {
	Human        // Структура Human
	Test  string // просто поле в структуре Action, для наглядности
}

// Ещё метод DoSomething для структуры Action
func (a *Action) DoSomething() {
	fmt.Printf("%s выполняет какое-то действие.\n", a.Name) // Обращение к полю структуры Human, которая встраивается в структуру Action
	fmt.Printf("%s\n", a.Test)                              // Обращение к полю структуры Action
}

func main() {
	/*
	   Создаём объект типа Action
	   Указываем тип Action, чтобы иметь доступ к полям и методам как Action, так и Human
	*/
	person := Action{
		Test: "Строчка из структуры Action",
		Human: Human{ // Встраиваем структуру Human в структуру Action и создаем новый объект структуры Human с указанными значениями полей
			Name:    "Иван",
			Age:     30,
			Address: "Москва",
		},
	}
	// Тут уже тестируем результат:
	// Вызываем методы как для Human, так и для Action
	person.SayHello()
	person.DoSomething()
}
