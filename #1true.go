package main

import "fmt"
// 1 уровень, 2 лекция
func main() {
// Вывод четных и нечетных значений
	for number := 0; number <= 10; number++ {
		if number%2 == 0 {
			fmt.Println(number, "четное")
		} else {
			fmt.Println(number, "не четное")
		}
	}
}
