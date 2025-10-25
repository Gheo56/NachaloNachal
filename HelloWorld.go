package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("//// Калькулятор жирности тела ")
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
	var a string
	fmt.Scan(&a)
	return
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваша жирность массы тела: %.0f", IMT)
	fmt.Print(result)
}

func calculateIMT(userKg float64, userHeight float64) float64 {
	const IMTPower = 2
	const Meter = 100
	IMT := userKg / math.Pow(userHeight/Meter, IMTPower)
	return IMT
}
func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост (в сантиметрах): ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес (в килограммах): ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
