package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func romanToArabic(roman string) (int, error) {

	var (
		romanNumerals = map[string]int{

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
			"L":    50,
			"C":    100,
		}
	)

	var result int
	prevValue := 0

	for _, char := range roman {
		value, found := romanNumerals[string(char)]
		if !found {
			return 0, fmt.Errorf("недопустимый символ римской цифры: %c", char)
		}

		result += value

		if value > prevValue {
			result -= 2 * prevValue
		}

		prevValue = value
	}

	return result, nil
}

func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("недопустимая математическая операция: %s", operator)
	}
}

func evaluate(expression string) (int, error) {
	parts := strings.Fields(expression)

	if len(parts) != 3 {
		return 0, fmt.Errorf("неверный формат ввода. Ожидается: число оператор число")
	}

	a, err := strconv.Atoi(parts[0])
	if err != nil {
		if matched, _ := regexp.MatchString(`^[IVXLCDM]+$`, parts[0]); matched {
			a, err = romanToArabic(parts[0])
			if err != nil {
				return 0, err
			}
		} else {
			return 0, fmt.Errorf("недопустимое число: %s", parts[0])
		}
	}

	operator := parts[1]

	b, err := strconv.Atoi(parts[2])
	if err != nil {
		if matched, _ := regexp.MatchString(`^[IVXLCDM]+$`, parts[2]); matched {
			b, err = romanToArabic(parts[2])
			if err != nil {
				return 0, err
			}
		} else {
			return 0, fmt.Errorf("недопустимое число: %s", parts[2])
		}
	}

	if (a < 1 || a > 10) || (b < 1 || b > 10) {
		return 0, fmt.Errorf("числа должны быть от 1 до 10 включительно")
	}

	return calculate(a, b, operator)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Использование: калькулятор \"выражение\"")
		os.Exit(1)
	}

	expression := os.Args[1]

	result, err := evaluate(expression)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Результат:", result)
}
