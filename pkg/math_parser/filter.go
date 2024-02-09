package mathparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	// регулярка математического выражения
	isExpr = regexp.MustCompile("^(\\d+\\.?\\d*|\\.\\d+)( [\\+\\-\\*\\/] (\\d+\\.?\\d*|\\.\\d+))*$").MatchString
)

func parceString(message string) error {
	// проверяем, является ли сообщение математическим выражением
	if !isExpr(message) {
		return fmt.Errorf("Bad message")
	}

	// разбиваем строку на токены
	tokens := strings.Split(message, " ")

	operands := make([]float64, len(tokens)/2+1) // операнды (числа)
	for i := range operands {                    // заполняем слайс операндов
		operands[i], _ = strconv.ParseFloat(tokens[i*2], 64)
	}

	operators := make([]string, len(tokens)/2) // операторы (операции)
	for i := range operators {                 // заполняем слайс операций
		operators[i] = tokens[i*2+1]
	}

	// меняем деление на умножение и вычитание на сложение
	for i, op := range operators {
		if op == "/" {
			operators[i] = "*"
			operands[i+1] = 1 / operands[i+1]
		} else if op == "-" {
			operators[i] = "+"
			operands[i+1] = -operands[i+1]
		}
	}

	return nil
}
