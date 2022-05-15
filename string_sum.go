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
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\t", "")
	input = strings.ReplaceAll(input, "\v", "")
	input = strings.ReplaceAll(input, "\r", "")
	input = strings.ReplaceAll(input, "\f", "")
	if input == "" {
		return "", errorEmptyInput
	}
	if len(input) < 3 {
		return "", errorNotTwoOperands
	}

	arr := strings.Split(input, "+")

	if len(arr) == 1 {
		arr = strings.Split(input, "-")
		if len(arr) != 2 {
			return "", errorNotTwoOperands
		}
		return calculator(arr[0], arr[1])
	}

	if len(arr) != 2 {
		return "", errorNotTwoOperands
	}

	return calculator(arr[0], arr[1])
}
func calculator(str1, str2 string) (string, error) {
	num1, errNum1 := strconv.Atoi(str1)
	if errNum1 != nil {
		return "", fmt.Errorf("invalid input. request must have numeric type")
	}

	num2, errNum2 := strconv.Atoi(str2)
	if errNum2 != nil {
		return "", fmt.Errorf("invalid input. request must have numeric type")
	}

	return strconv.Itoa(num1 + num2), nil
}
