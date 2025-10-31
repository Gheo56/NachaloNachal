package main

import (
	"fmt"
	"math"
)

func main() {
	PrintHeader1("Калькулятор Жирности Массы Тела 💪")
	userKg1, userHeight1 := getUserInput1()
	IMT1 := calculateIMT1(userKg1, userHeight1)

	clearScreen1()

	outputResult1(IMT1)

	GetUserExit1()
}

func outputResult1(IMT1 float64) {
	result1 := fmt.Sprintf("Ваша жирность массы тела: %.0f", IMT1)
	fmt.Print(result1)
}

func calculateIMT1(userKg1 float64, userHeight1 float64) float64 {
	const IMTPower = 2
	const Meter = 100
	IMT1 := userKg1 / math.Pow(userHeight1/Meter, IMTPower)
	return IMT1
}
func getUserInput1() (float64, float64) {
	var userHeight1 float64
	var userKg1 float64
	fmt.Print("Введите свой рост (в сантиметрах): ")
	fmt.Scanln(&userHeight1)
	fmt.Print("Введите свой вес (в килограммах): ")
	fmt.Scanln(&userKg1)
	return userKg1, userHeight1

}
func PrintHeader1(title string) {
	fmt.Println("\033[36m╔════════════════════════════════════════════╗")
	fmt.Printf("║ %-32s         ║\n", title)
	fmt.Println("╚════════════════════════════════════════════╝\033[0m")
	fmt.Println()
}

func clearScreen1() {
	fmt.Print("\033[H\033[2J")
}
func GetUserExit1() string {
	var ab1 string
	fmt.Println("\nНажмите Enter для выхода...")
	fmt.Scanln(&ab1)
	return ab1
}
