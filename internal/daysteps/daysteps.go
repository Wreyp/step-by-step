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
	// TODO: реализовать функцию

	part := strings.Split(data, ",")

	if len(part) != 2 {
		return 0, 0, fmt.Errorf("ошибка ввода формата")
	}

	number, err := strconv.Atoi(part[0])
	if err != nil {
		return 0, 0, err
	}

	durat, err := time.ParseDuration(part[1])
	if err != nil {
		return 0, 0, err
	}
	return number, durat, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}
	if steps <= 0 {
		return ""
	}

	distance := (float64(steps) * stepLength)
	distanceInKm := distance / mInKm

	calorii, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}
	result := fmt.Sprintf("Количество шагов: %d\nДистанция составила: %.2f км.\nВы сожгли: %.2f ккал.", steps, distanceInKm, calorii)
	return result
}
