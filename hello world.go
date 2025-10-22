package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64
	fmt.Println("___Калькулятор Жирности XD___")
	fmt.Print("Введите свой Рост в сантиметрах:")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес:")
	fmt.Scan(&userKg)
	var IMT = userKg / math.Pow(userHeight/100, IMTPower)
	fmt.Printf("___Твоя жирность___ %.0f", IMT)
}
