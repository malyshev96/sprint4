package spentcalories

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// Разделяем строку
	sepData := strings.Split(data, ",")

	// Проверяем правильность разделения
	if len(sepData) != 3 {
		return 0, 0, fmt.Errorf("Не удалось обработать данные")
	}

	// Преобразуем количество шагов в число
	stepsCount, errSteps := strconv.Atoi(sepData[0])
	if errSteps != nil {
		return 0, "0", 0, fmt.Errorf("Ошибка преобразования числа шагов: %s. Текст ошибки: %s", sepData[0], errSteps)
	}

	//Парсинг времени
	time, errTime := time.ParseDuration(sepData[2])
	if errTime != nil {
		return 0, "0", 0, fmt.Errorf("Ошибка преобразования времени ходьбы: %s. Текст ошибки: %s", sepData[1], errTime)
	}

	return stepsCount, sepData[1], time, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}
