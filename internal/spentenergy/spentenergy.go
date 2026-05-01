package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	// проверяем входные данные
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше нуля")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть больше нуля")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть больше нуля")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше нуля")
	}

	speed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := weight * speed * durationInMinutes / minInH
	// при ходьбе применяем корректирующий коэффициент
	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	// Проверяем входные данные на валидность
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше нуля")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть больше нуля")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть больше нуля")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше нуля")
	}

	// Для бега используем ту же формулу, что и для ходьбы, но без корректировки (Не Забыть. "нужно проверить")
	speed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := weight * speed * durationInMinutes / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию

	// Как и в других функциях, проверяем входные данные
	if duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	// Переводим продолжительнось тренировки в часы
	durationInHours := duration.Hours()
	return dist / durationInHours
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию

	// Шаг = рост * коэффициент для расчета длины шага
	stepLength := height * stepLengthCoefficient

	// Расстояние в метрах = шаги * длина шага
	distanceInMeters := float64(steps) * stepLength

	// Переводим из метров в километры
	return distanceInMeters / mInKm
}
