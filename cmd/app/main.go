package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func romanToInt(s string) int {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}
	total := 0
	prevValue := 0

	for _, char := range s {
		value := romanNumerals[char]
		if value > prevValue {
			total += value - 2*prevValue
		} else {
			total += value
		}
		prevValue = value
	}

	return total
}

func intToRoman(num int) string {
	val := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	rom := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""

	for i := 0; num > 0; i++ {
		for num >= val[i] {
			num -= val[i]
			result += rom[i]
		}
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input:")
	inputValue, _ := reader.ReadString('\n')
	inputValue = strings.TrimSpace(inputValue)
	re := regexp.MustCompile(`^(\d+|[IVXLCDM]+)\s+([\+\-\*\/])\s+(\d+|[IVXLCDM]+)$`)
	matches := re.FindStringSubmatch(inputValue)

	if matches == nil {
		panic("Ошибка: неверный ввод!")
	}

	firstValueStr, operator, secondValueStr := matches[1], matches[2], matches[3]

	var firstValue, secondValue int
	var isRoman bool

	firstValueIsArabic := false
	secondValueIsArabic := false

	if _, err := strconv.Atoi(firstValueStr); err == nil {
		firstValue, _ = strconv.Atoi(firstValueStr)
		firstValueIsArabic = true
	} else {
		firstValue = romanToInt(firstValueStr)
	}

	if _, err := strconv.Atoi(secondValueStr); err == nil {
		secondValue, _ = strconv.Atoi(secondValueStr)
		secondValueIsArabic = true
	} else {
		secondValue = romanToInt(secondValueStr)
	}

	if firstValueIsArabic != secondValueIsArabic {
		panic("Ошибка: нельзя использовать одновременно арабские и римские числа!")
	}

	if !firstValueIsArabic && !secondValueIsArabic {
		isRoman = true
	}

	if !isRoman && (firstValue < 1 || firstValue > 10 || secondValue < 1 || secondValue > 10) {
		panic("Ошибка: числа должны быть от 1 до 10 включительно!")
	}

	var result int

	switch operator {
	case "+":
		result = firstValue + secondValue
	case "-":
		result = firstValue - secondValue
	case "*":
		result = firstValue * secondValue
	case "/":
		if secondValue == 0 {
			panic("Ошибка: деление на ноль!")
		}
		result = firstValue / secondValue
	default:
		panic("Ошибка: неверный оператор!")
	}
	if isRoman {
		if result < 1 {
			panic("Ошибка: результат римских чисел должен быть положительным!")
		}
		fmt.Println("Output:")
		fmt.Println(intToRoman(result))
	} else {
		fmt.Println("Output:")
		fmt.Println(result)
	}
}