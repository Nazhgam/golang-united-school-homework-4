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

	input = strings.TrimSpace(input)
	if input == "" {
		return "", errorEmptyInput
	}

	ind := strings.LastIndex(input, "-")
	if ind > 0 {
		num1, errNum1 := strconv.Atoi(input[:ind])
		num2, errNum2 := strconv.Atoi(input[ind:])
		if errNum1 != nil || errNum2 != nil {
			return "", fmt.Errorf("alphabetical input don't allowed")
		}

		output = strconv.Itoa(num1 + num2)
		return output, nil
	}
	arr := strings.Split(input, "+")
	if len(arr) != 2 {
		return "", errorNotTwoOperands
	}
	num1, errNum1 := strconv.Atoi(arr[0])
	num2, errNum2 := strconv.Atoi(arr[1])
	if errNum1 != nil || errNum2 != nil {
		return "", fmt.Errorf("alphabetical input don't allowed")
	}

	output = strconv.Itoa(num1 + num2)
	return output, nil

}
