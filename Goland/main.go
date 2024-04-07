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
// ПРИМИЧАНИЕ, Я СДЕЛАЛ ЛИШЬ ПОЛОВИНУ ЗАДАНИЕ НО НЕ СМОГ СДЕЛАТЬ ОДИН ДЮЙМ ОТ ЭТОГО ЗАДАНИЕ, ТАК КАК Я НЕ СМОГ РЕШИТЬ КАК СДЕЛАТЬ ОТВЕТ В РИМСКИХ ЧИСЛАХ И ДЛЯ ЭТОГО Я ДУМАЮ МНЕ НУЖНО ЕЩЁ ВРЕМЕНИ НО ВРЕМЕНИ НЕТУ, УВЫ ТАК ЧТО ПРИСЫЛАЮ КАК ЕСТЬ, В ЛЮБОМ СЛУЧАЕ СПАСИБО ЗА ТЗ И ЗА ПОНИМАНИЕ!
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
				result := num1 + num2
				if result <= 10 {
					fmt.Println(result)
				} else {
					panic("АХТУНГ!!! ОТВЕТ НЕ ДОЛЖЕН БЫТЬ БОЛЬШЕ ДЕСЯТИ!!!! ヘ⁠（⁠。⁠□⁠°⁠）⁠ヘ")
				}
			case "-":
				result := num1 - num2
				if result <= 10 {
					fmt.Println(result)
				} else {
					panic("АХТУНГ!!! ОТВЕТ НЕ ДОЛЖЕН БЫТЬ БОЛЬШЕ ДЕСЯТИ!!!! ヘ⁠（⁠。⁠□⁠°⁠）⁠ヘ")
				}
			case "*":
				result := num1 * num2
				if result <= 10 {
					fmt.Println(result)
				} else {
					panic("АХТУНГ!!! ОТВЕТ НЕ ДОЛЖЕН БЫТЬ БОЛЬШЕ ДЕСЯТИ!!!! ヘ⁠（⁠。⁠□⁠°⁠）⁠ヘ")
				}
			case "/":
				result := num1 / num2
				if result <= 10 {
					fmt.Println(result)
				} else {
					panic("АХТУНГ!!! ОТВЕТ НЕ ДОЛЖЕН БЫТЬ БОЛЬШЕ ДЕСЯТИ!!!! ヘ⁠（⁠。⁠□⁠°⁠）⁠ヘ")
				}
			default:
				panic("АХТУНГ!! ВЫ ВВЕЛИ НЕ ТОТ СИМВОЛ ИЛИ ЦИФРУ ヘ⁠（⁠。⁠□⁠°⁠）⁠ヘ")
			}
		} else {
			switch operate {
			case "+":
				sum := num1 + num2
				if sum <= 10 {
					fmt.Println(sum)
				} else {
					panic("АХТУНГ!!! ОТВЕТ НЕ ДОЛЖЕН БЫТЬ БОЛЬШЕ ДЕСЯТИ!!!! ヘ⁠（⁠。⁠□⁠°⁠）⁠ヘ")
				}
			case "-":
				sum := num1 - num2
				if sum <= 10 {
					fmt.Println(sum)
				} else {
					panic("АХТУНГ!!! ОТВЕТ НЕ ДОЛЖЕН БЫТЬ БОЛЬШЕ ДЕСЯТИ!!!! ヘ⁠（⁠。⁠□⁠°⁠）⁠ヘ")
				}
			case "*":
				sum := num1 * num2
				if sum <= 10 {
					fmt.Println(sum)
				} else {
					panic("АХТУНГ!!! ОТВЕТ НЕ ДОЛЖЕН БЫТЬ БОЛЬШЕ ДЕСЯТИ!!!! ヘ⁠（⁠。⁠□⁠°⁠）⁠ヘ")
				}
			case "/":
				sum := num1 / num2
				if sum <= 10 {
					fmt.Println(sum)
				} else {
					panic("АХТУНГ!!! ОТВЕТ НЕ ДОЛЖЕН БЫТЬ БОЛЬШЕ ДЕСЯТИ!!!!ヘ⁠（⁠。⁠□⁠°⁠）⁠ヘ ")
				}
			default:
				panic("АХТУНГ!! ВЫ ВВЕЛИ НЕ ТОТ СИМВОЛ ИЛИ ЦИФРУ ヘ⁠（⁠。⁠□⁠°⁠）⁠ヘ")
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
