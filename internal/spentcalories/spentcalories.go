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
	// TODO: реализовать функцию
	part := strings.Split(data, ",")
	if len(part) != 3 {
		return 0, "", 0, fmt.Errorf("Ошибка ввода формата")
	}
	activity := part[1]
	number, err := strconv.Atoi(part[0])
	if err != nil {
		return 0, "", 0, err
	}
	if number <= 0 {
		return 0, "", 0, fmt.Errorf("Error")
	}
	durat, err := time.ParseDuration(part[2])
	if err != nil {
		return 0, "", 0, err
	}
	if durat <= 0 {
		return 0, "", 0, fmt.Errorf("Error")
	}
	return number, activity, durat, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLenth := height * stepLengthCoefficient
	lenth := (float64(steps) * stepLenth) / float64(mInKm)
	return lenth
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration.Hours() <= 0 {
		return 0
	}
	return distance(steps, height) / duration.Hours()

}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	number, activity, durat, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	duratInH := durat.Hours()
	speed := meanSpeed(number, height, durat)
	var calorii float64
	switch activity {
	case "Ходьба":
		calorii, err = WalkingSpentCalories(number, weight, height, durat)
		if err != nil {
			return "", err
		}
	case "Бег":
		calorii, err = RunningSpentCalories(number, weight, height, durat)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activity, duratInH, distance(number, height), speed, calorii)
	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("Ошибка шагов")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("Ошибка веса")
	}
	if height <= 0 {
		return 0, fmt.Errorf("Ошибка роста")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("Error")
	}

	return (weight * meanSpeed(steps, height, duration) * duration.Hours()), nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("Ошибка шагов")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("Ошибка веса")
	}
	if height <= 0 {
		return 0, fmt.Errorf("Ошибка веса")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("Ошибка времени")
	}

	durationInmin := duration.Minutes()

	return ((weight * meanSpeed(steps, height, duration)) * float64(durationInmin) / float64(minInH)) * walkingCaloriesCoefficient, nil
}
