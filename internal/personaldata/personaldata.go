package personaldata

import "fmt"

// Personal хранит личные данные пользователя.
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print выводит данные пользователя на экран.
func (p Personal) Print() {
	// TODO: реализовать функцию

	// Делаем данные в читабельном виде
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f\n", p.Weight)
	fmt.Printf("Рост: %.2f\n", p.Height)
}
