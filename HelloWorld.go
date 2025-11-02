package main

import (
	"fmt"
	"math"
)

func main() {
	PrintHeader("Калькулятор Жирности Массы Тела 💪")
	userKg, userHeight := getValidatedKg(), getValidatedHeight()
	IMT1 := calculateIMT(userKg, userHeight)

	clearScreen()

	outputResult(IMT1)

	GetUserExit()
}

func outputResult(IMT float64) {
	fmt.Printf("Индекс массы тела: %.0f \n", IMT)
	if IMT <= 18.5 {
		fmt.Println("Вы слишком худой! 🥵")
	} else if IMT > 18.5 && IMT < 25 {
		fmt.Println("Ваш вес в норме! 😎")
	} else if IMT >= 25 && IMT < 30 {
		fmt.Println("У вас избыточная масса тела! 🤔")
	} else if IMT >= 30 && IMT < 35 {
		fmt.Println("У вас ожирение первой степени! 😱")
	} else if IMT >= 35 && IMT < 40 {
		fmt.Println("У вас ожирение второй степени! 😨")
	} else if IMT >= 40 {
		fmt.Println("У вас ожирение третьей степени! 😱😱😱")
	}
}

func calculateIMT(userKg float64, userHeight float64) float64 {
	const IMTPower = 2
	const Meter = 100
	IMT := userKg / math.Pow(userHeight/Meter, IMTPower)
	return IMT
}
func getValidatedHeight() float64 {
	var userHeight float64
	for {
		fmt.Print("Введите свой рост (в сантиметрах): ")
		fmt.Scanln(&userHeight)
		if userHeight < 50 || userHeight > 260 {
			fmt.Println("Ошибка: Рост должен быть в рамках человеческой природы XD")
			continue // Повторить цикл, если пользователь ввел недопустимое значение
		}
		break // Выход из цикла, если пользователь ввел допустимое значение
	}
	return userHeight
}
func getValidatedKg() float64 {
	var userKg float64
	for {
		fmt.Print("Введите свой вес (в кг): ")
		fmt.Scanln(&userKg)
		if userKg < 25 || userKg > 600 {
			fmt.Println("Ошибка: Вес должен быть в рамках человеческой природы XD")
			continue // Повторить цикл, если пользователь ввел недопустимое значение
		}
		break // Выход из цикла, если пользователь ввел допустимое значение
	}
	return userKg
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
func GetUserExit() string {
	var ab1 string
	fmt.Println("\nНажмите Enter для выхода...")
	fmt.Scanln(&ab1)
	return ab1
}
