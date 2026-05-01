package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps хранит данные о дневных прогулках.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию

	// Разбиваем строку запятой
	parts := strings.Split(datastring, ",")

	// Должно быть только 2, если болше то ошибка, если меьше то ошибка:
	// Шаги, Продолжительность
	if len(parts) != 2 {
		return errors.New("неверный формат строки: ожидается 2 поля через запятую")
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
	ds.Steps = steps

	// Получаем продолжительность
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("не удалось распознать продолжительность: %w", err)
	}
	// Продолжительность должна быть больше нуля
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}
	ds.Duration = duration

	return nil

	//*******************Проверить
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	// Подсчитываем дистанцию
	dist := spentenergy.Distance(ds.Steps, ds.Height)

	// Подсчитываем калории для ходьбы
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	// Создаем строку с результатами
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		dist,
		calories,
	)

	return result, nil
}
