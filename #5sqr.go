package main

import "fmt"
// 1 уровень, 15 лекция
func number() {
// Функция "посчитать квадрат целого числа"
	fmt.Println("Введите число")
	var number float64
	fmt.Scanf("%f", &number)
	fmt.Println("Квадрат введенного числа равен ", number*number)
}
func main() {
// Вывод Hello World
	fmt.Println("С Новый")
	welcome()
	number()
}
func welcome() {
// Вызов функции в главной пограмме, выдать "Hello everyone"
	fmt.Println("2018 годом!")
}
// Пример однострочного комментария
/** а это -
			многострочного **/