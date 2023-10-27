/*
Паттерн "Адаптер" (Adapter) используется для соединения двух несовместимых интерфейсов
Он позволяет объектам с несовместимыми интерфейсами работать вместе
*/
package main

import "fmt"

// Интерфейс для метрической системы
type MetricSystem interface {
	GetLengthInMeters() float64
}

// Интерфейс для американской системы
type AmericanSystem interface {
	GetLengthInFeet() float64
}

// Структура для измерения длины в метрах
type LengthInMeters struct {
	Value float64
}

// Метод GetLengthInMeters возвращает длину измерения в метрах
func (l *LengthInMeters) GetLengthInMeters() float64 {
	return l.Value
}

// Структура для измерения длины в футах
type LengthInFeet struct {
	Value float64
}

// Метод GetLengthInFeet возвращает длину измерения в метрах
func (l *LengthInFeet) GetLengthInFeet() float64 {
	return l.Value
}

// Адаптер для преобразования метрической системы в американскую ( футы )
type MetricToAmericanAdapter struct {
	Metric LengthInMeters
}

// Метод GetLengthInFeet преобразует метрическое измерение (метры) в футы
func (a *MetricToAmericanAdapter) GetLengthInFeet() float64 {
	// Преобразовываем метры в футы (1 метр = 3.2808 фута)
	return a.Metric.GetLengthInMeters() * 3.2808
}

func main() {
	// Измерение в метрической системе
	metricLength := LengthInMeters{Value: 5.0}

	// Измерение в футах (с использованием адаптера)
	adapter := MetricToAmericanAdapter{Metric: metricLength}

	fmt.Printf("Длина в метрах: %.2f\n", metricLength.GetLengthInMeters())
	fmt.Printf("Длина в футах: %.2f\n", adapter.GetLengthInFeet())
}
