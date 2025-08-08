package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// Разделяем строку
	sepData := strings.Split(data, ",")

	// Проверяем правильность разделения
	if len(sepData) != 2 {
		return 0, 0, fmt.Errorf("Не удалось обработать данные")
	}

	// Преобразуем количество шагов в число
	stepsCount, errSteps := strconv.Atoi(sepData[0])
	if errSteps != nil {
		return 0, 0, fmt.Errorf("Ошибка преобразования числа шагов: %s. Текст ошибки: %s", sepData[0], errSteps)
	}

	//Проверка шагов на ноль
	if stepsCount == 0 {
		return 0, 0, fmt.Errorf("Количество шагов - 0")
	}

	//Парсинг времени
	duration, errTime := time.ParseDuration(sepData[1])
	if errTime != nil {
		return 0, 0, fmt.Errorf("Ошибка преобразования времени ходьбы: %s. Текст ошибки: %s", sepData[1], errTime)
	}

	return stepsCount, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	//Получаем данные из parsePackage
	stepsCount, duration, errParse := parsePackage(data)
	if errParse != nil {
		fmt.Println(errParse)
		return ""
	}

	//Проверка шагов на ноль
	if stepsCount == 0 {
		return ""
	}

	//Вычисление дистанции и калорий
	distance := float64(stepsCount) * stepLength / float64(mInKm)
	cal, _ := spentcalories.WalkingSpentCalories(stepsCount, weight, height, duration)

	//Формируем строку
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", stepsCount, distance, cal)
}
