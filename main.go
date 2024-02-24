package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Panic описывает ошибку для панического выхода из программы
type Panic struct {
	message string
}

func (p Panic) Error() string {
	return p.message
}

// arabicToRoman содержит соответствие римских цифр и арабских чисел
var arabicToRoman = map[string]int{
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

// romanNumerals содержит пары значений для преобразования арабских чисел в римские
var romanNumerals = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите пример:")
		scanner.Scan()
		input := scanner.Text()

		if input == "" {
			fmt.Println("Программа завершена.")
			return
		}

		result, err := calculate(input)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Результат: %v\n", result)
		}
	}
}

// calculate принимает строку с выражением, вычисляет и возвращает результат
func calculate(input string) (interface{}, error) {
	parts := strings.Split(input, " ")

	if len(parts) != 3 {
		return nil, Panic{"Паника!!! 😱😱😱 Неправильный ввод"}
	}

	if isMixed(parts[0], parts[2]) {
		return nil, Panic{"Паника!!! 😱😱😱 Нельзя одновременно вводить римские и арабские числа"}
	}

	a, err := convertToNumber(parts[0])
	if err != nil {
		return nil, err
	}

	op := parts[1]

	b, err := convertToNumber(parts[2])
	if err != nil {
		return nil, err
	}

	var result int

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return nil, Panic{"Паника!!! 😱😱😱 Деление на ноль"}
		}
		result = a / b
	default:
		return nil, Panic{"Паника!!! 😱😱😱 Недопустимая операция"}
	}

	if isRoman(parts[0]) || isRoman(parts[2]) {
		return convertToRoman(result)
	}

	return result, nil
}

// convertToNumber преобразует строку в число или возвращает ошибку
func convertToNumber(input string) (int, error) {
	if isRoman(input) {
		return arabicToRoman[input], nil
	}

	a, err := strconv.Atoi(input)
	if err != nil {
		return 0, Panic{"Паника!!! 😱😱😱 Неправильный ввод: " + input}
	}

	if a < 1 || a > 10 {
		return 0, Panic{"Паника!!! 😱😱😱 Неправильный ввод: " + input}
	}

	return a, nil
}

// isRoman проверяет, является ли строка римским числом
func isRoman(input string) bool {
	_, ok := arabicToRoman[input]
	return ok
}

// isMixed проверяет, содержит ли строка как римские, так и арабские числа
func isMixed(input1, input2 string) bool {
	hasRoman1 := isRoman(input1)
	hasRoman2 := isRoman(input2)

	return hasRoman1 && !hasRoman2 || !hasRoman1 && hasRoman2
}

// convertToRoman преобразует арабское число в римское
func convertToRoman(input int) (string, error) {
	if input <= 0 {
		return "", Panic{"Паника!!! 😱😱😱 Отрицательный результат или ноль"}
	}

	result := ""
	for _, numeral := range romanNumerals {
		for input >= numeral.Value {
			result += numeral.Symbol
			input -= numeral.Value
		}
	}

	return result, nil
}
