package main

import (
	"fmt"
	"strconv"
	"strings"
)

var romeToArabic = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func main() {
	fmt.Println("ДОБРО ПОЖАЛОВАТЬ В КАЛЬКУЛЯТОР ДВОЙНЫХ СТАНДАРТОВ И ДВОЙНЫХ ЧИСЕЛ ВЫ ЗДЕСЬ МОЖЕТЕ ПРОВОДИТЬ КАЛЬКУЛИТИВНЫЕ ОПЕРАЦИИ НА АРАБСКИХ И РИМСКИХ ЧИСЕЛ")
	fmt.Println("ЧТОБЫ ИСПОЛЬЗОВАТЬ КАЛЬКУЛЯТОР ПРОСТО НАПИШИТЕ ОПЕРАЦИЮ ОДНОЙ СТРОКОЙ")
	for {
		var num1Str string
		var operate string
		var num2Str string

		var num1 int
		var num2 int

		fmt.Scan(&num1Str, &operate, &num2Str)

		num1, _ = strconv.Atoi(num1Str)
		num2, _ = strconv.Atoi(num2Str)

		if IsRomanNumbers(num1Str) {
			num1 = romeToArabic[strings.TrimSpace(num1Str)]
			num2 = romeToArabic[strings.TrimSpace(num2Str)]

			switch operate {
			case "+":
				sum := num1 + num2
				fmt.Println(sum)
			case "-":
				sum := num1 - num2
				fmt.Println(sum)
			case "*":
				sum := num1 * num2
				fmt.Println(sum)
			case "/":
				sum := num1 / num2
				fmt.Println(sum)
			default:
				panic("АХТУНГ!! ВЫ ВВЕЛИ НЕ ТОТ СИМВОЛ ИЛИ ЦИФРУ")
			}
		} else {
			switch operate {
			case "+":
				sum := num1 + num2
				fmt.Println(sum)
			case "-":
				sum := num1 - num2
				fmt.Println(sum)
			case "*":
				sum := num1 * num2
				fmt.Println(sum)
			case "/":
				sum := num1 / num2
				fmt.Println(sum)
			default:
				panic("АХТУНГ!! ВЫ ВВЕЛИ НЕ ТОТ СИМВОЛ ИЛИ ЦИФРУ")
			}
		}
	}
}

func IsRomanNumbers(s string) bool {
	for _, c := range s {
		if _, ok := romeToArabic[string(c)]; !ok {
			return false
		}
	}
	return true
}
