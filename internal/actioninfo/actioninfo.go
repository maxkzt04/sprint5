package actioninfo

import "log"

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию

	// Проходим по всем строкам данных
	for _, data := range dataset {
		// Разобрать строку
		err := dp.Parse(data)
		if err != nil {
			// Если не получилось - логируем ошибку и идём дальше
			log.Println("ошибка парсинга:", err)
			continue
		}

		// Создаем строку с информацией о тренировке
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("ошибка формирования информации:", err)
			continue
		}

		// Инвормация в лог
		log.Print(info)
	}
}
