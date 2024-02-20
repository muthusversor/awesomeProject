package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Panic struct {
	message string
}

func (p Panic) Error() string {
	return p.message
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите пример:") // Считываем ввод
		scanner.Scan()                 // Сканируем ввод
		input := scanner.Text()        // Получаем текст из ввода

		if input == "" {
			fmt.Println("Программа завершена.") // Если ввод пустой, завершаем программу
			return
		}

		result, err := calculate(input) // Вызываем функцию calculate для обработки введенного примера
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Результат: %v\n", result)
		}
	}
}

func calculate(input string) (interface{}, error) {
	parts := strings.Split(input, " ") // Разделяем введенную строку на составляющие части

	if len(parts) != 3 {
		return nil, Panic{"Паника!!!😱😱😱 неправильный ввод"} // Возвращаем ошибку, если введенный пример некорректен
	}

	a, err := convertToNumber(parts[0]) // Преобразуем первую часть примера в число
	if err != nil {
		return nil, err
	}

	op := parts[1] // Получаем оператор

	b, err := convertToNumber(parts[2]) // Преобразуем вторую часть примера в число
	if err != nil {
		return nil, err
	}

	var result int // Инициализируем результат переменной типа int

	switch op { // Выполняем заданную операцию в зависимости от оператора
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return nil, Panic{"Паника!!!😱😱😱 деление на ноль"} // Возвращаем ошибку при попытке деления на ноль
		} // Возвращаем ошибку при делении на 0
		result = a / b
	default:
		return nil, Panic{"Паника!!!😱😱😱 недопустимая операция"} // Возвращаем ошибку при недопустимой операции
	}

	if isRoman(parts[0]) && isRoman(parts[2]) { // Проверяем, являются ли оба числа римскими
		return convertToRoman(result), nil // Возвращаем результат вычисления в римском формате
	}

	return result, nil // Возвращаем результат вычисления в виде арабского числа
}

func convertToNumber(input string) (int, error) {
	arabicToRoman := map[string]int{ // Создаем map для преобразования римских чисел в арабские
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

	if num, ok := arabicToRoman[input]; ok { // Проверяем, является ли введенное число римским
		return num, nil
	}

	if a, err := strconv.Atoi(input); err == nil { // Пробуем преобразовать введенное число в числовой формат
		return a, nil
	}

	return 0, Panic{"Паника!!!😱😱😱 неправильный ввод: " + input}
}

func isRoman(input string) bool {
	arabicToRoman := map[string]int{ // Создаем map для преобразования римских чисел в арабские
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

	_, ok := arabicToRoman[input]
	return ok
}

func convertToRoman(input int) string {
	romanNumerals := []struct {
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

	result := ""
	for _, numeral := range romanNumerals {
		for input >= numeral.Value {
			result += numeral.Symbol
			input -= numeral.Value
		}
	}

	return result
}
