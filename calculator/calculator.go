package calculator

import "errors"

func DoOperation(dig1, dig2 float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return dig1 + dig2, nil
	case "-":
		return dig1 - dig2, nil
	case "*":
		return dig1 * dig2, nil
	case "/":
		return dig1 / dig2, nil
	default:
		return 0, errors.New("Error operation")
	}
}
