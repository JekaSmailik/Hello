package main

import "fmt"
// 1 уровень, 4 лекция
func main() {
// Оператор выбора.
	fmt.Println("Укажите одно из следующих чисел: 1; 30; 55 или 100")
	var number float64
	fmt.Scanf("%f", &number)
	switch number {
	case 100:
		fmt.Println("Поздравляю, Вы сделали правильный выбор, указав число 100")
	case 55:
		fmt.Println("Поздравляю, Вы сделали верный выбор, указав число 55")
	case 30:
		fmt.Println("Поздравляю, Вы сделали обоснованный выбор, указав число 30")
	case 1:
		fmt.Println("Поздравляю, Вы сделали очевидный выбор, указав число 1")
	default:
		fmt.Println("Введите число из предложенных вариантов")
	return
	}
}

