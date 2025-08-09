package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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
		return 0, "0", 0, fmt.Errorf("Не удалось обработать данные")
	}

	// Преобразуем количество шагов в число
	stepsCount, errSteps := strconv.Atoi(sepData[0])
	if errSteps != nil {
		return 0, "0", 0, fmt.Errorf("Ошибка преобразования числа шагов: %s. Текст ошибки: %s", sepData[0], errSteps)
	}

	//Парсинг времени
	time, errTime := time.ParseDuration(sepData[2])
	if errTime != nil {
		return 0, "0", 0, fmt.Errorf("Ошибка преобразования времени ходьбы: %s. Текст ошибки: %s", sepData[2], errTime)
	}

	if stepsCount <= 0 || time <= 0 {
		log.Println("Не положительные данные")
		return 0, "0", 0, nil
	}

	return stepsCount, sepData[1], time, nil
}

func distance(steps int, height float64) float64 {
	// Длина шага
	stepLength := height * stepLengthCoefficient

	// Дистанция
	dist := stepLength * float64(steps) / float64(mInKm)

	return dist
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Проверка длительности на ноль
	if duration <= 0 {
		return 0
	}

	// Скорость
	speed := distance(steps, height) / duration.Hours()

	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// Получение данных из строки
	steps, training, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	//Формирование информации по типу тренировки
	switch training {
	case "Ходьба":
		dist := distance(steps, height)
		speed := meanSpeed(steps, height, duration)
		cal, _ := WalkingSpentCalories(steps, weight, height, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", training, duration.Hours(), dist, speed, cal), nil
	case "Бег":
		dist := distance(steps, height)
		speed := meanSpeed(steps, height, duration)
		cal, _ := RunningSpentCalories(steps, weight, height, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", training, duration.Hours(), dist, speed, cal), nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка входных параметров
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Параметры должны быть больше 0")
	}

	//Расчет калорий
	cal := (meanSpeed(steps, height, duration) * weight * duration.Minutes()) / float64(minInH)

	return cal, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка входных параметров
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Параметры должны быть больше 0")
	}

	//Расчет калорий
	cal := (meanSpeed(steps, height, duration) * weight * duration.Minutes()) / float64(minInH) * walkingCaloriesCoefficient

	return cal, nil
}
