package main

import (
	"fmt"
	"math"
)

func main() {
	PrintHeader("Калькулятор Жирности Массы Тела 💪")
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)

	clearScreen()

	outputResult(IMT)
	fmt.Println("\nЧто бы выйти введите любую клавишу и нажмите Enter")
	var a string
	fmt.Scan(&a)
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
	clearScreen()
	return userKg, userHeight
}
func PrintHeader(title string) {
	fmt.Println("\033[36m╔════════════════════════════════════════════╗")
	fmt.Printf("║ %-32s         ║\n", title)
	fmt.Println("╚════════════════════════════════════════════╝\033[0m")
	fmt.Println()
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
