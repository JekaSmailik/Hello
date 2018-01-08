package main

import "fmt"
// 1 уровень, 4 лекция
func main() {
// Если условие выполняется, сделать действие, если нет - сделать другое
	fmt.Println("Сколько Вам лет? ")
	var age float64
	fmt.Scanf("%f", &age)
	for {
		if age >= 18 {
			fmt.Println("Доступ разрешен ")
		} else {
			fmt.Println("Вам отказано в доступе ")
		}
		break
	}
}
