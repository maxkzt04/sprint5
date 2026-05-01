package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля

	// Количество шагов, тип тренировки, продолжительность, личные данные
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию

	// Разбиваем строку запятой
	parts := strings.Split(datastring, ",")

	// Долхжно быть только 3, если болше то ошибка, если меьше то ошибка:
	// Шаги, Тип тренировки, Продолжительность
	//
	if len(parts) != 3 {
		return errors.New("неверный формат строки: ожидается 3 поля через запятую")
	}

	// Получаем количество шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("не удалось преобразовать шаги в число: %w", err)
	}
	// Шаги должны быть больше нуля
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}
	t.Steps = steps

	// Сохраняем тип тренировки в структуре TrainingType
	t.TrainingType = parts[1]

	// Получаем продолжительность
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("не удалось распознать продолжительность: %w", err)
	}
	// Продолжительность должна быть больше нуля
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}
	t.Duration = duration

	return nil
}

// ActionInfo формирует и возвращает строку с информацией о тренировке.
func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	// Подсчитываем дистанцию
	dist := spentenergy.Distance(t.Steps, t.Height)
	// Тоже самое только узнаем среднюю скорость
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64

	// Узнаем сколько калорий было сожжено в зависимости от типа тренировки при беге или ходьбе, если дргое то ошибка
	switch t.TrainingType {
	case "Бег":
		cal, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		calories = cal
	case "Ходьба":
		cal, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		calories = cal
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	// Вормируем результат об тренировки
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		dist,
		speed,
		calories,
	)

	return result, nil
}
