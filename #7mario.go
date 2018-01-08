package main

import "fmt"
// 1 уровень, 18 лекция
func main() {
// Строительство полупирамидки с помощью хэша (#) высотой от 0 до 23
	fmt.Println("Укажите высоту полупирамидки")
	var height int
	fmt.Scanf("%d", &height)
// (height < 0 || height > 23) необходимо выполнить проверку на условие от 0 до 23
	for line := 0; line < height; line++ {
		for space := line + 1; space-height < 0; space++ {
			fmt.Printf(" ")
		}
		for hes := 0; hes < line+1; hes++ {
			fmt.Printf("#")
		}
		fmt.Println(" ")
	}
}
