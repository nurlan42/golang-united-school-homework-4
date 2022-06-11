package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	errorWrongOperator  = errors.New("wrong operator. Please, choose either plus or minus signs")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	var sign int

	for i, el := range input {
		if i == 0 && string(el) == "+" {
			sign = 1
			input = input[i+1:]
		} else if i == 0 && string(el) == "-" {
			sign = -1
			input = input[i+1:]
		}
	}

	nbrs := strings.Split(input, "+")
	operator := "+"
	if len(nbrs) != 2 {
		nbrs = strings.Split(input, "-")
		operator = "-"
	}
	if len(nbrs) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	nbr1, err := strconv.Atoi(nbrs[0])
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	nbr2, err := strconv.Atoi(nbrs[1])
	if err != nil {
		return "", fmt.Errorf("%w", err)

	}

	if sign == -1 {
		nbr1 = -nbr1
	}

	switch operator {
	case "+":
		res := nbr1 + nbr2
		return strconv.Itoa(res), nil
	case "-":
		res := nbr1 - nbr2
		return strconv.Itoa(res), nil
	default:
		return "", fmt.Errorf("%w", errorWrongOperator)
	}
}
