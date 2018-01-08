package main

import "fmt"
// 1 уровень, 17 лекция
func main() {
// Подсчитать, сколько бутылочек воды уходит на душ в завичимости от времени
	fmt.Println("Укажите количество проведенных минут Вами в душе")
	var time float64
	fmt.Scanf("%f", &time)
	water := time * 12
	fmt.Println("Принимая душ Вы израсходовали", water, "бутылок воды")
}
