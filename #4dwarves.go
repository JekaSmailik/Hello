package main

import "fmt"
// 1 уровень, 4 лекция
func main() {
	// Цыкл вложеный в другой цикл, повторяющий действие заданное количество раз.
	for dwarves := 0; dwarves < 3; dwarves++ {
		for SnowWhite := 0; SnowWhite < 1; SnowWhite++ {
			fmt.Println("Нет! мне не нужна помощь!!!")
		}
		fmt.Println("Я помогу тебе, Белоснежка!")
	}
}
